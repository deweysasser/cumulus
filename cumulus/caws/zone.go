package caws

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/deweysasser/cumulus/cumulus"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"strings"
	"time"
)

func (a Account) Zones(ctx context.Context) chan cumulus.Zone {
	results := make(chan cumulus.Zone)

	go func() {
		defer close(results)
		s, e := a.session()
		if e != nil {
			cumulus.HandleError(ctx, e)
			close(results)
		}

		svc := route53.New(s)

		// TODO:  put in API call limiter
		start := time.Now()
		zones, e := svc.ListHostedZonesWithContext(ctx, &route53.ListHostedZonesInput{})
		CallTimer.Done(start)

		if e != nil {
			cumulus.HandleError(ctx, e)
			close(results)
		}

		for _, r := range zones.HostedZones {
			l := zerolog.Ctx(ctx).With().Str("zone_id", aws.StringValue(r.Id)).Logger()
			select {
			case <-ctx.Done():
				return
			case results <- zone{a, l.WithContext(ctx), r}:
			}
		}
	}()

	return results
}

type zone struct {
	Account
	context.Context
	*route53.HostedZone
}

func (z zone) Source() string {
	return z.Account.Name()
}

func (z zone) Ctx() context.Context {
	return z.Context
}

func (z zone) Id() cumulus.ID {
	return cumulus.ID(aws.StringValue(z.HostedZone.Id))
}

func (z zone) JSON() string {
	b, err := json.Marshal(z.HostedZone)

	if err != nil {
		log.Fatal().Err(err).Msg("Unable to marshal")
	}

	return string(b)
}

func (z zone) Text() string {

	zonetype := "public"

	if aws.BoolValue(z.HostedZone.Config.PrivateZone) {
		zonetype = "private"
	}

	return strings.Join([]string{
		aws.StringValue(z.HostedZone.Id),
		aws.StringValue(z.HostedZone.Name),
		zonetype,
	}, "\t")
	//TODO implement me
	panic("implement me")
}
