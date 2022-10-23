package stats

import (
	"context"
	"github.com/rs/zerolog/log"
	"time"
)

type Counter struct {
	Name              string
	count, TotalCalls int
	periodStart       time.Time
	c                 chan operation
}

func NewCounter(ctx context.Context, name string) *Counter {
	c := &Counter{
		Name:        name,
		count:       0,
		periodStart: time.Now(),
		c:           make(chan operation, 100),
	}

	ticker := time.NewTicker(period)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-c.c:
				c.count++
				c.TotalCalls++
			case <-ticker.C:
				c.Report()
				c.count = 0
				c.periodStart = time.Now()
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
		Msg("counter")
}

func (c *Counter) Total() {
	log.Info().
		Str("name", c.Name).
		Int("TotalCalls", c.TotalCalls).
		Msg("counter totals")
}

func (c *Counter) Inc() {
	c.c <- inc
}

func (c Counter) Rate() float64 {
	return float64(c.count) * float64(1*time.Second) / float64(time.Since(c.periodStart))
}
