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
