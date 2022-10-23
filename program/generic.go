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

type Texter interface {
	Text() string
}

type CommonList struct {
	CredentialsFile string   `group:"AWS" short:"c" help:"AWS Credentials File" type:"existingfile" default:"~/.aws/credentials"`
	Fields          []string `group:"Output" short:"i" help:"List of field regexps to include"`
	Exclude         []string `group:"Output" short:"x" help:"List of fields regexps to exclude" default:"tag:.*"`
	IncludeAll      bool     `group:"Output" short:"A" help:"Include all fields"`
	// TODO:  support a "fields" list to specify the membership and order of the fields printed
}

func listOnRegionalAccounts[T cumulus.Common](list *CommonList, method func(account cumulus.RegionalAccounts, ctx context.Context) chan T, typename string) error {
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
	log.Debug().Str("ras", fmt.Sprint(ra)).Msg("Using accounts")

	collectedErrors := cumulus.ErrorCollector{}
	ctx := cumulus.WithErrorHandler(context.Background(), collectedErrors.Handle)

	var count atomic.Int32
	defer func() {
		log.Info().Int32("count", count.Load()).Msg("discovered " + typename)
		stats.Report()
		stats.Total()
	}()

	items := method(ra, ctx)

	Display(list, typename, items, count)
	return collectedErrors.Error

}

func Display[T cumulus.Common](list *CommonList, typename string, items chan T, count atomic.Int32) {
	f := NewFilter(list.Fields, list.Exclude)

	log.Debug().Msg("Listing all " + typename)
	acc := cumulus.NewAccumulator()
	for info := range items {
		count.Add(1)
		acc.Add(info.Fields())
	}
	acc.Print(f)
}

func listOnAccounts[T cumulus.Common](list *CommonList, method func(account cumulus.Accounts, ctx context.Context) chan T, typename string) error {
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
	var count atomic.Int32
	defer func() {
		log.Info().Int32("count", count.Load()).Msg("discovered " + typename)
		stats.Report()
		stats.Total()
	}()

	items := method(ra, ctx)

	Display(list, typename, items, count)
	return errors.Error
}
