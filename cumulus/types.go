package cumulus

//go:generate ./generate-type-signatures types.go

import (
	"context"
	"fmt"
)

// ErrorHandler is called by individual methods to handle and possibly abort processing
type ErrorHandler func(ctx context.Context, err error)
type Provider[T any] func(ctx context.Context) chan T

type ID string

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

type Common interface {
	Source() RegionalAccount
	Ctx() context.Context
	//Logger() zerolog.Logger
}

// AUTOGENERATE below this line

// AccountInfo gets details about a specific account
type AccountInfo interface {
	Source() Account
	Ctx() context.Context
	Name() string
	ID() string
}

type Instance interface {
	Common
	Id() ID
	JSON() string
	Text() string
}

type Zone interface {
	Id() ID
	JSON() string
	Text() string
}
