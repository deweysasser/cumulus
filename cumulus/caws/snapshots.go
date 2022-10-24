package caws

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/deweysasser/cumulus/cumulus"
	"github.com/rs/zerolog/log"
	"time"
)

func (a RegionalAccount) Snapshots(ctx context.Context) chan cumulus.Snapshot {
	results := make(chan cumulus.Snapshot)

	s, e := a.session()
	if e != nil {
		cumulus.HandleError(ctx, e)
		close(results)
		return results
	}

	stssvc := sts.New(s)

	a.Read.Wait(ctx)
	start := time.Now()
	out, e := stssvc.GetCallerIdentity(&sts.GetCallerIdentityInput{})
	CallTimer.Done(start)

	if e != nil {
		cumulus.HandleError(ctx, e)
		close(results)
		return results
	}

	svc := ec2.New(s)

	go func() {
		defer close(results)

		a.Read.Wait(ctx)

		start = time.Now()
		snapshots, err := svc.DescribeSnapshotsWithContext(ctx, &ec2.DescribeSnapshotsInput{OwnerIds: []*string{out.Account}})
		CallTimer.Done(start)

		if err != nil {
			cumulus.HandleError(ctx, e)
			return
		}

		for _, snap := range snapshots.Snapshots {
			select {
			case <-ctx.Done():
				cumulus.HandleError(ctx, ctx.Err())
				return
			case results <- snapshot{ctx, svc, snap, a}:
			}
		}
	}()

	return results
}

type snapshot struct {
	context.Context
	*ec2.EC2
	*ec2.Snapshot
	RegionalAccount
}

func (i snapshot) Source() cumulus.Fielder {
	return i.RegionalAccount
}

func (i snapshot) Ctx() context.Context {
	return i.Context
}

func (i snapshot) Delete(ctx context.Context, dryRun bool) error {

	defer CallTimer.Call().Done()
	i.RegionalAccount.Modify.Wait(ctx)
	_, err := i.EC2.DeleteSnapshotWithContext(ctx, &ec2.DeleteSnapshotInput{SnapshotId: i.SnapshotId, DryRun: aws.Bool(dryRun)})
	if ae, ok := err.(awserr.Error); ok {
		if dryRun && ae.Code() == "DryRunOperation" {
			return nil
		}
	}
	return err
}

func (i snapshot) Id() cumulus.ID {
	return cumulus.ID(aws.StringValue(i.SnapshotId))
}

func (i snapshot) JSON() string {
	bytes, err := json.Marshal(i)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to marshal")
	}

	return string(bytes)
}

func (i snapshot) GetFields(builder cumulus.IFieldBuilder) {

	builder.
		GID(aws.StringValue(i.SnapshotId)).
		What("size", fmt.Sprint(aws.Int64Value(i.Snapshot.VolumeSize), "G")).
		When("start_time", aws.TimeValue(i.Snapshot.StartTime)).
		Description(aws.StringValue(i.Description))

	tagFields(builder, i.Tags)
}
