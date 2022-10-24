package program

import (
	"github.com/deweysasser/cumulus/cumulus"
)

type ZoneList struct {
	CommonList `embed:""`
}

func (list *ZoneList) Run(program Options) error {
	return listOnAccounts[cumulus.Zone](
		program,
		&list.CommonList,
		cumulus.Accounts.Zones,
		"Zone")
}

type RecordList struct {
	CommonList `embed:""`
}

func (list *RecordList) Run(program Options) error {
	return listOnAccounts[cumulus.NameRecord](
		program,
		&list.CommonList,
		cumulus.Accounts.NameRecords,
		"NameRecord")
}
