package program

import (
	"fmt"
	"github.com/deweysasser/cumulus/cumulus"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"regexp"
	"strings"
)

type LineFilter func(fields cumulus.Fields) bool

func AcceptAllLines(fields cumulus.Fields) bool {
	return true
}

func ParseFilters(s []string) (LineFilter, error) {
	filters := make([]LineFilter, 0, len(s))

	for _, s1 := range s {
		f, e := ParseFilter(s1)
		if e != nil {
			return nil, e
		}
		filters = append(filters, f)
	}

	// Implements "OR"
	return func(fields cumulus.Fields) bool {
		for _, f := range filters {
			if f(fields) {
				return true
			}
		}
		return false
	}, nil
}

func ParseFilter(s string) (LineFilter, error) {
	filters := make([]LineFilter, 0, len(s))

	parts := strings.Split(s, ",")

	for _, s1 := range parts {
		f, e := ParseExpression(strings.TrimSpace(s1))
		if e != nil {
			return nil, e
		}
		filters = append(filters, f)
	}

	// Implements "AND"
	return func(fields cumulus.Fields) bool {
		for _, f := range filters {
			if !f(fields) {
				return false
			}
		}
		return true
	}, nil
}

func ParseExpression(expr string) (LineFilter, error) {
	s := strings.Split(expr, "=")
	if len(s) != 2 {
		return nil, fmt.Errorf("Failed to split expression.  Expression has %d parts", len(s))
	}

	nrc, err := regexp.Compile(s[0])
	if err != nil {
		return nil, errors.Wrap(err, "failed to compile filter expression")
	}
	vrc, err := regexp.Compile(s[1])
	if err != nil {
		return nil, errors.Wrap(err, "failed to compile filter expression")
	}

	return func(fields cumulus.Fields) bool {
		for f, v := range fields {
			log.Debug().Str("field", f.Name).Str("value", v.String()).Msg("Evaluating expression against field")
			if nrc.MatchString(f.Name) {
				if vrc.MatchString(v.String()) {
					log.Debug().Str("field", f.Name).Str("value", v.String()).Msg("Matches")
					return true
				}
			}
		}
		return false
	}, nil
}
