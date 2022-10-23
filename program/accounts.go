package program

import (
	"github.com/deweysasser/cumulus/cumulus"
)

type Accounts struct {
	List List `cmd:""`
}

type List struct {
	CommonList
}

func (list *List) Run(program Options) error {
	return listOnAccounts[cumulus.AccountInfo](
		program,
		&list.CommonList,
		cumulus.Accounts.AccountInfos,
		"AccountInfo")
}
