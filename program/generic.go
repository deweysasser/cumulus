package program

import (
	"context"
	"fmt"
	"github.com/deweysasser/cumulus/cumulus"
	"github.com/deweysasser/cumulus/cumulus/caws"
	"github.com/deweysasser/cumulus/cumulus/stats"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"sync/atomic"
	"time"
)

type CommonList struct {
	CredentialsFile string   `group:"AWS" short:"c" help:"AWS Credentials File" type:"existingfile" default:"~/.aws/credentials"`
	Include         []string `group:"Output" short:"I" help:"List of field regexps to include"`
	Exclude         []string `group:"Output" short:"X" help:"List of fields regexps to exclude"`
	IncludeAll      bool     `group:"Output" short:"A" help:"Include all fields"`

	// TODO:  support a "fields" list to specify the membership and order of the fields printed
}

func Display[T cumulus.Common](options Options, list *CommonList, typename string, items chan T) {
	var count atomic.Int32
	defer func() {
		log.Info().Int32("count", count.Load()).Msg("discovered " + typename)
		stats.Report()
		stats.Total()
	}()

	var f *Filter
	if list.IncludeAll || options.Verbose {
		log.Debug().Msg("Including all fields.  Removing filter")
		f = NoFilter
	} else {
		f = NewFilter(list.Include, list.Exclude)
	}

	acc := cumulus.NewAccumulator()
	for info := range items {
		count.Add(1)
		acc.Add(info)
	}
	acc.Print(f, !options.Quiet)
}

func listOnRegionalAccounts[T cumulus.Common](options Options, list *CommonList, method func(account cumulus.RegionalAccounts, ctx context.Context) chan T, typename string) error {
	start := time.Now()
	defer func() {
		log.Info().Dur("duration", time.Since(start)).Msg("run time")
	}()
	log.Debug().Str("file", list.CredentialsFile).Msg("Opening file")

	accounts, err := caws.AvailableAccountsFrom(list.CredentialsFile)
	if err != nil {
		return err
	}

	ra := accounts.Unique(cumulus.WithErrorHandler(context.Background(), cumulus.IgnoreErrors)).InRegion(caws.DefaultRegions...)

	if len(ra) == 0 {
		return errors.New("No accounts discovered")
	}

	collectedErrors := cumulus.ErrorCollector{}
	ctx := cumulus.WithErrorHandler(context.Background(), collectedErrors.Handle)

	log.Debug().Str("ras", fmt.Sprint(ra)).Msg("Using accounts")
	log.Debug().Msg("Listing all " + typename)
	items := method(ra, ctx)

	Display(options, list, typename, items)
	return collectedErrors.Error
}

func listOnAccounts[T cumulus.Common](options Options, list *CommonList, method func(account cumulus.Accounts, ctx context.Context) chan T, typename string) error {
	start := time.Now()
	defer func() {
		log.Info().Dur("duration", time.Since(start)).Msg("run time")
	}()

	log.Debug().Str("file", list.CredentialsFile).Msg("Opening file")

	accounts, err := caws.AvailableAccountsFrom(list.CredentialsFile)
	if err != nil {
		return err
	}

	ra := accounts.Unique(cumulus.WithErrorHandler(context.Background(), cumulus.IgnoreErrors))

	if len(ra) == 0 {
		return errors.New("No accounts discovered")
	}

	errors := cumulus.ErrorCollector{}
	ctx := cumulus.WithErrorHandler(context.Background(), errors.Handle)

	log.Debug().Str("ras", fmt.Sprint(ra)).Msg("Using accounts")
	log.Debug().Msg("Listing all " + typename)

	items := method(ra, ctx)

	Display(options, list, typename, items)
	return errors.Error
}
