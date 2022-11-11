package caws

import (
	"fmt"
	aws "github.com/aws/aws-sdk-go/aws"
	cumulus "github.com/deweysasser/cumulus/cumulus"
)

// Code generated. DO NOT EDIT.  Code generated from ec2_Instance.yaml

func (i instance) GeneratedFields(builder cumulus.IFieldBuilder) {

	if i.obj.AmiLaunchIndex != nil {
		builder.What("ami_launch_index", fmt.Sprint(aws.Int64Value(i.obj.AmiLaunchIndex)), cumulus.DefaultHidden)
	}

	if i.obj.Architecture != nil {
		builder.What("architecture", aws.StringValue(i.obj.Architecture), cumulus.DefaultHidden)
	}

	if i.obj.BootMode != nil {
		builder.What("boot_mode", aws.StringValue(i.obj.BootMode), cumulus.DefaultHidden)
	}

	if i.obj.CapacityReservationId != nil {
		builder.What("capacity_reservation_id", aws.StringValue(i.obj.CapacityReservationId), cumulus.DefaultHidden)
	}

	if i.obj.ClientToken != nil {
		builder.What("client_token", aws.StringValue(i.obj.ClientToken), cumulus.DefaultHidden)
	}

	if i.obj.EbsOptimized != nil {
		builder.What("ebs_optimized", fmt.Sprint(boolToString(i.obj.EbsOptimized)), cumulus.DefaultHidden)
	}

	if i.obj.EnaSupport != nil {
		builder.What("ena_support", fmt.Sprint(boolToString(i.obj.EnaSupport)), cumulus.DefaultHidden)
	}

	if i.obj.Hypervisor != nil {
		builder.What("hypervisor", aws.StringValue(i.obj.Hypervisor), cumulus.DefaultHidden)
	}

	if i.obj.ImageId != nil {
		builder.What("image_id", aws.StringValue(i.obj.ImageId))
	}

	builder.GID(aws.StringValue(i.obj.InstanceId))

	if i.obj.InstanceLifecycle != nil {
		builder.What("instance_lifecycle", aws.StringValue(i.obj.InstanceLifecycle), cumulus.DefaultHidden)
	}

	if i.obj.InstanceType != nil {
		builder.What("instance_type", aws.StringValue(i.obj.InstanceType))
	}

	if i.obj.Ipv6Address != nil {
		builder.What("ipv_6_address", aws.StringValue(i.obj.Ipv6Address), cumulus.DefaultHidden)
	}

	if i.obj.KernelId != nil {
		builder.What("kernel_id", aws.StringValue(i.obj.KernelId), cumulus.DefaultHidden)
	}

	if i.obj.KeyName != nil {
		builder.What("key_name", aws.StringValue(i.obj.KeyName), cumulus.DefaultHidden)
	}

	if i.obj.LaunchTime != nil {
		builder.When("launch_time", aws.TimeValue(i.obj.LaunchTime), cumulus.DefaultHidden)
	}

	if i.obj.OutpostArn != nil {
		builder.What("outpost_arn", aws.StringValue(i.obj.OutpostArn), cumulus.DefaultHidden)
	}

	if i.obj.Platform != nil {
		builder.What("platform", aws.StringValue(i.obj.Platform), cumulus.DefaultHidden)
	}

	if i.obj.PlatformDetails != nil {
		builder.What("platform_details", aws.StringValue(i.obj.PlatformDetails), cumulus.DefaultHidden)
	}

	if i.obj.PrivateDnsName != nil {
		builder.Where("private_dns_name", aws.StringValue(i.obj.PrivateDnsName))
	}

	if i.obj.PrivateIpAddress != nil {
		builder.Where("private_ip_address", aws.StringValue(i.obj.PrivateIpAddress))
	}

	if i.obj.PublicDnsName != nil {
		builder.Where("public_dns_name", aws.StringValue(i.obj.PublicDnsName))
	}

	if i.obj.PublicIpAddress != nil {
		builder.Where("public_ip_address", aws.StringValue(i.obj.PublicIpAddress))
	}

	if i.obj.RamdiskId != nil {
		builder.What("ramdisk_id", aws.StringValue(i.obj.RamdiskId), cumulus.DefaultHidden)
	}

	if i.obj.RootDeviceName != nil {
		builder.What("root_device_name", aws.StringValue(i.obj.RootDeviceName), cumulus.DefaultHidden)
	}

	if i.obj.RootDeviceType != nil {
		builder.What("root_device_type", aws.StringValue(i.obj.RootDeviceType), cumulus.DefaultHidden)
	}

	if i.obj.SourceDestCheck != nil {
		builder.What("source_dest_check", fmt.Sprint(boolToString(i.obj.SourceDestCheck)), cumulus.DefaultHidden)
	}

	if i.obj.SpotInstanceRequestId != nil {
		builder.What("spot_instance_request_id", aws.StringValue(i.obj.SpotInstanceRequestId), cumulus.DefaultHidden)
	}

	if i.obj.SriovNetSupport != nil {
		builder.What("sriov_net_support", aws.StringValue(i.obj.SriovNetSupport), cumulus.DefaultHidden)
	}

	if i.obj.StateTransitionReason != nil {
		builder.What("state_transition_reason", aws.StringValue(i.obj.StateTransitionReason), cumulus.DefaultHidden)
	}

	if i.obj.SubnetId != nil {
		builder.What("subnet_id", aws.StringValue(i.obj.SubnetId), cumulus.DefaultHidden)
	}

	ec2_Tag_to_fields(builder, i.Ctx(), i.obj.Tags)

	if i.obj.TpmSupport != nil {
		builder.What("tpm_support", aws.StringValue(i.obj.TpmSupport), cumulus.DefaultHidden)
	}

	if i.obj.UsageOperation != nil {
		builder.What("usage_operation", aws.StringValue(i.obj.UsageOperation), cumulus.DefaultHidden)
	}

	if i.obj.UsageOperationUpdateTime != nil {
		builder.When("usage_operation_update_time", aws.TimeValue(i.obj.UsageOperationUpdateTime), cumulus.DefaultHidden)
	}

	if i.obj.VirtualizationType != nil {
		builder.What("virtualization_type", aws.StringValue(i.obj.VirtualizationType), cumulus.DefaultHidden)
	}

	if i.obj.VpcId != nil {
		builder.What("vpc_id", aws.StringValue(i.obj.VpcId), cumulus.DefaultHidden)
	}

}
