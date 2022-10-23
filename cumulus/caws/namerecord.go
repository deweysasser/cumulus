package caws

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/deweysasser/cumulus/cumulus"
	"github.com/rs/zerolog"
	"strings"
	"sync"
	"time"
)

func (a Account) NameRecords(ctx context.Context) chan cumulus.NameRecord {
	results := make(chan cumulus.NameRecord)

	s, e := a.session()
	if e != nil {
		cumulus.HandleError(ctx, e)
		close(results)
	}

	wg := sync.WaitGroup{}

	for zone := range a.Zones(ctx) {
		wg.Add(1)

		go func(zone cumulus.Zone) {
			defer wg.Done()
			svc := route53.New(s)

			// TODO:  put in API call limiter
			start := time.Now()
			// TODO:  call the "pages" variation, or handle page iteration ourselves
			records, e := svc.ListResourceRecordSets(&route53.ListResourceRecordSetsInput{HostedZoneId: aws.String(string(zone.Id()))})
			CallTimer.Done(start)

			if e != nil {
				cumulus.HandleError(ctx, e)
				close(results)
			}

			for _, r := range records.ResourceRecordSets {
				l := zerolog.Ctx(zone.Ctx()).With().Str("record_name", aws.StringValue(r.Name)).Logger()

				select {
				case <-ctx.Done():
					cumulus.HandleError(ctx, ctx.Err())
				case results <- &resourceRecord{a, zone, l.WithContext(ctx), r}:

				}
			}
		}(zone)
	}

	go func() {
		defer close(results)
		wg.Wait()
	}()

	return results
}

type resourceRecord struct {
	Account
	cumulus.Zone
	context.Context
	*route53.ResourceRecordSet
}

func (r resourceRecord) Text() string {
	return strings.Join([]string{
		aws.StringValue(r.ResourceRecordSet.Name),
		aws.StringValue(r.ResourceRecordSet.Type),
		r.Zone.Id().String(),
	}, "\t")
}

func (r resourceRecord) GetFields(builder cumulus.IFieldBuilder) {
	builder.
		Name(aws.StringValue(r.ResourceRecordSet.Name)).
		What("record_type", aws.StringValue(r.ResourceRecordSet.Type)).
		Where("zone_id", r.Zone.Id().String()).
		Where("account", r.Account.Name()).
		Done()
}

func (r resourceRecord) Ctx() context.Context {
	return r.Context
}

func (r resourceRecord) Source() cumulus.Fielder {
	return r.Account
}

func (r resourceRecord) Id() cumulus.ID {
	return cumulus.ID(aws.StringValue(r.ResourceRecordSet.Name))
}
