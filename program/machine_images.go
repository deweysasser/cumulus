package program

import (
	"github.com/deweysasser/cumulus/cumulus"
)

type MachineImageList struct {
	CommonList `embed:""`
}

func (list *MachineImageList) Run(program Options) error {
	return listOnRegionalAccounts[cumulus.MachineImage](
		program,
		&list.CommonList,
		cumulus.RegionalAccounts.MachineImages,
		"instance")
}
