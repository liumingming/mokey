package object

const (
	IntegerObject = "INTEGER"
	BooleanObject = "BOOLEAN"
)

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}
