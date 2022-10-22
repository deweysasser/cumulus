package caws

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/deweysasser/golang-program/cumulus"
	"github.com/hashicorp/go-multierror"
	"github.com/rs/zerolog/log"
	"strings"
	"time"
)

func (a RegionalAccount) VisitInstance(ctx context.Context, cancel context.CancelFunc, visitor cumulus.InstanceVisitor) error {

	s, e := a.session()
	if e != nil {
		return e
	}

	fmt.Sprint("Logger should print next")

	svc := ec2.New(s)

	a.Read.Wait(ctx)
	start := time.Now()
	instances, err := svc.DescribeInstancesWithContext(ctx, &ec2.DescribeInstancesInput{})
	CallTimer.Done(start)

	if err != nil {
		return e
	}

	for _, r := range instances.Reservations {
		for _, i := range r.Instances {
			select {
			case <-ctx.Done():
				return err
			default:
				l := log.Ctx(ctx).With().Str("instance_id", aws.StringValue(i.InstanceId)).Logger()
				err = multierror.Append(err,
					visitor(
						l.WithContext(ctx),
						cancel,
						instance{i}))
			}
		}
	}
	return err
}

type instance struct{ *ec2.Instance }

func (i instance) Id() cumulus.ID {
	return cumulus.ID(aws.StringValue(i.InstanceId))
}

func (i instance) JSON() string {
	bytes, err := json.Marshal(i)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to marshal")
	}

	return string(bytes)
}

func (i instance) Text() string {

	name := ""
	for _, t := range i.Tags {
		if aws.StringValue(t.Key) == "Name" {
			name = aws.StringValue(t.Value)
			break
		}
	}

	return strings.Join([]string{
		fieldValue(i.InstanceId),
		fieldValue(i.InstanceType),
		fieldValue(i.PrivateDnsName),
		fieldValue(i.PrivateIpAddress),
		fieldValue(i.PublicDnsName),
		fieldValue(i.PublicIpAddress),
		name,
	},
		"\t",
	)
}

func fieldValue(s *string) string {
	if s == nil || *s == "" {
		return "-"
	}
	return *s
}
