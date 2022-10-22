package program

import (
	"context"
	"fmt"
	"github.com/deweysasser/golang-program/cumulus"
	"github.com/deweysasser/golang-program/cumulus/caws"
	"github.com/deweysasser/golang-program/cumulus/stats"
	"github.com/rs/zerolog/log"
	"strings"
	"sync/atomic"
)

type Instances struct {
	List InstanceList `cmd:""`
}

type InstanceList struct {
	CredentialsFile string `group:"AWS" short:"c" help:"AWS Credentials File" type:"existingfile" default:"~/.aws/credentials"`
	Arg             string `arg:"" optional:""`
}

func (list *InstanceList) Run() error {
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
		return ra.VisitInstance(context.Background(), func(ctx context.Context, cancel context.CancelFunc, info cumulus.Instance) error {
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
		log.Debug().Msg("Listing all instances")
		return ra.VisitInstance(context.Background(), func(ctx context.Context, cancel context.CancelFunc, info cumulus.Instance) error {
			count.Add(1)
			a := cumulus.RegionalAccountContext(ctx)
			fmt.Print(a.Name(), "\t", a.Region(), "\t")
			fmt.Println(info.Text())
			return nil
		})

	}

}
