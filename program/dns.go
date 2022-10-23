package program

import (
	"github.com/deweysasser/cumulus/cumulus"
)

type DNS struct {
	Zone   Zones   `cmd:""`
	Record Records `cmd:""`
}

type Zones struct {
	List ZoneList `cmd:""`
}

type ZoneList struct {
	CommonList `embed:""`
}

func (list *ZoneList) Run(program *Options) error {
	return listOnAccounts[cumulus.Zone](&list.CommonList,
		cumulus.Accounts.Zones,
		"Zone")
}

type Records struct {
	List RecordList `cmd:""`
}

type RecordList struct {
	CommonList `embed:""`
}

func (list *RecordList) Run(program *Options) error {
	return listOnAccounts[cumulus.NameRecord](&list.CommonList,
		cumulus.Accounts.NameRecords,
		"NameRecord")
}
