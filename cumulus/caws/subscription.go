package caws

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/deweysasser/cumulus/cumulus"
	"github.com/rs/zerolog"
	"time"
)

func (a RegionalAccount) Subscriptions(ctx context.Context) chan cumulus.Subscription {
	results := make(chan cumulus.Subscription)

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
		Subscriptions, e := svc.ListSubscriptionsWithContext(ctx, nil)
		CallTimer.Done(start)

		if e != nil {
			cumulus.HandleError(ctx, e)
			return
		}

		for _, r := range Subscriptions.Subscriptions {
			l := zerolog.Ctx(ctx).With().Str("Subscription_arn", aws.StringValue(r.SubscriptionArn)).Logger()
			select {
			case <-ctx.Done():
				return
			case results <- subscription{a, l.WithContext(ctx), r}:
			}
		}
	}()

	return results
}

type subscription struct {
	account RegionalAccount
	context.Context
	obj *sns.Subscription
}

func (i subscription) Source() cumulus.Fielder {
	return i.account
}

func (i subscription) Ctx() context.Context {
	return i.Context
}

func (i subscription) Id() cumulus.ID {
	return cumulus.ID(aws.StringValue(i.obj.SubscriptionArn))
}

func (i subscription) GetFields(builder cumulus.IFieldBuilder) {
	i.GeneratedFields(builder)
}
