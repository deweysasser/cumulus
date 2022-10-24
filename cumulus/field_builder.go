package cumulus

import "time"

type IFieldBuilder interface {
	Who(name, value string, options ...MetadataOption) *FieldBuilder
	What(name, value string, options ...MetadataOption) *FieldBuilder
	Where(name, value string, options ...MetadataOption) *FieldBuilder
	When(name string, t time.Time, options ...MetadataOption) *FieldBuilder
	How(name, value string, options ...MetadataOption) *FieldBuilder
	Why(name, value string, options ...MetadataOption) *FieldBuilder

	Tag(name, value string, options ...MetadataOption) *FieldBuilder

	GID(value string) *FieldBuilder
	Name(value string) *FieldBuilder
	Description(value string) *FieldBuilder
	Done()
}

type FieldBuilder struct {
	Fields Fields
}

type MetadataOption func(meta FieldMeta) FieldMeta

var DefaultHidden = func(meta FieldMeta) FieldMeta {
	meta.DefaultHidden = true
	return meta
}

func (f *FieldBuilder) Add(meta FieldMeta, value string, options ...MetadataOption) *FieldBuilder {
	for _, o := range options {
		meta = o(meta)
	}

	if v, ok := f.Fields[meta]; ok {
		v.Add(value)
	} else {
		f.Fields[meta] = NewFieldValue(value)
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

func (f *FieldBuilder) Who(name, value string, options ...MetadataOption) *FieldBuilder {
	meta := FieldMeta{
		Kind: WHO,
		Name: name,
	}
	return f.Add(meta, value, options...)
}

func (f *FieldBuilder) What(name, value string, options ...MetadataOption) *FieldBuilder {
	meta := FieldMeta{
		Kind: WHAT,
		Name: name,
	}
	return f.Add(meta, value, options...)
}

func (f *FieldBuilder) Where(name, value string, options ...MetadataOption) *FieldBuilder {
	meta := FieldMeta{
		Kind: WHERE,
		Name: name,
	}

	return f.Add(meta, value, options...)
}

func (f *FieldBuilder) When(name string, t time.Time, options ...MetadataOption) *FieldBuilder {
	meta := FieldMeta{
		Kind: WHEN,
		Name: name,
	}

	return f.Add(meta, t.String(), options...)
}

func (f *FieldBuilder) How(name, value string, options ...MetadataOption) *FieldBuilder {
	meta := FieldMeta{
		Kind: HOW,
		Name: name,
	}

	return f.Add(meta, value, options...)
}

func (f *FieldBuilder) Why(name, value string, options ...MetadataOption) *FieldBuilder {
	meta := FieldMeta{
		Kind: WHY,
		Name: name,
	}
	return f.Add(meta, value, options...)

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

func (f *FieldBuilder) Tag(name, value string, options ...MetadataOption) *FieldBuilder {
	meta := FieldMeta{
		Kind:          TAG,
		Name:          "tag:" + name,
		DefaultHidden: true,
	}
	return f.Add(meta, value, options...)
}
