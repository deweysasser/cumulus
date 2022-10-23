package cumulus

import "time"

type IFieldBuilder interface {
	Who(name, value string) *FieldBuilder
	What(name, value string) *FieldBuilder
	Where(name, value string) *FieldBuilder
	When(name string, t time.Time) *FieldBuilder
	How(name, value string) *FieldBuilder
	Why(name, value string) *FieldBuilder
	GID(value string) *FieldBuilder
	Name(value string) *FieldBuilder
	Description(value string) *FieldBuilder
	Tag(name, value string) *FieldBuilder
	Done()
}

type FieldBuilder struct {
	Fields Fields
}

func NewBuilder() *FieldBuilder {
	return &FieldBuilder{
		Fields: make([]Field, 0),
	}
}

// Done is syntactic sugar to help make a chain of calls more convenient
func (f *FieldBuilder) Done() {
}

func (f *FieldBuilder) Who(name, value string) *FieldBuilder {
	f.Fields = append(f.Fields, Field{
		FieldMeta: FieldMeta{
			Kind: WHO,
			Name: name,
		},
		Value: value,
	})

	return f
}

func (f *FieldBuilder) What(name, value string) *FieldBuilder {
	f.Fields = append(f.Fields, Field{
		FieldMeta: FieldMeta{
			Kind: WHAT,
			Name: name,
		},

		Value: value,
	})

	return f
}

func (f *FieldBuilder) Where(name, value string) *FieldBuilder {
	f.Fields = append(f.Fields, Field{
		FieldMeta: FieldMeta{
			Kind: WHERE,
			Name: name,
		},

		Value: value,
	})

	return f
}

func (f *FieldBuilder) When(name string, t time.Time) *FieldBuilder {
	f.Fields = append(f.Fields, Field{
		FieldMeta: FieldMeta{
			Kind: WHEN,
			Name: name,
		},

		Value: t.String(),
	})

	return f
}

func (f *FieldBuilder) How(name, value string) *FieldBuilder {
	f.Fields = append(f.Fields, Field{
		FieldMeta: FieldMeta{
			Kind: HOW,
			Name: name,
		},

		Value: value,
	})

	return f
}

func (f *FieldBuilder) Why(name, value string) *FieldBuilder {
	f.Fields = append(f.Fields, Field{
		FieldMeta: FieldMeta{
			Kind: WHY,
			Name: name,
		},

		Value: value,
	})

	return f
}

func (f *FieldBuilder) GID(value string) *FieldBuilder {
	f.Fields = append(f.Fields, Field{
		FieldMeta: IDMeta,
		Value:     value,
	})

	return f
}

func (f *FieldBuilder) Name(value string) *FieldBuilder {
	f.Fields = append(f.Fields, Field{
		FieldMeta: NameMeta,
		Value:     value,
	})

	return f
}

func (f *FieldBuilder) Description(value string) *FieldBuilder {
	f.Fields = append(f.Fields, Field{
		FieldMeta: DescriptionMeta,
		Value:     value,
	})

	return f
}

func (f *FieldBuilder) Tag(name, value string) *FieldBuilder {
	f.Fields = append(f.Fields, Field{
		FieldMeta: FieldMeta{
			Kind:          TAG,
			Name:          "tag:" + name,
			DefaultHidden: true,
		},

		Value: value,
	})

	return f
}
