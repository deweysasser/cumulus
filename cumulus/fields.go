package cumulus

import (
	"fmt"
	set "github.com/deckarep/golang-set/v2"
	"sort"
	"strings"
)

type FieldType int

const (
	GID FieldType = iota
	NAME
	WHO
	WHAT
	WHEN
	WHERE
	WHY
	HOW
	DESCRIPTION
	TAG
)

var (
	DescriptionMeta = FieldMeta{
		Kind:          DESCRIPTION,
		Name:          "description",
		DefaultHidden: true,
	}

	NameMeta = FieldMeta{
		Kind: NAME,
		Name: "name",
	}

	IDMeta = FieldMeta{
		Kind: GID,
		Name: "ID",
	}
)

type FieldMeta struct {
	Kind          FieldType
	Name          string
	DefaultHidden bool
}

type FieldValue struct {
	Values []string
}

func (f FieldValue) String() string {
	switch {
	case f.Values == nil:
		return ""
	case len(f.Values) < 1:
		return ""
	case len(f.Values) == 1:
		return f.Values[0]
	default:
		return strings.Join(f.Values, ", ")
	}
}

func fValue(s string) FieldValue {
	return FieldValue{Values: []string{s}}
}

func (f *FieldValue) Add(s string) {
	f.Values = append(f.Values, s)
}

type Fields map[FieldMeta]FieldValue

//
//func (fields Fields) String() string {
//	sort.Slice(fields, func(i, j int) bool {
//		return fields[i].Kind < fields[j].Kind
//	})
//
//	s := make([]string, len(fields))
//
//	for i, f := range fields {
//		s[i] = f.Value
//		if f.Value == "" {
//			s[i] = "-"
//		}
//		if strings.Contains(f.Value, " ") {
//			s[i] = "\"" + f.Value + "\""
//		}
//	}
//
//	return strings.Join(s, "\t")
//}

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

type Filter interface {
	Accept(meta FieldMeta) bool
}

func (acc *FieldsAccumulator) Add(fielder Fielder) {
	b := NewBuilder()
	fielder.GetFields(b)

	// If we can find a source, also get the fields there
	if s, ok := fielder.(Sourcer); ok {
		s.Source().GetFields(b)
	}

	fields := b.Fields

	m := make(map[string]string)
	for k, v := range fields {
		acc.fields.Add(k)
		s := v.String()
		acc.lengths[k.Name] = max(acc.lengths[k.Name], len(s))
		m[k.Name] = s
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

func (acc *FieldsAccumulator) Print(f Filter, printTitles bool) {
	padding := 3
	fields := acc.Fields()
	// TODO:  put the title in only when we're verbose

	printFields := make([]FieldMeta, 0)

	for _, field := range fields {
		if !f.Accept(field) {
			continue
		}
		printFields = append(printFields, field)

		//name := strings.Replace(field.Name, "_", " ", -1)

		name := field.Name

		acc.lengths[field.Name] = max(acc.lengths[field.Name], len(name))

		if printTitles {
			fmt.Printf("%-*s", acc.lengths[field.Name]+padding, name)
		}
	}
	fmt.Println()

	for _, line := range acc.Lines {
		for _, field := range printFields {

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
