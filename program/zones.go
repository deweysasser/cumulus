package program

import (
	"context"
	"fmt"
	"github.com/deweysasser/golang-program/cumulus"
	"github.com/deweysasser/golang-program/cumulus/caws"
)

type Zones struct {
	List ZoneList `cmd:""`
}

type ZoneList struct {
	CredentialsFile string `group:"AWS" short:"c" help:"AWS Credentials File" type:"existingfile" default:"~/.aws/credentials"`
	Arg             string `arg:"" optional:""`
}

func (list *ZoneList) Run() error {
	accounts, err := caws.AvailableAccountsFrom(list.CredentialsFile)
	if err != nil {
		return err
	}

	return accounts.VisitZones(context.Background(), func(ctx context.Context, cancel context.CancelFunc, info cumulus.Zone) error {
		acct := cumulus.AccountContext(ctx)
		fmt.Print(acct)
		fmt.Print("\t")
		fmt.Println(info.Text())
		return nil
	})
}
