package object

type String struct {
	ObjectType ObjectType
	Value      string
}

func (i String) Type() ObjectType {
	return StringObject
}

func (i String) Inspect() string {
	return i.Value
}
