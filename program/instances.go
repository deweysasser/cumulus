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

	ra := accounts.Unique(cumulus.WithErrorHandler(context.Background(), cumulus.IgnoreErrors)).InRegion("us-east-2")

	log.Debug().Str("ras", fmt.Sprint(ra)).Msg("Using accounts")
	var count atomic.Int32
	defer func() {
		log.Info().Int32("count", count.Load()).Msg("discovered instances")
		stats.Report()
	}()

	log.Debug().Str("arg", list.Arg).Msg("Searching for this IP")
	if list.Arg != "" {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		for info := range ra.Instances(ctx) {
			count.Add(1)
			s := info.Text()
			if strings.Contains(s, list.Arg) {
				logger := zerolog.Ctx(info.Ctx())
				logger.Info().Msg("Found instance and exiting")
				fmt.Println(info.Text())
				cancel()
			}
		}
		return nil
	} else {
		log.Debug().Msg("Listing all instances")
		for info := range ra.Instances(context.Background()) {
			count.Add(1)
			fmt.Println(info.Text())
		}
		return nil
		//return ra.VisitInstance(context.Background(), func(ctx context.Context, cancel context.CancelFunc, info cumulus.Instance) error {
		//	count.Add(1)
		//	a := cumulus.RegionalAccountContext(ctx)
		//	fmt.Print(a.Name(), "\t", a.Region(), "\t")
		//})

	}

}
