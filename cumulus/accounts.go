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

func (a Accounts) AccountInfos(ctx context.Context) chan AccountInfo {
	var providers []Provider[AccountInfo]

	for _, acct := range a {
		providers = append(providers, acct.AccountInfos)
	}

	zerolog.Ctx(ctx).Debug().Msg("Collecting")
	return collect(ctx, providers)
}

func (a Accounts) Unique(ctx context.Context) Accounts {

	seen := make(map[string]bool)
	var unique Accounts

	log := zerolog.Ctx(ctx)

	log.Debug().Msg("finding unique accounts")

	for info := range a.AccountInfos(ctx) {
		log.Debug().Msg("checking")
		log := zerolog.Ctx(info.Ctx())
		log.Debug().
			Msg("Checking account")
		if !seen[info.ID()] {
			seen[info.ID()] = true
			unique = append(unique, info.Source())
			log.Debug().Str("id", info.ID()).Str("account", info.Source().Name()).Msg("using account")
		} else {
			log.Debug().Str("id", info.ID()).Str("account", info.Source().Name()).Msg("ignoring duplicate account")
		}
	}

	return unique
}

func (a Accounts) VisitAccountInfo(ctx context.Context, visitor AccountInfoVisitor) error {
	wg := sync.WaitGroup{}
	results := make(chan error, 100)

	for _, acct := range a {
		if v, implements := acct.(VisitAccountInfoer); implements {
			wg.Add(1)
			logger := log.With().Str("account", acct.Name()).Logger()

			go func(v VisitAccountInfoer, logger zerolog.Logger, acct Account) {
				//ctx := logger.WithContext(withAccountContext(ctx, acct))

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

				//ctx := logger.WithContext(withAccountContext(ctx, acct))

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
