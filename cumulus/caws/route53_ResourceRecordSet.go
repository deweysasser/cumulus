package caws

import (
	"fmt"
	aws "github.com/aws/aws-sdk-go/aws"
	cumulus "github.com/deweysasser/cumulus/cumulus"
)

// Code generated. DO NOT EDIT.  Code generated from route53_ResourceRecordSet.yaml

func (i resourcerecordset) GeneratedFields(builder cumulus.IFieldBuilder) {

	if i.obj.Failover != nil {
		builder.What("failover", aws.StringValue(i.obj.Failover), cumulus.DefaultHidden)
	}

	if i.obj.HealthCheckId != nil {
		builder.What("health_check_id", aws.StringValue(i.obj.HealthCheckId), cumulus.DefaultHidden)
	}

	if i.obj.MultiValueAnswer != nil {
		builder.What("multi_value_answer", fmt.Sprint(boolToString(i.obj.MultiValueAnswer)), cumulus.DefaultHidden)
	}

	builder.Name(aws.StringValue(i.obj.Name))

	if i.obj.Region != nil {
		builder.What("region", aws.StringValue(i.obj.Region), cumulus.DefaultHidden)
	}

	if i.obj.SetIdentifier != nil {
		builder.What("set_identifier", aws.StringValue(i.obj.SetIdentifier), cumulus.DefaultHidden)
	}

	if i.obj.TTL != nil {
		builder.What("ttl", fmt.Sprint(aws.Int64Value(i.obj.TTL)), cumulus.DefaultHidden)
	}

	if i.obj.TrafficPolicyInstanceId != nil {
		builder.What("traffic_policy_instance_id", aws.StringValue(i.obj.TrafficPolicyInstanceId), cumulus.DefaultHidden)
	}

	if i.obj.Type != nil {
		builder.What("record_type", aws.StringValue(i.obj.Type))
	}

	if i.obj.Weight != nil {
		builder.What("weight", fmt.Sprint(aws.Int64Value(i.obj.Weight)), cumulus.DefaultHidden)
	}

}
