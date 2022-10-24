package caws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/deweysasser/cumulus/cumulus"
	"github.com/rs/zerolog/log"
)

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

func (a RegionalAccount) Source() cumulus.Fielder {
	return a
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
