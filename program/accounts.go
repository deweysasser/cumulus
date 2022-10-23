package program

import (
	"context"
	"fmt"
	"github.com/deweysasser/golang-program/cumulus"
	"github.com/deweysasser/golang-program/cumulus/caws"
)

type Accounts struct {
	List List `cmd:""`
}

type List struct {
	CredentialsFile string `group:"AWS" short:"c" help:"AWS Credentials File" type:"existingfile" default:"~/.aws/credentials"`
}

func (list *List) Run() error {
	accounts, err := caws.AvailableAccountsFrom(list.CredentialsFile)
	if err != nil {
		return err
	}

	ctx := cumulus.WithErrorHandler(context.Background(), cumulus.IgnoreErrors)

	for info := range accounts.AccountInfos(ctx) {
		fmt.Println(info.Name(), "\t", info.ID())
	}

	return nil
}
