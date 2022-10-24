package program

import (
	"github.com/deweysasser/cumulus/cumulus"
)

type InstanceList struct {
	CommonList `embed:""`
}

func (list *InstanceList) Run(program Options) error {
	return listOnRegionalAccounts[cumulus.Instance](
		program,
		&list.CommonList,
		cumulus.RegionalAccounts.Instances,
		"instance")
}
