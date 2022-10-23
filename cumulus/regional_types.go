package cumulus

import "context"

type Instance interface {
	Common
	Id() ID
	Fields() []Field
}

type Snapshot interface {
	Common
	Id() ID
	Delete(ctx context.Context, dryRun bool) error
}
