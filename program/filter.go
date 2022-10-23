package program

import (
	"github.com/deweysasser/cumulus/cumulus"
	"github.com/rs/zerolog/log"
	"regexp"
)

type Filter struct {
	Include, Exclude []*regexp.Regexp
}

var NoFilter = &Filter{
	Include: []*regexp.Regexp{regexp.MustCompile(".*")},
	Exclude: []*regexp.Regexp{},
}

func NewFilter(include, exclude []string) *Filter {
	inc := make([]*regexp.Regexp, 0, len(include))
	ex := make([]*regexp.Regexp, 0, len(exclude))

	for _, i := range include {
		r, e := regexp.Compile(i)
		if e != nil {
			log.Fatal().Err(e).Msg("Failed to parse field include regexp")
		}
		inc = append(inc, r)
	}

	for _, i := range exclude {
		r, e := regexp.Compile(i)
		if e != nil {
			log.Fatal().Err(e).Msg("Failed to parse field exclude regexp")
		}
		ex = append(ex, r)
	}

	return &Filter{inc, ex}
}

func (f *Filter) Accept(meta cumulus.FieldMeta) bool {
	log := log.With().Str("fieldname", meta.Name).Logger()

	if len(f.Include) > 0 {
		log.Debug().Msg("Has includes")
		for _, r := range f.Include {
			if r.MatchString(meta.Name) {
				return true
			}
		}
	}

	if len(f.Exclude) > 0 {
		log.Debug().Msg("Has excludes")
		for _, r := range f.Exclude {
			if r.MatchString(meta.Name) {
				return false
			}
		}
		return true
	}

	log.Debug().Bool("hidden", meta.DefaultHidden).Msg("returning default decision")

	return !meta.DefaultHidden
}
