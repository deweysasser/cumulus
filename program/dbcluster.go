package program

import (
	"github.com/deweysasser/cumulus/cumulus"
)

type ClusterList struct {
	CommonList `embed:""`
}

func (list *ClusterList) Run(program Options) error {
	return listOnRegionalAccounts[cumulus.DBCluster](
		program,
		&list.CommonList,
		cumulus.RegionalAccounts.DBClusters,
		"volume")
}
