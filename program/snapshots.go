package program

import (
	"github.com/deweysasser/golang-program/cumulus"
)

type Snapshots struct {
	List SnapshotList `cmd:""`
}

type SnapshotList struct {
	CommonList `embed:""`
}

func (list *SnapshotList) Run() error {
	return doList[cumulus.Snapshot](&list.CommonList,
		cumulus.RegionalAccounts.Snapshots,
		"snapshot")
}
