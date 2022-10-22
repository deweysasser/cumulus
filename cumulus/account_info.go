package cumulus

import "context"

type AccountInfo interface {
	Name() string
	ID() string
}

type AccountInfoVisitor func(ctx context.Context, info AccountInfo) error

type VisitAccountInfoer interface {
	VisitAccountInfo(ctx context.Context, visitor AccountInfoVisitor) error
}
