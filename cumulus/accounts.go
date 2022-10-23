package cumulus

import (
	"context"
	"github.com/rs/zerolog"
	"strings"
)

func (a Accounts) String() string {
	s := make([]string, len(a))
	for n, ac := range a {
		s[n] = ac.String()
	}

	return strings.Join(s, ", ")
}

//type AccountProducer func() Accounts
//
//var AccountProducers []AccountProducer

func (a Accounts) InRegion(region ...string) RegionalAccounts {

	var list RegionalAccounts

	for _, acct := range a {
		for _, r := range region {
			list = append(list, acct.InRegion(r))
		}
	}

	return list
}

func (a Accounts) Unique(ctx context.Context) Accounts {
	// TODO:  evaluate uniqueness in order so the set statys consistent with consistent order of the credentials file.
	// Still do the lookup concurrently though.
	seen := make(map[string]bool)
	var unique Accounts

	log := zerolog.Ctx(ctx)

	log.Debug().Msg("finding unique accounts")

	for info := range a.AccountInfos(ctx) {
		//log := zerolog.Ctx(info.Ctx())
		if !seen[info.ID()] {
			seen[info.ID()] = true
			unique = append(unique, info.Account())
			//log.Debug().Str("id", info.ID()).Str("source", info.Source()).Msg("unique: using account")
		} else {
			//log.Debug().Str("id", info.ID()).Str("source", info.Source()).Msg("unique: ignoring duplicate account")
		}
	}

	return unique
}
