package program

import (
	"github.com/deweysasser/cumulus/cumulus"
	"github.com/stretchr/testify/assert"
	"strings"
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

func TestJoinedExpressions(t *testing.T) {

	tests := []struct {
		name   string
		args   string
		fields []string

		want bool
	}{
		{"single expression, single field", "a=alpha", []string{"a=alpha"}, true},
		{"single expression, single field", "a=alpha", []string{"a=alpha", "b=beta"}, true},
		{"single expression, single field", "a=alpha", []string{"a=something"}, false},

		{"single expression, single field", "a=alpha,b=beta", []string{"a=alpha"}, false},
		{"single expression, single field", "a=alpha,b=beta", []string{"a=alpha", "b=beta"}, true},
		{"single expression, single field", "a=alpha,b=beta", []string{"a=alpha", "b=gamma"}, false},

		{"single expression, single field", "tag:Environment=staging,region=us-east-2", []string{"tag:Environment=staging", "region=us-east-2"}, true},
		{"single expression, single field", "tag:Environment=staging,region=us-east-2", []string{"tag:Environment=staging", "region=us-east-1"}, false},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseFilter(tt.args)
			assert.NoError(t, err)

			fields := make(cumulus.Fields)
			for _, f := range tt.fields {
				ss := strings.Split(f, "=")
				fields[cumulus.FieldMeta{
					Kind:          cumulus.WHERE,
					Name:          ss[0],
					DefaultHidden: false,
				}] = cumulus.NewFieldValue(ss[1])
			}

			assert.Equalf(t, tt.want, got(fields), "ParseFilter(%v) on %v", tt.args, tt.fields)
		})
	}
}

func TestMultipleExpressions(t *testing.T) {

	tests := []struct {
		name   string
		args   []string
		fields []string

		want bool
	}{
		{"single expression, single field", []string{"a=alpha"}, []string{"a=alpha"}, true},
		{"single expression, single field", []string{"a=alpha"}, []string{"a=alpha", "b=beta"}, true},
		{"single expression, single field", []string{"a=alpha"}, []string{"a=something"}, false},

		{"single expression, single field", []string{"a=alpha,b=beta"}, []string{"a=alpha"}, false},
		{"single expression, single field", []string{"a=alpha,b=beta"}, []string{"a=alpha", "b=beta"}, true},
		{"single expression, single field", []string{"a=alpha,b=beta"}, []string{"a=alpha", "b=gamma"}, false},

		{"single expression, single field", []string{"a=alpha,b=beta", "b=gamma"}, []string{"a=alpha", "b=gamma"}, true},
		{"single expression, single field", []string{"a=alpha,b=beta", "a=delta, b=gamma"}, []string{"a=delta", "b=gamma"}, true},
		{"single expression, single field", []string{"a=alpha,b=beta", "a=delta, b=gamma"}, []string{"a=delta", "b=gamma"}, true},
		{"single expression, single field", []string{"a=alpha,b=beta", "a=delta, b=gamma"}, []string{"a=delta", "b=episilon"}, false},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseFilters(tt.args)
			assert.NoError(t, err)

			fields := make(cumulus.Fields)
			for _, f := range tt.fields {
				ss := strings.Split(f, "=")
				fields[cumulus.FieldMeta{
					Kind:          cumulus.WHERE,
					Name:          ss[0],
					DefaultHidden: false,
				}] = cumulus.NewFieldValue(ss[1])
			}

			assert.Equalf(t, tt.want, got(fields), "ParseFilter(%v) on %v", tt.args, tt.fields)
		})
	}
}
