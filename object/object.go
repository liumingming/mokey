package object

const (
	IntegerObject = "INTEGER"
	BooleanObject = "BOOLEAN"
	StringObject  = "STRING"
)

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}
