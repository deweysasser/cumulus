package caws

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/deweysasser/cumulus/cumulus"
	"github.com/rs/zerolog/log"
	"time"
)

func (a RegionalAccount) Volumes(ctx context.Context) chan cumulus.Volume {

	result := make(chan cumulus.Volume)

	s, e := a.session()
	if e != nil {
		//return e
		close(result)
		return result
	}

	start := time.Now()

	svc := ec2.New(s)

	if e = a.Read.Wait(ctx); e != nil {
		cumulus.HandleError(ctx, e)
		close(result)
		return result
	}

	out, err := svc.DescribeVolumesWithContext(ctx, &ec2.DescribeVolumesInput{})
	CallTimer.Done(start)

	if err != nil {
		close(result)
		return result
	}

	go func() {
		defer close(result)

		for _, r := range out.Volumes {
			select {
			case <-ctx.Done():
				return
			default:
				l := log.Ctx(ctx).With().Str("volume_id", aws.StringValue(r.VolumeId)).Logger()
				ctx := l.WithContext(ctx)
				result <- &volume{a, ctx, r}
			}
		}
	}()
	return result
}

type volume struct {
	RegionalAccount
	context.Context
	obj *ec2.Volume
}

func (a volume) Ctx() context.Context {
	return a.Context
}

func (a volume) GetFields(builder cumulus.IFieldBuilder) {

	a.GeneratedFields(builder)

	for _, a := range a.obj.Attachments {
		builder.Where("attached_to", aws.StringValue(a.InstanceId))
	}
}
