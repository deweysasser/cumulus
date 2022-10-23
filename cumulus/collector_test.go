package cumulus

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCollect(t *testing.T) {

	expected := map[string]bool{
		"one":   true,
		"two":   true,
		"three": true,
		"four":  true,
	}

	p := func(ctx context.Context) chan string {
		return arrayIterator([]string{"one"})
	}
	p2 := func(ctx context.Context) chan string {
		return arrayIterator([]string{"two"})
	}
	p3 := func(ctx context.Context) chan string {
		return arrayIterator([]string{"three", "four"})
	}

	r := collect(context.Background(), []Provider[string]{p, p2, p3})

	count := 0
	for v := range r {
		assert.True(t, expected[v])
		count++
	}

	assert.Equal(t, 4, count)

}

func arrayIterator[T any](items []T) chan T {
	c := make(chan T)
	go func() {
		for _, t := range items {
			c <- t
		}
		close(c)
	}()

	return c
}
