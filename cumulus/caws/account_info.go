package caws

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/deweysasser/golang-program/cumulus"
	"time"
)

func (a Account) VisitAccountInfo(ctx context.Context, visitor cumulus.AccountInfoVisitor) error {

	s, e := a.session()
	if e != nil {
		return e
	}

	svc := sts.New(s)

	start := time.Now()
	out, err := svc.GetCallerIdentity(&sts.GetCallerIdentityInput{})
	CallTimer.Done(start)

	if err != nil {
		return e
	}

	return visitor(ctx, accountInfo{
		Account:                 a,
		GetCallerIdentityOutput: out,
	})
}

type accountInfo struct {
	Account
	*sts.GetCallerIdentityOutput
}

func (a accountInfo) Name() string {
	return string(a.Account)
}

func (a accountInfo) ID() string {
	return aws.StringValue(a.GetCallerIdentityOutput.Account)
}
