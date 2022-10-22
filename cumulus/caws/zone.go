package caws

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/deweysasser/golang-program/cumulus"
	"github.com/hashicorp/go-multierror"
	"github.com/rs/zerolog/log"
	"strings"
	"time"
)

func (a Account) VisitZones(ctx context.Context, cancel context.CancelFunc, visitor cumulus.ZoneVisitor) error {

	s, e := a.session()
	if e != nil {
		return e
	}

	fmt.Sprint("Logger should print next")

	svc := route53.New(s)

	start := time.Now()
	zones, err := svc.ListHostedZonesWithContext(ctx, &route53.ListHostedZonesInput{})
	CallTimer.Done(start)

	if err != nil {
		return e
	}

	for _, r := range zones.HostedZones {
		select {
		case <-ctx.Done():
			return err
		default:
			err = multierror.Append(err, visitor(ctx, cancel, zone{r}))
		}
	}
	return err
}

type zone struct {
	*route53.HostedZone
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
