package program

import "github.com/deweysasser/cumulus/cumulus"

type TopicList struct {
	CommonList `embed:""`
}

func (list *TopicList) Run(program Options) error {
	return listOnRegionalAccounts[cumulus.Topic](
		program,
		&list.CommonList,
		cumulus.RegionalAccounts.Topics,
		"Topic")
}

type SubscriptionList struct {
	CommonList `embed:""`
}

func (list *SubscriptionList) Run(program Options) error {
	return listOnRegionalAccounts[cumulus.Subscription](
		program,
		&list.CommonList,
		cumulus.RegionalAccounts.Subscriptions,
		"Subscription")
}
