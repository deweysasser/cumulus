package caws

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/deweysasser/cumulus/cumulus"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"time"
)

func (a Account) AccountInfos(ctx context.Context) chan cumulus.AccountInfo {
	results := make(chan cumulus.AccountInfo)

	go func() {

		l := zerolog.Ctx(ctx)

		log := l.With().Str("account", string(a)).Logger()

		log.Debug().Msg("Getting account info")

		s, e := a.session()

		if e != nil {
			cumulus.HandleError(ctx, errors.Wrap(e, "Error getting session"))
			close(results)
			return
		}

		svc := sts.New(s)

		start := time.Now()
		out, e := svc.GetCallerIdentity(&sts.GetCallerIdentityInput{})
		CallTimer.Done(start)

		if e != nil {
			cumulus.HandleError(ctx, errors.Wrap(e, "Error calling STS"))
			close(results)
			return
		}

		log = log.With().Str("account_id", aws.StringValue(out.Account)).Logger()

		log.Debug().Msg("Sending results")
		results <- &accountInfo{
			account:                 a,
			GetCallerIdentityOutput: out,
			ctx:                     log.WithContext(ctx),
		}

		close(results)
	}()

	return results
}

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
		account:                 a,
		ctx:                     ctx,
		GetCallerIdentityOutput: out,
	})
}

type accountInfo struct {
	account Account
	*sts.GetCallerIdentityOutput
	ctx context.Context
}

func (a accountInfo) Account() cumulus.Account {
	return a.account
}

func (a accountInfo) Source() cumulus.Fielder {
	return a.account
}

func (a accountInfo) Ctx() context.Context {
	return a.ctx
}

func (a accountInfo) Name() string {
	return string(a.account)
}

func (a accountInfo) ID() string {
	return aws.StringValue(a.GetCallerIdentityOutput.Account)
}

func (a accountInfo) Text() string {
	return a.Name()
}

func (a accountInfo) GetFields(builder cumulus.IFieldBuilder) {
	builder.
		Name(a.Name()).
		GID(aws.StringValue(a.Arn)).
		Who("account_id", aws.StringValue(a.GetCallerIdentityOutput.Account)).
		How("user_id", aws.StringValue(a.GetCallerIdentityOutput.UserId)).
		Done()
}
