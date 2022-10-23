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

func (f *FieldBuilder) Add(meta FieldMeta, value string) *FieldBuilder {
	if v, ok := f.Fields[meta]; ok {
		v.Add(value)
	} else {
		f.Fields[meta] = fValue(value)
	}

	return f
}

func NewBuilder() *FieldBuilder {
	return &FieldBuilder{
		Fields: make(Fields),
	}
}

// Done is syntactic sugar to help make a chain of calls more convenient
func (f *FieldBuilder) Done() {
}

func (f *FieldBuilder) Who(name, value string) *FieldBuilder {
	meta := FieldMeta{
		Kind: WHO,
		Name: name,
	}
	return f.Add(meta, value)
}

func (f *FieldBuilder) What(name, value string) *FieldBuilder {
	meta := FieldMeta{
		Kind: WHAT,
		Name: name,
	}
	return f.Add(meta, value)
}

func (f *FieldBuilder) Where(name, value string) *FieldBuilder {
	meta := FieldMeta{
		Kind: WHERE,
		Name: name,
	}
	return f.Add(meta, value)
}

func (f *FieldBuilder) When(name string, t time.Time) *FieldBuilder {
	meta := FieldMeta{
		Kind: WHEN,
		Name: name,
	}

	return f.Add(meta, t.String())
}

func (f *FieldBuilder) How(name, value string) *FieldBuilder {
	meta := FieldMeta{
		Kind: HOW,
		Name: name,
	}

	return f.Add(meta, value)
}

func (f *FieldBuilder) Why(name, value string) *FieldBuilder {
	meta := FieldMeta{
		Kind: WHY,
		Name: name,
	}
	return f.Add(meta, value)

}

func (f *FieldBuilder) GID(value string) *FieldBuilder {
	meta := IDMeta
	return f.Add(meta, value)

}

func (f *FieldBuilder) Name(value string) *FieldBuilder {
	meta := NameMeta
	return f.Add(meta, value)

}

func (f *FieldBuilder) Description(value string) *FieldBuilder {
	meta := DescriptionMeta
	return f.Add(meta, value)

}

func (f *FieldBuilder) Tag(name, value string) *FieldBuilder {
	meta := FieldMeta{
		Kind:          TAG,
		Name:          "tag:" + name,
		DefaultHidden: true,
	}
	return f.Add(meta, value)
}
