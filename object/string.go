package object

import "fmt"

type String struct {
	ObjectType ObjectType
	Value      string
}

func (i String) Type() ObjectType {
	return i.ObjectType
}

func (i String) Inspect() string {
	return fmt.Sprintf("%s", i.Value)
}
