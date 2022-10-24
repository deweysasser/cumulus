package program

import (
	"github.com/deweysasser/cumulus/cumulus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewExpression(t *testing.T) {

	exp, err := ParseExpression("foo=bar")
	assert.NoError(t, err)

	fields := make(cumulus.Fields)

	fields[cumulus.FieldMeta{
		Kind:          cumulus.WHERE,
		Name:          "foo",
		DefaultHidden: false,
	}] = cumulus.NewFieldValue("baroom")

	assert.True(t, exp(fields))

	exp, err = ParseExpression("blah=whatever")
	assert.NoError(t, err)

	assert.False(t, exp(fields))
}

func TestParseFilter(t *testing.T) {

	exp, err := ParseFilter("foo=bar, baz=something")
	assert.NoError(t, err)

	fields := make(cumulus.Fields)

	fields[cumulus.FieldMeta{
		Kind:          cumulus.WHERE,
		Name:          "foo",
		DefaultHidden: false,
	}] = cumulus.NewFieldValue("baroom")

	fields[cumulus.FieldMeta{
		Kind:          cumulus.WHERE,
		Name:          "baz",
		DefaultHidden: false,
	}] = cumulus.NewFieldValue("something else")

	assert.True(t, exp(fields))

	exp, err = ParseFilter("foo=bar, baz=nothing")
	assert.NoError(t, err)

	assert.False(t, exp(fields))
}
