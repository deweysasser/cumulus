package caws

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/deweysasser/cumulus/cumulus"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"time"
)

func (a RegionalAccount) Instances(ctx context.Context) chan cumulus.Instance {

	result := make(chan cumulus.Instance)

	s, e := a.session()
	if e != nil {
		//return e
		close(result)
		return result
	}

	svc := ec2.New(s)

	a.Read.Wait(ctx)
	start := time.Now()
	instances, err := svc.DescribeInstancesWithContext(ctx, &ec2.DescribeInstancesInput{})
	CallTimer.Done(start)

	if err != nil {
		//return e
		close(result)
		return result
	}

	go func() {
		defer close(result)

		for _, r := range instances.Reservations {
			for _, i := range r.Instances {
				select {
				case <-ctx.Done():
					return
				default:
					l := log.Ctx(ctx).With().Str("instance_id", aws.StringValue(i.InstanceId)).Logger()
					ctx := l.WithContext(ctx)
					result <- &instance{a, ctx, i}
				}
			}
		}
	}()
	return result
}

type instance struct {
	account RegionalAccount
	ctx     context.Context
	obj     *ec2.Instance
}

func (i instance) Ctx() context.Context {
	return i.ctx
}

func (i instance) Source() cumulus.Fielder {
	return i.account
}

func (i instance) Id() cumulus.ID {
	return cumulus.ID(aws.StringValue(i.obj.InstanceId))
}

func (i instance) JSON() string {
	bytes, err := json.Marshal(i)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to marshal")
	}

	return string(bytes)
}

func (i instance) GetFields(builder cumulus.IFieldBuilder) {
	i.GeneratedFields(builder)
}

func addNetworkInterfaces(builder cumulus.IFieldBuilder, ctx context.Context, inf []*ec2.InstanceNetworkInterface) {
	for _, inf := range inf {
		zerolog.Ctx(ctx).Debug().Msg("found additional network interface")
		for _, ip := range inf.PrivateIpAddresses {
			builder.Where("private_dns_additional", aws.StringValue(ip.PrivateDnsName), cumulus.DefaultHidden)
			builder.Where("private_ip_additional", aws.StringValue(ip.PrivateIpAddress), cumulus.DefaultHidden)
		}
	}
}

func ec2_Tag_to_fields(builder cumulus.IFieldBuilder, _ context.Context, tags []*ec2.Tag) {
	for _, t := range tags {
		if aws.StringValue(t.Key) == "Name" {
			builder.Name(aws.StringValue(t.Value))
		} else {
			builder.Tag(aws.StringValue(t.Key), aws.StringValue(t.Value))
		}
	}
}
