package cumulus

import "context"

// WARNING:  this file is generated. DO NOT EDIT.  Edit the source instead

func (a Accounts) AccountInfos(ctx context.Context) chan AccountInfo {
	var providers []Provider[AccountInfo]

	for _, acct := range a {

		if v, ok := acct.(AccountInfoer); ok {
			providers = append(providers, v.AccountInfos)
		}
	}

	return collect(ctx, providers)
}

func (a Accounts) Zones(ctx context.Context) chan Zone {
	var providers []Provider[Zone]

	for _, acct := range a {

		if v, ok := acct.(Zoner); ok {
			providers = append(providers, v.Zones)
		}
	}

	return collect(ctx, providers)
}

func (a Accounts) NameRecords(ctx context.Context) chan NameRecord {
	var providers []Provider[NameRecord]

	for _, acct := range a {

		if v, ok := acct.(NameRecorder); ok {
			providers = append(providers, v.NameRecords)
		}
	}

	return collect(ctx, providers)
}
