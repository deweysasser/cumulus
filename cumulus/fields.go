package cumulus

import (
	"fmt"
	set "github.com/deckarep/golang-set/v2"
	"sort"
	"strings"
	"time"
)

type FieldType int

const (
	LUID FieldType = iota
	WUID
	Name
	WHO
	WHAT
	WHEN
	WHERE
	WHY
	HOW
)

type FieldMeta struct {
	Kind FieldType
	Name string
}

type Field struct {
	FieldMeta
	Value string
}

type Fields []Field

func (fields Fields) String() string {
	sort.Slice(fields, func(i, j int) bool {
		return fields[i].Kind < fields[j].Kind
	})

	s := make([]string, len(fields))

	for i, f := range fields {
		s[i] = f.Value
		if f.Value == "" {
			s[i] = "-"
		}
		if strings.Contains(f.Value, " ") {
			s[i] = "\"" + f.Value + "\""
		}
	}

	return strings.Join(s, "\t")
}

type FieldsAccumulator struct {
	Lines   []map[string]string
	lengths map[string]int
	fields  set.Set[FieldMeta]
}

func NewAccumulator() FieldsAccumulator {
	return FieldsAccumulator{
		Lines:   make([]map[string]string, 0),
		lengths: make(map[string]int),
		fields:  set.NewSet[FieldMeta](),
	}
}

func (acc *FieldsAccumulator) Add(fields Fields) {
	m := make(map[string]string)
	for _, f := range fields {
		acc.fields.Add(f.FieldMeta)
		acc.lengths[f.Name] = max(acc.lengths[f.Name], len(f.Value))
		m[f.Name] = f.Value
	}

	acc.Lines = append(acc.Lines, m)
}

func (acc *FieldsAccumulator) Fields() []FieldMeta {
	fields := acc.fields.ToSlice()
	sort.Slice(fields, func(i, j int) bool {
		if fields[i].Kind == fields[j].Kind {
			return strings.Compare(fields[i].Name, fields[j].Name) < 1
		}
		return fields[i].Kind < fields[j].Kind
	})

	return fields
}

func (acc *FieldsAccumulator) Print() {
	padding := 3
	fields := acc.Fields()
	for _, field := range fields {
		name := strings.ToUpper(field.Name)
		name = strings.Replace(name, "_", " ", -1)

		fmt.Printf("%-*s", acc.lengths[field.Name]+padding, name)
	}
	fmt.Println()

	for _, line := range acc.Lines {
		for _, field := range fields {
			s := line[field.Name]
			switch {
			case s == "":
				s = "-"
			case strings.Contains(s, " "):
				s = "\"" + s + "\""
			}

			fmt.Printf("%-*s", acc.lengths[field.Name]+padding, s)
		}
		fmt.Println()
	}
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func NewBuilder() *FieldBuilder {
	return &FieldBuilder{
		Fields: make([]Field, 0),
	}
}

type FieldBuilder struct {
	Fields Fields
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

func (f *FieldBuilder) LUID(value string) *FieldBuilder {
	f.Fields = append(f.Fields, Field{
		FieldMeta: FieldMeta{
			Kind: LUID,
			Name: "ID",
		},

		Value: value,
	})

	return f
}

func (f *FieldBuilder) Name(value string) *FieldBuilder {
	f.Fields = append(f.Fields, Field{
		FieldMeta: FieldMeta{
			Kind: Name,
			Name: "name",
		},

		Value: value,
	})

	return f
}

func (f *FieldBuilder) WUID(value string) *FieldBuilder {
	f.Fields = append(f.Fields, Field{
		FieldMeta: FieldMeta{
			Kind: WUID,
			Name: "WUID",
		},

		Value: value,
	})

	return f
}
