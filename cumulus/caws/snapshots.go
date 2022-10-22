package caws

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/deweysasser/golang-program/cumulus"
	"github.com/hashicorp/go-multierror"
	"github.com/rs/zerolog/log"
	"strings"
	"time"
)

func (a RegionalAccount) VisitSnapshot(ctx context.Context, cancel context.CancelFunc, visitor cumulus.SnapshotVisitor) error {

	s, e := a.session()
	if e != nil {
		return e
	}

	fmt.Sprint("Logger should print next")

	stssvc := sts.New(s)

	a.Read.Wait(ctx)
	start := time.Now()
	out, err := stssvc.GetCallerIdentity(&sts.GetCallerIdentityInput{})
	CallTimer.Done(start)

	if err != nil {
		return e
	}

	svc := ec2.New(s)

	a.Read.Wait(ctx)

	start = time.Now()
	snapshots, err := svc.DescribeSnapshotsWithContext(ctx, &ec2.DescribeSnapshotsInput{OwnerIds: []*string{out.Account}})
	CallTimer.Done(start)

	if err != nil {
		return e
	}

	for _, snap := range snapshots.Snapshots {
		select {
		case <-ctx.Done():
			return err
		default:
			l := log.Ctx(ctx).With().Str("instance_id", aws.StringValue(snap.SnapshotId)).Logger()
			err = multierror.Append(err, visitor(l.WithContext(ctx), cancel, snapshot{svc, snap, a}))
		}
	}
	return err
}

type snapshot struct {
	*ec2.EC2
	*ec2.Snapshot
	RegionalAccount
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

func (i snapshot) Text() string {

	name := ""
	for _, t := range i.Tags {
		if aws.StringValue(t.Key) == "Name" {
			name = aws.StringValue(t.Value)
			break
		}
	}

	return strings.Join([]string{
		aws.StringValue(i.SnapshotId),
		name,
		fmt.Sprint(aws.Int64Value(i.Snapshot.VolumeSize), "G"),
		fmt.Sprint(aws.TimeValue(i.Snapshot.StartTime)),
		aws.StringValue(i.Description),
	},
		"\t",
	)
}
