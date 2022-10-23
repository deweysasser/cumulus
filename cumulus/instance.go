package cumulus

import "context"

type InstanceVisitor func(ctx context.Context, cancel context.CancelFunc, instance Instance) error

type VisitInstancer interface {
	VisitInstance(ctx context.Context, cancel context.CancelFunc, visitor InstanceVisitor) error
}
