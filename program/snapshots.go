package program

import (
	"github.com/deweysasser/cumulus/cumulus"
)

type SnapshotList struct {
	CommonList `embed:""`
}

func (list *SnapshotList) Run(program Options) error {
	return listOnRegionalAccounts[cumulus.Snapshot](
		program,
		&list.CommonList,
		cumulus.RegionalAccounts.Snapshots,
		"snapshot")
}
