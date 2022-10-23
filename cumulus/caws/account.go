package caws

import (
	"fmt"
	// TODO:  move to AWS SDK for go v2
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/deweysasser/cumulus/cumulus"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/sasbury/mini"
	"golang.org/x/time/rate"
)

// TODO:  turn this into a struct so we can have an associated rate limiter.  Alternatively, make it call always through
// a specific region so it's a regional account

// Account represents an AWS account
type Account string

func (a Account) Name() string {
	return string(a)
}

func (a Account) String() string {
	return string(a)
}

func (a Account) GetFields(builder cumulus.IFieldBuilder) {
	builder.Where("account_name", string(a))
}

func (a Account) session() (*session.Session, error) {
	//logger := log.With().Str("profile", string(a)).Str("region", profile.region).Logger()

	if s, err := session.NewSessionWithOptions(
		session.Options{
			Profile: string(a),
		},
	); err == nil {
		//if s, err := session.NewSessionWithOptions(session.Options{Profile: string(a), Config: aws.Config{Region: aws.String(profile.region)}}); err == nil {
		return s, nil
	} else {
		log.Error().Err(err).Msg("Failed to create session")
		return nil, err
	}
}

type Limit struct {
	Read   *rate.Limiter
	Modify *rate.Limiter
}

var DefaultRateLimit = Limit{
	Read:   rate.NewLimiter(200, 200),
	Modify: rate.NewLimiter(8, 100),
}

type RegionalAccount struct {
	Account
	region string
	*Limit
}

func (a RegionalAccount) String() string {
	return fmt.Sprint(a.Account, "/", a.region)
}

func (a RegionalAccount) GetFields(builder cumulus.IFieldBuilder) {
	a.Account.GetFields(builder)
	builder.Where("region", a.region)
}

func (a RegionalAccount) Region() string {
	return a.region
}

func (a RegionalAccount) session() (*session.Session, error) {
	//logger := log.With().Str("profile", string(a)).Str("region", profile.region).Logger()

	if s, err := session.NewSessionWithOptions(session.Options{Profile: a.Account.Name(), Config: aws.Config{Region: aws.String(a.region)}}); err == nil {
		return s, nil
	} else {
		log.Error().Err(err).Msg("Failed to create session")
		return nil, err
	}
}

func (a Account) InRegion(region string) cumulus.RegionalAccount {
	return &RegionalAccount{
		Account: a,
		region:  region,
		Limit:   &DefaultRateLimit,
	}
}

func AvailableAccountsFrom(path string) (cumulus.Accounts, error) {

	config, err := mini.LoadConfiguration(path)

	if err != nil {
		return nil, errors.Wrap(err, "Error reading "+path)
	}

	var accounts cumulus.Accounts

	for _, s := range config.SectionNames() {
		accounts = append(accounts, Account(s))
	}

	return accounts, nil
}
