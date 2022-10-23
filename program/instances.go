package program

import (
	"github.com/deweysasser/cumulus/cumulus"
)

type Instances struct {
	List InstanceList `cmd:""`
}

type InstanceList struct {
	CommonList `embed:""`
}

func (list *InstanceList) Run() error {
	return listOnRegionalAccounts[cumulus.Instance](&list.CommonList,
		cumulus.RegionalAccounts.Instances,
		"instance")
}
