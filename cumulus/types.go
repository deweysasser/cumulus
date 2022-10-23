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
type Provider[T Texter] func(ctx context.Context) chan T
type ProviderMethod[T Texter] func(p Provider[T], ctx context.Context) chan T

type ID string

func (i ID) String() string {
	return string(i)
}

type Account interface {
	fmt.Stringer
	InRegion(region string) RegionalAccount
	Name() string
	AccountInfoer
}

type Accounts []Account

type RegionalAccount interface {
	fmt.Stringer
	Account
	Region() string
}

type RegionalAccounts []RegionalAccount

type Texter interface {
	Text() string
}

type Fielder interface {
	Fields() Fields
}

type Common interface {
	Texter
	Fielder
	Source() string
	Ctx() context.Context
	Text() string
	//JSON() string
}
