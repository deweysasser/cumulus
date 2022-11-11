package caws

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/deweysasser/cumulus/cumulus"
	"github.com/rs/zerolog"
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
				case results <- &resourcerecordset{a, zone, l.WithContext(ctx), r}:

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

type resourcerecordset struct {
	Account
	cumulus.Zone
	context.Context
	obj *route53.ResourceRecordSet
}

func (r resourcerecordset) GetFields(builder cumulus.IFieldBuilder) {
	r.GeneratedFields(builder)
	builder.
		Where("zone_id", r.Zone.Id().String()).
		Done()

	for _, record := range r.obj.ResourceRecords {
		builder.What("record", aws.StringValue(record.Value), cumulus.DefaultHidden)
	}
}

func (r resourcerecordset) Ctx() context.Context {
	return r.Context
}

func (r resourcerecordset) Source() cumulus.Fielder {
	return r.Account
}

func (r resourcerecordset) Id() cumulus.ID {
	return cumulus.ID(aws.StringValue(r.obj.Name))
}
