package program

import (
	"github.com/deweysasser/golang-program/cumulus"
)

type Instances struct {
	List InstanceList `cmd:""`
}

type InstanceList struct {
	CommonList `embed:""`
}

func (list *InstanceList) Run() error {
	return doList[cumulus.Instance](&list.CommonList,
		cumulus.RegionalAccounts.Instances,
		"instance")
}
