package caws

import (
	aws "github.com/aws/aws-sdk-go/aws"
	cumulus "github.com/deweysasser/cumulus/cumulus"
)

// Code generated. DO NOT EDIT.  Code generated from sns_Subscription.yaml

func (i subscription) GeneratedFields(builder cumulus.IFieldBuilder) {

	if i.obj.Endpoint != nil {
		builder.What("endpoint", aws.StringValue(i.obj.Endpoint), cumulus.DefaultHidden)
	}

	if i.obj.Owner != nil {
		builder.What("owner", aws.StringValue(i.obj.Owner), cumulus.DefaultHidden)
	}

	if i.obj.Protocol != nil {
		builder.What("protocol", aws.StringValue(i.obj.Protocol), cumulus.DefaultHidden)
	}

	if i.obj.SubscriptionArn != nil {
		builder.What("subscription_arn", aws.StringValue(i.obj.SubscriptionArn), cumulus.DefaultHidden)
	}

	if i.obj.TopicArn != nil {
		builder.What("topic_arn", aws.StringValue(i.obj.TopicArn), cumulus.DefaultHidden)
	}

}
