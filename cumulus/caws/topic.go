package caws

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/deweysasser/cumulus/cumulus"
	"github.com/rs/zerolog"
	"time"
)

func (a RegionalAccount) Topics(ctx context.Context) chan cumulus.Topic {
	results := make(chan cumulus.Topic)

	go func() {
		defer close(results)
		s, e := a.session()
		if e != nil {
			cumulus.HandleError(ctx, e)
			return
		}

		svc := sns.New(s)

		// TODO:  put in API call limiter
		start := time.Now()
		topics, e := svc.ListTopicsWithContext(ctx, nil)
		CallTimer.Done(start)

		if e != nil {
			cumulus.HandleError(ctx, e)
			return
		}

		for _, r := range topics.Topics {
			l := zerolog.Ctx(ctx).With().Str("topic_arn", aws.StringValue(r.TopicArn)).Logger()
			select {
			case <-ctx.Done():
				return
			case results <- topic{a, l.WithContext(ctx), r}:
			}
		}
	}()

	return results
}

type topic struct {
	account RegionalAccount
	context.Context
	obj *sns.Topic
}

func (i topic) Source() cumulus.Fielder {
	return i.account
}

func (i topic) Ctx() context.Context {
	return i.Context
}

func (i topic) Id() cumulus.ID {
	return cumulus.ID(aws.StringValue(i.obj.TopicArn))
}

func (i topic) GetFields(builder cumulus.IFieldBuilder) {
	builder.GID(aws.StringValue(i.obj.TopicArn))
	//i.GeneratedFields(builder)
}
