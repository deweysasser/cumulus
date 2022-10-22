package cumulus

import (
	"context"
	"errors"
	"fmt"
	"github.com/hashicorp/go-multierror"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"sync"
)

var accountCtxKey struct{}

func withAccountContext(ctx context.Context, a Account) context.Context {
	return context.WithValue(ctx, accountCtxKey, a)
}

func AccountContext(ctx context.Context) Account {
	value := ctx.Value(accountCtxKey)
	if a, yes := value.(Account); yes {
		return a
	} else {
		log.Fatal().Str("value", fmt.Sprint(value)).Msg("value is not an Account")
		return nil
	}
}

func RegionalAccountContext(ctx context.Context) RegionalAccount {
	return ctx.Value(accountCtxKey).(RegionalAccount)
}

func (a RegionalAccounts) VisitInstance(ctx context.Context, visitor InstanceVisitor) error {
	wg := sync.WaitGroup{}
	results := make(chan error, 100)

	ctx, ctxCancel := context.WithCancel(ctx)

	cancel := func() {
		log.Debug().Msg("visitor cancel called")
		ctxCancel()
	}

	defer ctxCancel()

	foundAny := false
	for _, acct := range a {
		logger := zerolog.Ctx(ctx).With().Str("account", acct.Name()).Str("region", acct.Region()).Logger()

		if v, implements := acct.(VisitInstancer); implements {
			wg.Add(1)
			foundAny = true
			go func(v VisitInstancer, logger zerolog.Logger) {

				ctx := logger.WithContext(withAccountContext(ctx, acct))

				defer wg.Done()
				select {
				case <-ctx.Done():
					logger.Debug().Msg("Context canceled")
					return
				default:
					if err := v.VisitInstance(ctx, cancel, visitor); err != nil {
						logger.Debug().Err(err).Msg("visitor return error")
						results <- err
					}
				}
			}(v, logger)
		}
	}
	go func() {
		defer close(results)
		wg.Wait()
	}()

	if !foundAny {
		return errors.New("no accounts implement VisitInstance")
	}

	var list error

	for e := range results {
		list = multierror.Append(list, e)
	}

	return nil
}

func (a RegionalAccounts) VisitSnapshot(ctx context.Context, visitor SnapshotVisitor) error {
	wg := sync.WaitGroup{}
	results := make(chan error, 100)

	// TODO:  remove this
	cancel := func() {}

	foundAny := false
	for _, acct := range a {
		logger := log.Logger.With().Str("account", acct.Name()).Str("region", acct.Region()).Logger()

		if v, implements := acct.(VisitSnapshotr); implements {
			wg.Add(1)
			foundAny = true
			go func(v VisitSnapshotr, logger zerolog.Logger, acct RegionalAccount) {

				ctx := logger.WithContext(withAccountContext(ctx, acct))

				defer wg.Done()
				select {
				case <-ctx.Done():
					logger.Debug().Msg("Context canceled")
					return
				default:
					if err := v.VisitSnapshot(ctx, cancel, visitor); err != nil {
						logger.Debug().Err(err).Msg("visitor return error")
						results <- err
					}
				}
			}(v, logger, acct)
		}
	}
	go func() {
		defer close(results)
		wg.Wait()
	}()

	if !foundAny {
		return errors.New("no accounts implement VisitSnapshot")
	}

	var list error

	for e := range results {
		list = multierror.Append(list, e)
	}

	return nil
}
