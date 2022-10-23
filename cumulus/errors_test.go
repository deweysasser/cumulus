package cumulus

import (
	"context"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrorContexts(t *testing.T) {

	ctx := context.Background()

	assert.NotNil(t, ErrorContext(ctx))

	called := false
	ctx = WithErrorHandler(ctx, func(ctx context.Context, err error) {
		called = true
	})

	HandleError(ctx, errors.New("testing"))

	assert.True(t, called, "Error handler not called")

}
