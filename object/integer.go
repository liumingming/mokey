package object

import "fmt"

type Integer struct {
	ObjectType ObjectType
	Value      int64
}

func (i Integer) Type() ObjectType {
	return i.ObjectType
}

func (i Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}
