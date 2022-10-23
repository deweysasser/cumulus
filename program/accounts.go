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

func (list *List) Run() error {
	return doAccountList[cumulus.AccountInfo](&list.CommonList,
		cumulus.Accounts.AccountInfos,
		"AccountInfo")
}
