package cumulus

import (
	"context"
	"sync"
)

func collect[T Texter](ctx context.Context, from []Provider[T]) chan T {
	c := make(chan T)

	wg := sync.WaitGroup{}

	wg.Add(len(from))

	for _, p := range from {
		c1 := p(ctx)
		go func(from chan T) {
			defer wg.Done()
			for a := range from {
				c <- a
			}
		}(c1)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	return c
}
