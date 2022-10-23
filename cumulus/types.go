package cumulus

//go:generate ./generate-type-signatures regional_types.go nonregional_types.go
//go:generate ./generate-account-wrappers RegionalAccounts generated-regional-accounts.go regional_types.go
//go:generate ./generate-account-wrappers Accounts generated-accounts.go nonregional_types.go

import (
	"context"
	"fmt"
)

// ErrorHandler is called by individual methods to handle and possibly abort processing
type ErrorHandler func(ctx context.Context, err error)
type Provider[T Fielder] func(ctx context.Context) chan T
type ProviderMethod[T Fielder] func(p Provider[T], ctx context.Context) chan T

type ID string

func (i ID) String() string {
	return string(i)
}

type Account interface {
	fmt.Stringer
	InRegion(region string) RegionalAccount
	Name() string
	Fielder
	AccountInfoer
}

type Accounts []Account

type RegionalAccount interface {
	Account
	Region() string
}

type RegionalAccounts []RegionalAccount

type Fielder interface {
	GetFields(builder IFieldBuilder)
}

type Sourcer interface {
	Source() Fielder
}

type Common interface {
	Fielder
	Sourcer
	Ctx() context.Context
}
