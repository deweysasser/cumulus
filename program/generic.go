package program

import (
	"context"
	"fmt"
	"github.com/deweysasser/golang-program/cumulus"
	"github.com/deweysasser/golang-program/cumulus/caws"
	"github.com/deweysasser/golang-program/cumulus/stats"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"strings"
	"sync/atomic"
	"time"
)

type Texter interface {
	Text() string
}

type CommonList struct {
	CredentialsFile string `group:"AWS" short:"c" help:"AWS Credentials File" type:"existingfile" default:"~/.aws/credentials"`
	Arg             string `arg:"" optional:""`
}

func doList[T cumulus.Common](list *CommonList, method func(account cumulus.RegionalAccounts, ctx context.Context) chan T, typename string) error {
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

	errors := cumulus.ErrorCollector{}
	ctx := cumulus.WithErrorHandler(context.Background(), errors.Handle)

	log.Debug().Str("ras", fmt.Sprint(ra)).Msg("Using accounts")
	var count atomic.Int32
	defer func() {
		log.Info().Int32("count", count.Load()).Msg("discovered " + typename)
		stats.Report()
		stats.Total()
	}()

	log.Debug().Str("arg", list.Arg).Msg("Searching for this IP")
	if list.Arg != "" {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		for info := range method(ra, ctx) {
			count.Add(1)
			s := info.Text()
			if strings.Contains(s, list.Arg) {
				logger := zerolog.Ctx(info.Ctx())
				logger.Info().Msg("Found " + typename + " and exiting")
				fmt.Println(info.Text())
				cancel()
			}
		}
		return errors.Error
	} else {
		log.Debug().Msg("Listing all " + typename)
		for info := range method(ra, ctx) {
			count.Add(1)
			fmt.Println(info.Text())
		}
		return errors.Error
	}

}

func doAccountList[T cumulus.Common](list *CommonList, method func(account cumulus.Accounts, ctx context.Context) chan T, typename string) error {
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

	errors := cumulus.ErrorCollector{}
	ctx := cumulus.WithErrorHandler(context.Background(), errors.Handle)

	log.Debug().Str("ras", fmt.Sprint(ra)).Msg("Using accounts")
	var count atomic.Int32
	defer func() {
		log.Info().Int32("count", count.Load()).Msg("discovered " + typename)
		stats.Report()
		stats.Total()
	}()

	log.Debug().Str("arg", list.Arg).Msg("Searching for this IP")
	if list.Arg != "" {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		for info := range method(ra, ctx) {
			count.Add(1)
			s := info.Text()
			if strings.Contains(s, list.Arg) {
				logger := zerolog.Ctx(info.Ctx())
				logger.Info().Msg("Found " + typename + " and exiting")
				fmt.Println(info.Text())
				cancel()
			}
		}
		return errors.Error
	} else {
		log.Debug().Msg("Listing all " + typename)
		for info := range method(ra, ctx) {
			count.Add(1)
			fmt.Println(info.Text())
		}
		return errors.Error
	}
}
