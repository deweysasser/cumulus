package caws

import "fmt"

//go:generate go run ../../build/generate.go caws ec2_Instance.yaml ec2_Snapshot.yaml ec2_Volume.yaml route53_HostedZone.yaml route53_ResourceRecordSet.yaml

func boolToString(b *bool) string {
	if b != nil && *b {
		return "true"
	}
	return "false"
}

func toSizeInG(i *int64) string {
	return fmt.Sprintf("%dG", *i)
}
