package program

import (
	"github.com/deweysasser/cumulus/cumulus"
)

type DBInstanceList struct {
	CommonList `embed:""`
}

func (list *DBInstanceList) Run(program Options) error {
	return listOnRegionalAccounts[cumulus.DBInstance](
		program,
		&list.CommonList,
		cumulus.RegionalAccounts.DBInstances,
		"volume")
}
