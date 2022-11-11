package caws

import (
	"fmt"
	aws "github.com/aws/aws-sdk-go/aws"
	cumulus "github.com/deweysasser/cumulus/cumulus"
)

// Code generated. DO NOT EDIT.  Code generated from route53_HostedZone.yaml

func (i hostedzone) GeneratedFields(builder cumulus.IFieldBuilder) {

	if i.obj.CallerReference != nil {
		builder.What("caller_reference", aws.StringValue(i.obj.CallerReference), cumulus.DefaultHidden)
	}

	builder.GID(aws.StringValue(i.obj.Id))

	builder.Name(aws.StringValue(i.obj.Name))

	if i.obj.ResourceRecordSetCount != nil {
		builder.What("resource_record_set_count", fmt.Sprint(aws.Int64Value(i.obj.ResourceRecordSetCount)), cumulus.DefaultHidden)
	}

}
