package cumulus

import (
	"context"
)

func (a RegionalAccounts) AccountInfos(ctx context.Context) chan AccountInfo {
	var providers []Provider[AccountInfo]

	for _, acct := range a {

		if v, ok := acct.(AccountInfoer); ok {
			providers = append(providers, v.AccountInfos)
		}
	}

	return collect(ctx, providers)
}

func (a RegionalAccounts) Instances(ctx context.Context) chan Instance {
	var providers []Provider[Instance]

	for _, acct := range a {

		if v, ok := acct.(Instancer); ok {
			providers = append(providers, v.Instances)
		}
	}

	return collect(ctx, providers)
}
