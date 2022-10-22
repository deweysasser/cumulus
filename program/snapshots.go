package program

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"github.com/deweysasser/golang-program/cumulus"
	"github.com/deweysasser/golang-program/cumulus/caws"
	"github.com/deweysasser/golang-program/cumulus/stats"
	"github.com/hashicorp/go-multierror"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
)

type Snapshots struct {
	List   SnapshotList   `cmd:""`
	Delete SnapshotDelete `cmd:""`
}

type SnapshotList struct {
	CredentialsFile string `group:"AWS" short:"c" help:"AWS Credentials File" type:"existingfile" default:"~/.aws/credentials"`
	Arg             string `arg:"" optional:""`
}

func (list *SnapshotList) Run() error {
	accounts, err := caws.AvailableAccountsFrom(list.CredentialsFile)
	if err != nil {
		return err
	}

	ra := accounts.Unique().InRegion("us-east-2")
	log.Debug().Str("ras", fmt.Sprint(ra)).Msg("Using accounts")
	var count atomic.Int32

	defer func() {
		fmt.Println("Visited", count.Load())
		stats.Report()
	}()

	log.Debug().Str("arg", list.Arg).Msg("Searching for this IP")
	if list.Arg != "" {
		return ra.VisitSnapshot(context.Background(), func(ctx context.Context, cancel context.CancelFunc, info cumulus.Snapshot) error {
			count.Add(1)
			s := info.Text()
			if strings.Contains(s, list.Arg) {
				log.Ctx(ctx).Info().Msg("Found instance and exiting")
				fmt.Println(info.Text())
				cancel()
			}
			return nil
		})
	} else {
		log.Debug().Msg("Listing all snapshots")
		return ra.VisitSnapshot(context.Background(), func(ctx context.Context, cancel context.CancelFunc, info cumulus.Snapshot) error {
			count.Add(1)
			a := cumulus.RegionalAccountContext(ctx)
			fmt.Print(a.Name(), "\t", a.Region(), "\t")
			fmt.Println(info.Text())
			return nil
		})

	}

}

type SnapshotDelete struct {
	CredentialsFile string   `group:"AWS" short:"c" help:"AWS Credentials File" type:"existingfile" default:"~/.aws/credentials"`
	DryRun          bool     `group:"AWS" short:"n" help:"Do not actually delete"`
	Arg             []string `arg:"" optional:""`
}

func (cmd *SnapshotDelete) Run() error {
	accounts, err := caws.AvailableAccountsFrom(cmd.CredentialsFile)
	if err != nil {
		return err
	}

	victimMap := make(map[cumulus.ID]bool)
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	ctx, c := context.WithCancel(context.Background())
	defer log.Debug().Msg("Run exiting")

	cancel := func() {
		log.Debug().Msg("Canceling context")
		c()
	}
	defer cancel()

	go func() {
		sig := <-sigs

		signal.Reset(syscall.SIGINT, syscall.SIGTERM)
		cancel()

		log.Debug().
			Str("signal", sig.String()).
			Msg("Canceling on signal")
		//time.Sleep(5 * time.Second)
		//log.Error().Msg("Exiting on timeout")
		//os.Exit(1)
	}()

	for _, victim := range cmd.Arg {
		if victim == "-" {
			log.Debug().Msg("Reading from stdin")
			readStdin(&victimMap)
		}
		victimMap[cumulus.ID(victim)] = true
	}

	log.Debug().Int("count", len(victimMap)).Msg("Search for one of victims")
	if len(victimMap) < 1 {
		return errors.New("must delete at least one item")
	}

	ra := accounts.Unique().InRegion("us-east-2")
	log.Debug().Str("ras", fmt.Sprint(ra)).Msg("Using accounts")

	deletions := stats.NewCounter(ctx, "deletions")
	deletionErrors := stats.NewCounter(ctx, "deletion_errors")
	deletionTimer := stats.NewTimer(ctx, "deletion_calls")
	retries := stats.NewCounter(ctx, "deletion_retries")

	defer func() {
	}()
	var count atomic.Int32
	defer func() {
		fmt.Println("Visited", count.Load())
		stats.Report()
	}()

	victims := make(chan Worker)

	var deleteErrors error

	errors := workerPool(ctx, 15, victims)

	go func() {
		defer log.Debug().Msg("victims closed")
		defer close(victims)

		// TODO:  handle the error return value of this function
		ra.VisitSnapshot(ctx, func(ctx context.Context, cancel context.CancelFunc, info cumulus.Snapshot) error {
			count.Add(1)
			//a := cumulus.RegionalAccountContext(ctx)
			logger := log.Ctx(ctx)
			if victimMap[info.Id()] {
				victims <- func() error {

					logger.Info().Bool("dry_run", cmd.DryRun).Msg("Deleting snapshot")

					defer deletionTimer.Call().Done()

					err := retry(ctx, retries, 10,
						func() error {
							logger.Info().Bool("dry_run", cmd.DryRun).Err(ctx.Err()).Msg("calling AWS delete")
							return info.Delete(ctx, cmd.DryRun)
						})
					if err != nil {
						deletionErrors.Inc()
					} else {
						deletions.Inc()
					}
					return err
				}
			}

			return nil
		})
	}()

	for e := range errors {
		deleteErrors = multierror.Append(deleteErrors, e)
	}

	log.Debug().Msg("errors closed")
	err = multierror.Append(err, deleteErrors)

	log.Debug().Msg("returning summary of errors")
	return err
}

// retry retries the given function the specified number of times
func retry(ctx context.Context, retries *stats.Counter, nubmerOfRetries int, f func() error) error {
	var err error
	for count := 0; count < nubmerOfRetries; count++ {
		err = f()
		if err == nil {
			return nil
		} else if ctx.Err() != nil {
			return err
		} else {
			retries.Inc()
			log.Debug().Err(err).Int("count", count).Msg("attempt failed.  Retrying")
		}
	}
	return err
}

func readStdin(m *map[cumulus.ID]bool) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := scanner.Text()
		//log.Debug().Str("victim", text).Msg("looking to delete")
		(*m)[cumulus.ID(text)] = true
	}
}

type Worker func() error

func workerPool(ctx context.Context, size int, c <-chan Worker) chan error {
	wg := sync.WaitGroup{}
	errors := make(chan error, size)

	wg.Add(size)

	for i := 0; i < size; i++ {
		go func() {
			defer wg.Done()
			for worker := range c {
				if ctx.Err() != nil {
					errors <- ctx.Err()
					return
				}
				e := worker()
				if e != nil {
					log.Error().Err(e).Msg("Error processing worker")
					errors <- e
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		log.Debug().Msg("Closing errors")
		close(errors)
	}()

	return errors
}
