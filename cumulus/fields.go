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
	Values set.Set[string]
}

func (f FieldValue) String() string {
	//switch {
	//case f.Values == nil:
	//	return ""
	//case f.Values.Cardinality() < 1:
	//	return ""
	//case f.Values.Cardinality() == 1:
	//	c := f.Values.Iter()
	//	defer close(c)
	//	return <-c
	//default:
	//	return strings.Join(f.Values, ", ")
	//}
	// TODO: this probably puts more weight on the GC than we need, and copies stuff unnecessarily
	return strings.Join(f.Values.ToSlice(), ", ")
}

func fValue(s string) *FieldValue {
	if s != "" {
		return &FieldValue{Values: set.NewThreadUnsafeSet[string](s)}
	} else {
		return &FieldValue{Values: set.NewThreadUnsafeSet[string]()}
	}
}

func (f *FieldValue) Add(s string) {
	if s != "" {
		f.Values.Add(s)
	}
}

type Fields map[FieldMeta]*FieldValue

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
	Lines   []Fields
	lengths map[FieldMeta]int
	fields  set.Set[FieldMeta]
}

func NewAccumulator() FieldsAccumulator {
	return FieldsAccumulator{
		Lines:   make([]Fields, 0),
		lengths: make(map[FieldMeta]int),
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

	m := make(Fields)
	for k, v := range fields {
		acc.fields.Add(k)
		s := v.String()
		acc.lengths[k] = max(acc.lengths[k], len(s))
		m[k] = v
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

		acc.lengths[field] = max(acc.lengths[field], len(name))

		if printTitles {
			fmt.Printf("%-*s", acc.lengths[field]+padding, name)
		}
	}
	fmt.Println()

	for _, line := range acc.Lines {
		for _, field := range printFields {

			s := ""
			if v, ok := line[field]; ok {
				s = v.String()
			}
			
			switch {
			case s == "":
				s = "-"
			case strings.Contains(s, " "):
				s = "\"" + s + "\""
			}

			fmt.Printf("%-*s", acc.lengths[field]+padding, s)
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
