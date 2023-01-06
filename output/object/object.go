package object

const (
	IntegerObject = "INTEGER"
)

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}
