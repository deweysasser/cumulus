package cumulus

import "context"

// WARNING:  this file is generated. DO NOT EDIT.  Edit the source instead

func (a RegionalAccounts) Instances(ctx context.Context) chan Instance {
	var providers []Provider[Instance]

	for _, acct := range a {

		if v, ok := acct.(Instancer); ok {
			providers = append(providers, v.Instances)
		}
	}

	return collect(ctx, providers)
}

func (a RegionalAccounts) Snapshots(ctx context.Context) chan Snapshot {
	var providers []Provider[Snapshot]

	for _, acct := range a {

		if v, ok := acct.(Snapshoter); ok {
			providers = append(providers, v.Snapshots)
		}
	}

	return collect(ctx, providers)
}
