package cumulus

import (
	"context"
	"errors"
	"github.com/hashicorp/go-multierror"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"strings"
	"sync"
)

type Accounts []Account

func (a Accounts) String() string {
	s := make([]string, len(a))
	for n, ac := range a {
		s[n] = ac.String()
	}

	return strings.Join(s, ", ")
}

//type AccountProducer func() Accounts
//
//var AccountProducers []AccountProducer

func (a Accounts) InRegion(region ...string) RegionalAccounts {

	var list RegionalAccounts

	for _, acct := range a {
		for _, r := range region {
			list = append(list, acct.InRegion(r))
		}
	}

	return list
}

func (a Accounts) Unique() Accounts {
	type combined struct {
		Account
		id string
	}
	c := make(chan combined)
	wg := sync.WaitGroup{}

	wg.Add(1)

	go func() {
		defer close(c)
		log.Debug().Msg("Waiting on info channel close")
		wg.Wait()
	}()

	go func() {
		defer wg.Done()

		a.VisitAccountInfo(context.Background(), func(ctx context.Context, info AccountInfo) error {
			log.Ctx(ctx).Debug().Str("id", info.ID()).Msg("reporting information")
			ra := AccountContext(ctx)
			c <- combined{ra, info.ID()}
			return nil
		})
	}()

	already := make(map[string]bool, len(a))
	results := make(Accounts, 0)

	for i := range c {
		if _, seen := already[i.id]; !seen {
			already[i.id] = true
			results = append(results, i.Account)
			log.Debug().Str("id", i.id).Str("account", i.Account.Name()).Msg("using account")
		} else {
			log.Debug().Str("id", i.id).Str("account", i.Account.Name()).Msg("ignoring duplicate account")
		}
	}

	log.Debug().Str("unique accounts", results.String()).Msg("Unique Accounts")

	return results
}

type RegionalAccounts []RegionalAccount

func (a Accounts) VisitAccountInfo(ctx context.Context, visitor AccountInfoVisitor) error {
	wg := sync.WaitGroup{}
	results := make(chan error, 100)

	for _, acct := range a {
		if v, implements := acct.(VisitAccountInfoer); implements {
			wg.Add(1)
			logger := log.With().Str("account", acct.Name()).Logger()

			go func(v VisitAccountInfoer, logger zerolog.Logger, acct Account) {
				ctx := logger.WithContext(withAccountContext(ctx, acct))

				defer wg.Done()
				select {
				case <-ctx.Done():
					return
				default:
					if err := v.VisitAccountInfo(ctx, visitor); err != nil {
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

	var list error

	for e := range results {
		list = multierror.Append(list, e)
	}

	return nil
}

func (a Accounts) VisitZones(ctx context.Context, visitor ZoneVisitor) error {
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
		logger := zerolog.Ctx(ctx).With().Str("account", acct.Name()).Logger()

		if v, implements := acct.(VisitZoneser); implements {
			wg.Add(1)
			foundAny = true
			go func(v VisitZoneser, logger zerolog.Logger, acct Account) {

				ctx := logger.WithContext(withAccountContext(ctx, acct))

				defer wg.Done()
				select {
				case <-ctx.Done():
					logger.Debug().Msg("Context canceled")
					return
				default:
					if err := v.VisitZones(ctx, cancel, visitor); err != nil {
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
		return errors.New("no accounts implement VisitInstance")
	}

	var list error

	for e := range results {
		list = multierror.Append(list, e)
	}

	return nil
}
