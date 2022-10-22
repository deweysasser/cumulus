package stats

import (
	"context"
	"github.com/rs/zerolog/log"
	"time"
)

type Counter struct {
	Name         string
	count, Total int
	c            chan operation
}

func NewCounter(ctx context.Context, name string) *Counter {
	c := &Counter{
		Name:  name,
		count: 0,
		c:     make(chan operation, 100),
	}

	ticker := time.NewTicker(period)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-c.c:
				c.count++
				c.Total++
			case <-ticker.C:
				c.Report()
				c.count = 0
			}
		}
	}()

	reportersChannel <- c

	return c
}

func (c *Counter) Report() {
	rate := c.Rate()
	log.Info().
		Str("name", c.Name).
		Float64("rate", rate).
		Int("count", c.count).
		Int("Total", c.Total).
		Msg("counter")
}

func (c *Counter) Inc() {
	c.c <- inc
}

func (c Counter) Rate() float64 {
	return float64(c.count) * float64(1*time.Second) / float64(period)

}
