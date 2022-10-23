package cumulus

type FieldType int

const (
	WHO FieldType = iota
	WHAT
	WHEN
	WHERE
	WHY
	HOW
)

type Field struct {
	Kind        FieldType
	Name, Value string
}

//func FieldBuilder
