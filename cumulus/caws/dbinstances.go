package caws

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/deweysasser/cumulus/cumulus"
	"github.com/rs/zerolog/log"
	"time"
)

func (a RegionalAccount) DBInstances(ctx context.Context) chan cumulus.DBInstance {

	result := make(chan cumulus.DBInstance)

	s, e := a.session()
	if e != nil {
		//return e
		close(result)
		return result
	}

	start := time.Now()

	svc := rds.New(s)

	if e = a.Read.Wait(ctx); e != nil {
		cumulus.HandleError(ctx, e)
		close(result)
		return result
	}

	out, err := svc.DescribeDBInstancesWithContext(ctx, nil)
	CallTimer.Done(start)

	if err != nil {
		close(result)
		return result
	}

	go func() {
		defer close(result)

		for _, r := range out.DBInstances {
			select {
			case <-ctx.Done():
				return
			default:
				l := log.Ctx(ctx).With().Str("Instance_id", aws.StringValue(r.DBInstanceArn)).Logger()
				ctx := l.WithContext(ctx)
				result <- &dbinstance{a, ctx, r}
			}
		}
	}()
	return result
}

type dbinstance struct {
	RegionalAccount
	context.Context
	obj *rds.DBInstance
}

func (a dbinstance) Ctx() context.Context {
	return a.Context
}

func (a dbinstance) GetFields(builder cumulus.IFieldBuilder) {

	a.GeneratedFields(builder)

}
