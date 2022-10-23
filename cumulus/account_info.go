package cumulus

import "context"

type AccountInfoVisitor func(ctx context.Context, info AccountInfo) error

type VisitAccountInfoer interface {
	VisitAccountInfo(ctx context.Context, visitor AccountInfoVisitor) error
}
