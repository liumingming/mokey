package eval

import (
	"mokey/ast"
	"mokey/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.IntegerLiteral:
		return &object.Integer{
			ObjectType: object.IntegerObject,
			Value:      node.Value,
		}

	}
	return nil
}
