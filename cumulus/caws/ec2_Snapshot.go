package caws

import (
	"fmt"
	aws "github.com/aws/aws-sdk-go/aws"
	cumulus "github.com/deweysasser/cumulus/cumulus"
)

// Code generated. DO NOT EDIT.  Code generated from ec2_Snapshot.yaml

func (i snapshot) GeneratedFields(builder cumulus.IFieldBuilder) {

	if i.obj.DataEncryptionKeyId != nil {
		builder.What("data_encryption_key_id", aws.StringValue(i.obj.DataEncryptionKeyId), cumulus.DefaultHidden)
	}

	if i.obj.Description != nil {
		builder.What("description", aws.StringValue(i.obj.Description), cumulus.DefaultHidden)
	}

	if i.obj.Encrypted != nil {
		builder.What("encrypted", fmt.Sprint(boolToString(i.obj.Encrypted)), cumulus.DefaultHidden)
	}

	if i.obj.KmsKeyId != nil {
		builder.What("kms_key_id", aws.StringValue(i.obj.KmsKeyId), cumulus.DefaultHidden)
	}

	if i.obj.OutpostArn != nil {
		builder.What("outpost_arn", aws.StringValue(i.obj.OutpostArn), cumulus.DefaultHidden)
	}

	if i.obj.OwnerAlias != nil {
		builder.What("owner_alias", aws.StringValue(i.obj.OwnerAlias), cumulus.DefaultHidden)
	}

	if i.obj.OwnerId != nil {
		builder.What("owner_id", aws.StringValue(i.obj.OwnerId), cumulus.DefaultHidden)
	}

	if i.obj.Progress != nil {
		builder.What("progress", aws.StringValue(i.obj.Progress), cumulus.DefaultHidden)
	}

	if i.obj.RestoreExpiryTime != nil {
		builder.When("restore_expiry_time", aws.TimeValue(i.obj.RestoreExpiryTime), cumulus.DefaultHidden)
	}

	if i.obj.SnapshotId != nil {
		builder.What("snapshot_id", aws.StringValue(i.obj.SnapshotId), cumulus.DefaultHidden)
	}

	if i.obj.StartTime != nil {
		builder.When("start_time", aws.TimeValue(i.obj.StartTime), cumulus.DefaultHidden)
	}

	if i.obj.State != nil {
		builder.What("state", aws.StringValue(i.obj.State), cumulus.DefaultHidden)
	}

	if i.obj.StateMessage != nil {
		builder.What("state_message", aws.StringValue(i.obj.StateMessage), cumulus.DefaultHidden)
	}

	if i.obj.StorageTier != nil {
		builder.What("storage_tier", aws.StringValue(i.obj.StorageTier), cumulus.DefaultHidden)
	}

	ec2_Tag_to_fields(builder, i.Ctx(), i.obj.Tags)

	if i.obj.VolumeId != nil {
		builder.What("volume_id", aws.StringValue(i.obj.VolumeId), cumulus.DefaultHidden)
	}

	if i.obj.VolumeSize != nil {
		builder.What("volume_size", fmt.Sprint(aws.Int64Value(i.obj.VolumeSize)), cumulus.DefaultHidden)
	}

}
