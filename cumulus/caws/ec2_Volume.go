package caws

import (
	"fmt"
	aws "github.com/aws/aws-sdk-go/aws"
	cumulus "github.com/deweysasser/cumulus/cumulus"
)

// Code generated. DO NOT EDIT.  Code generated from ec2_Volume.yaml

func (i volume) GeneratedFields(builder cumulus.IFieldBuilder) {

	if i.obj.AvailabilityZone != nil {
		builder.What("availability_zone", aws.StringValue(i.obj.AvailabilityZone), cumulus.DefaultHidden)
	}

	if i.obj.CreateTime != nil {
		builder.When("create_time", aws.TimeValue(i.obj.CreateTime), cumulus.DefaultHidden)
	}

	if i.obj.Encrypted != nil {
		builder.What("encrypted", fmt.Sprint(boolToString(i.obj.Encrypted)), cumulus.DefaultHidden)
	}

	if i.obj.FastRestored != nil {
		builder.What("fast_restored", fmt.Sprint(boolToString(i.obj.FastRestored)), cumulus.DefaultHidden)
	}

	if i.obj.Iops != nil {
		builder.What("iops", fmt.Sprint(aws.Int64Value(i.obj.Iops)), cumulus.DefaultHidden)
	}

	if i.obj.KmsKeyId != nil {
		builder.What("kms_key_id", aws.StringValue(i.obj.KmsKeyId), cumulus.DefaultHidden)
	}

	if i.obj.MultiAttachEnabled != nil {
		builder.What("multi_attach_enabled", fmt.Sprint(boolToString(i.obj.MultiAttachEnabled)), cumulus.DefaultHidden)
	}

	if i.obj.OutpostArn != nil {
		builder.What("outpost_arn", aws.StringValue(i.obj.OutpostArn), cumulus.DefaultHidden)
	}

	if i.obj.Size != nil {
		builder.What("size", fmt.Sprint(toSizeInG(i.obj.Size)), cumulus.DefaultHidden)
	}

	if i.obj.SnapshotId != nil {
		builder.What("snapshot_id", aws.StringValue(i.obj.SnapshotId), cumulus.DefaultHidden)
	}

	if i.obj.State != nil {
		builder.What("state", aws.StringValue(i.obj.State), cumulus.DefaultHidden)
	}

	ec2_Tag_to_fields(builder, i.Ctx(), i.obj.Tags)

	if i.obj.Throughput != nil {
		builder.What("throughput", fmt.Sprint(aws.Int64Value(i.obj.Throughput)), cumulus.DefaultHidden)
	}

	if i.obj.VolumeId != nil {
		builder.What("volume_id", aws.StringValue(i.obj.VolumeId), cumulus.DefaultHidden)
	}

	if i.obj.VolumeType != nil {
		builder.What("volume_type", aws.StringValue(i.obj.VolumeType), cumulus.DefaultHidden)
	}

}
