package stats

import (
	"context"
	"github.com/rs/zerolog/log"
	"time"
)

type Timer struct {
	Name                    string
	count, totalCalls       int
	duration, totalDuration time.Duration
	periodStart             time.Time
	c                       chan time.Duration
}

func NewTimer(ctx context.Context, name string) *Timer {
	c := &Timer{
		Name:        name,
		count:       0,
		periodStart: time.Now(),
		c:           make(chan time.Duration, 100),
	}

	ticker := time.NewTicker(period)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case dur := <-c.c:
				c.count++
				c.totalCalls++
				c.totalDuration = c.totalDuration + dur
				c.duration = c.duration + dur
			case <-ticker.C:
				c.Report()
				c.count = 0
				c.duration = 0
				c.periodStart = time.Now()
			}
		}
	}()

	reportersChannel <- c
	return c
}

type Doner interface {
	Done()
}

type doner struct {
	start time.Time
	timer *Timer
}

func (d doner) Done() {
	d.timer.Done(d.start)
}

func (c *Timer) Done(start time.Time) {
	c.c <- time.Since(start)
}

func (c *Timer) Call() Doner {
	return doner{
		start: time.Now(),
		timer: c,
	}
}

func (c *Timer) Report() {
	log.Info().
		Str("name", c.Name).
		Float64("rate", c.Rate()).
		Dur("average_latency", c.AverageLatency()).
		Int("count", c.count).
		Msg("timer")
}

func (c *Timer) Total() {
	log.Info().
		Str("name", c.Name).
		Int("TotalCalls", c.totalCalls).
		Dur("TotalDuration", c.totalDuration).
		Msg("timer totals")
}

func (c Timer) Rate() float64 {
	return float64(c.count) * float64(1*time.Second) / float64(time.Since(c.periodStart))
}

func (c Timer) AverageLatency() time.Duration {
	if c.count > 1 {
		return c.duration / time.Duration(c.count)
	}
	return 0
}
