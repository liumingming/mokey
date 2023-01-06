package object

import "fmt"

type Boolean struct {
	ObjectType ObjectType
	Value      bool
}

func (b Boolean) Type() ObjectType {
	return BooleanObject
}

func (b Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}
