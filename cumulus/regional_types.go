package cumulus

import "context"

type Instance interface {
	Common
	Id() ID
}

type Snapshot interface {
	Common
	Id() ID
	Delete(ctx context.Context, dryRun bool) error
}

type MachineImage interface {
	Common
}

type Volume interface {
	Common
}

type Subscription interface {
	Common
}

type Topic interface {
	Common
}

type DBCluster interface {
	Common
}

type DBInstance interface {
	Common
}
