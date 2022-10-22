package program

import (
	"context"
	"github.com/rs/zerolog/log"
	"golang.org/x/time/rate"
	"sync"
	"time"
)

type adjustment int

const (
	up adjustment = iota
	down
)

type AdaptiveLimit struct {
	*rate.Limiter
	MaxRetries                       int
	requiredSuccesses                int64
	successesSeen                    int64
	lock                             sync.Locker
	adjustments                      chan adjustment
	downwardCooldown, upwardCooldown time.Duration
	history                          RingArray
}

func NewAdaptiveLimit(ctx context.Context, limit *rate.Limiter) *AdaptiveLimit {
	a := &AdaptiveLimit{
		Limiter:           limit,
		MaxRetries:        5,
		requiredSuccesses: 10,
		successesSeen:     0,
		lock:              &sync.Mutex{},
		adjustments:       make(chan adjustment),
		downwardCooldown:  10 * time.Second,
		upwardCooldown:    1 * time.Second,
		history:           RingArray{Values: make([]int, 0, 20)},
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case adjust := <-a.adjustments:
				a.lock.Lock()
				switch adjust {
				case up:
					//a.success()
				case down:
					//a.decrease()
				}
				a.lock.Unlock()
			}
		}
	}()

	return a
}

func (a *AdaptiveLimit) adjust(ad adjustment) {
	select {
	case a.adjustments <- ad:
	default:
	}
}

func (a *AdaptiveLimit) Do(ctx context.Context, w Worker) error {
	var err error
	for i := 0; i < a.MaxRetries; i++ {
		if err := a.Wait(ctx); err != nil {
			return err
		}

		if err := ctx.Err(); err != nil {
			return err
		}

		if err = w(); err != nil {
			log.Debug().Err(err).Int("try_number", i).Msg("try failed.  retrying")
			a.adjust(down)
		} else {
			a.adjust(up)
			return nil
		}
	}

	log.Debug().Err(err).Msg("worker failed for the last time.  giving up")

	return err
}

func (a *AdaptiveLimit) success() {
	//if a.successesSeen > a.requiredSuccesses {

	limit := a.Limit() + 1
	log.Warn().Int("new_limit", int(limit)).Msg("increasing rate limit")
	a.SetLimit(limit)

	a.successesSeen = 0
	time.Sleep(a.upwardCooldown)
	//} else {
	//	a.successesSeen++
	//}
}

func (a *AdaptiveLimit) decrease() {
	limit := a.Limit()

	a.history.Add(int(limit))

	limit = rate.Limit(a.history.Average() * 90 / 100)

	if limit > 0 {
		a.SetLimit(limit)
	}
	log.Warn().Int("new_limit", int(limit)).Msg("Decreasing rate limit")

	a.successesSeen = 0

	time.Sleep(a.downwardCooldown)

}
