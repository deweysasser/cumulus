package caws

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/deweysasser/cumulus/cumulus"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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
			case results <- hostedzone{a, l.WithContext(ctx), r}:
			}
		}
	}()

	return results
}

type hostedzone struct {
	account Account
	context.Context
	obj *route53.HostedZone
}

func (z hostedzone) Source() cumulus.Fielder {
	return z.account
}

func (z hostedzone) Ctx() context.Context {
	return z.Context
}

func (z hostedzone) Id() cumulus.ID {
	return cumulus.ID(aws.StringValue(z.obj.Id))
}

func (z hostedzone) JSON() string {
	b, err := json.Marshal(z.obj)

	if err != nil {
		log.Fatal().Err(err).Msg("Unable to marshal")
	}

	return string(b)
}

func (z hostedzone) GetFields(builder cumulus.IFieldBuilder) {

	z.GeneratedFields(builder)
	if aws.BoolValue(z.obj.Config.PrivateZone) {
		builder.What("type", "private")
	} else {
		builder.What("type", "public")
	}
}
