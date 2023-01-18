package caws

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/deweysasser/cumulus/cumulus"
	"github.com/rs/zerolog/log"
	"time"
)

func (a RegionalAccount) DBClusters(ctx context.Context) chan cumulus.DBCluster {

	result := make(chan cumulus.DBCluster)

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

	out, err := svc.DescribeDBClustersWithContext(ctx, nil)
	CallTimer.Done(start)

	if err != nil {
		close(result)
		return result
	}

	go func() {
		defer close(result)

		for _, r := range out.DBClusters {
			select {
			case <-ctx.Done():
				return
			default:
				l := log.Ctx(ctx).With().Str("cluster_id", aws.StringValue(r.DBClusterArn)).Logger()
				ctx := l.WithContext(ctx)
				result <- &dbcluster{a, ctx, r}
			}
		}
	}()
	return result
}

type dbcluster struct {
	RegionalAccount
	context.Context
	obj *rds.DBCluster
}

func (a dbcluster) Ctx() context.Context {
	return a.Context
}

func (a dbcluster) GetFields(builder cumulus.IFieldBuilder) {

	a.GeneratedFields(builder)

}
