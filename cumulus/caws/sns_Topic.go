package caws

import (
	aws "github.com/aws/aws-sdk-go/aws"
	cumulus "github.com/deweysasser/cumulus/cumulus"
)

// Code generated. DO NOT EDIT.  Code generated from sns_Topic.yaml

func (i topic) GeneratedFields(builder cumulus.IFieldBuilder) {
	if i.obj.TopicArn != nil {
		builder.What("topic_arn", aws.StringValue(i.obj.TopicArn))
	}

}
