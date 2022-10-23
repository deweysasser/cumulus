package cumulus

import "context"

// AccountInfo gets details about a specific account
type AccountInfo interface {
	Common
	Account() Account
	Ctx() context.Context
	Name() string
	ID() string
}

type Zone interface {
	Common
	Ctx() context.Context
	Id() ID
}

type NameRecord interface {
	Common
	Ctx() context.Context
	Id() ID
}
