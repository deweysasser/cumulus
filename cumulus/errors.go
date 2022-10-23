package cumulus

import (
	"context"
	"github.com/hashicorp/go-multierror"
	"github.com/rs/zerolog"
)

// IgnoreErrors ignores all errors.  If it returns true, procesing should be aborted
func IgnoreErrors(ctx context.Context, err error) {}
func LogErrors(ctx context.Context, err error) {
	zerolog.Ctx(ctx).Error().Err(err).Msg("error during operation")
}

type contextErrorKey struct{}

func WithErrorHandler(ctx context.Context, handler ErrorHandler) context.Context {
	return context.WithValue(ctx, contextErrorKey{}, handler)
}

func ErrorContext(ctx context.Context) ErrorHandler {
	h := ctx.Value(contextErrorKey{})
	if h == nil {
		return LogErrors
	}
	return h.(ErrorHandler)
}

func HandleError(ctx context.Context, err error) {
	ErrorContext(ctx)(ctx, err)
}

///////////////////////////////////////////////////

// ErrorCollector uses multierror to accumulate errors into a large list
type ErrorCollector struct {
	Error error
}

func (e *ErrorCollector) Handle(ctx context.Context, err error) {
	e.Error = multierror.Append(e.Error, err)
}
