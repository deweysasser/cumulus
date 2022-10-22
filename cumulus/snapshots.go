package cumulus

import "context"

type Snapshot interface {
	Id() ID
	JSON() string
	Text() string
	Delete(ctx context.Context, dryRun bool) error
}

type SnapshotVisitor func(ctx context.Context, cancel context.CancelFunc, instance Snapshot) error

type VisitSnapshotr interface {
	VisitSnapshot(ctx context.Context, cancel context.CancelFunc, visitor SnapshotVisitor) error
}
