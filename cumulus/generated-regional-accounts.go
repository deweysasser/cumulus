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

func (a RegionalAccounts) MachineImages(ctx context.Context) chan MachineImage {
	var providers []Provider[MachineImage]

	for _, acct := range a {

		if v, ok := acct.(MachineImager); ok {
			providers = append(providers, v.MachineImages)
		}
	}

	return collect(ctx, providers)
}

func (a RegionalAccounts) Volumes(ctx context.Context) chan Volume {
	var providers []Provider[Volume]

	for _, acct := range a {

		if v, ok := acct.(Volumer); ok {
			providers = append(providers, v.Volumes)
		}
	}

	return collect(ctx, providers)
}

func (a RegionalAccounts) Subscriptions(ctx context.Context) chan Subscription {
	var providers []Provider[Subscription]

	for _, acct := range a {

		if v, ok := acct.(Subscriptioner); ok {
			providers = append(providers, v.Subscriptions)
		}
	}

	return collect(ctx, providers)
}

func (a RegionalAccounts) Topics(ctx context.Context) chan Topic {
	var providers []Provider[Topic]

	for _, acct := range a {

		if v, ok := acct.(Topicer); ok {
			providers = append(providers, v.Topics)
		}
	}

	return collect(ctx, providers)
}

func (a RegionalAccounts) DBClusters(ctx context.Context) chan DBCluster {
	var providers []Provider[DBCluster]

	for _, acct := range a {

		if v, ok := acct.(DBClusterer); ok {
			providers = append(providers, v.DBClusters)
		}
	}

	return collect(ctx, providers)
}

func (a RegionalAccounts) DBInstances(ctx context.Context) chan DBInstance {
	var providers []Provider[DBInstance]

	for _, acct := range a {

		if v, ok := acct.(DBInstancer); ok {
			providers = append(providers, v.DBInstances)
		}
	}

	return collect(ctx, providers)
}
