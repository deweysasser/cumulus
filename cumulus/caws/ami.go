package caws

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/deweysasser/cumulus/cumulus"
	"github.com/rs/zerolog/log"
	"time"
)

func (a RegionalAccount) MachineImages(ctx context.Context) chan cumulus.MachineImage {

	result := make(chan cumulus.MachineImage)

	s, e := a.session()
	if e != nil {
		//return e
		close(result)
		return result
	}

	stssvc := sts.New(s)

	if e = a.Read.Wait(ctx); e != nil {
		cumulus.HandleError(ctx, e)
		close(result)
		return result
	}
	start := time.Now()
	stsout, e := stssvc.GetCallerIdentity(&sts.GetCallerIdentityInput{})
	CallTimer.Done(start)

	if e != nil {
		cumulus.HandleError(ctx, e)
		close(result)
		return result
	}

	svc := ec2.New(s)

	if e = a.Read.Wait(ctx); e != nil {
		cumulus.HandleError(ctx, e)
		close(result)
		return result
	}

	out, err := svc.DescribeImagesWithContext(ctx, &ec2.DescribeImagesInput{
		Owners: []*string{
			stsout.Account,
		},
	})
	CallTimer.Done(start)

	if err != nil {
		close(result)
		return result
	}

	go func() {
		defer close(result)

		for _, r := range out.Images {
			select {
			case <-ctx.Done():
				return
			default:
				l := log.Ctx(ctx).With().Str("image_id", aws.StringValue(r.ImageId)).Logger()
				ctx := l.WithContext(ctx)
				result <- &ami{a, ctx, r}
			}
		}
	}()
	return result
}

type ami struct {
	RegionalAccount
	context.Context
	*ec2.Image
}

func (a ami) Ctx() context.Context {
	return a.Context
}

func (a ami) GetFields(builder cumulus.IFieldBuilder) {
	pub := "private"
	if aws.BoolValue(a.Image.Public) {
		pub = "public"
	}

	ec2_Tag_to_fields(builder, a.Context, a.Tags)

	builder.
		GID(aws.StringValue(a.ImageId)).
		Name(aws.StringValue(a.Image.Name)).
		Description(aws.StringValue(a.Image.Description)).
		What("type", pub).
		What("image_type", aws.StringValue(a.Image.ImageType)).
		Add(cumulus.FieldMeta{
			Kind:          cumulus.WHEN,
			Name:          "created",
			DefaultHidden: false,
		}, aws.StringValue(a.Image.CreationDate)).
		Done()
}
