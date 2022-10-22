package cumulus

import "fmt"

type Account interface {
	fmt.Stringer
	InRegion(region string) RegionalAccount
	Name() string
}

type RegionalAccount interface {
	fmt.Stringer
	Account
	Region() string
}
