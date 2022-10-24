package program

import (
	"github.com/deweysasser/cumulus/cumulus"
)

type VolumeList struct {
	CommonList `embed:""`
}

func (list *VolumeList) Run(program Options) error {
	return listOnRegionalAccounts[cumulus.Volume](
		program,
		&list.CommonList,
		cumulus.RegionalAccounts.Volumes,
		"volume")
}
