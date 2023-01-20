package eval

import (
	"fmt"
	"monkey/ast"
	"monkey/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.IntegerLiteral:
		return &object.Integer{
			ObjectType: object.IntegerObject,
			Value:      node.Value,
		}
	case *ast.StringLiteral:
		return &object.String{
			ObjectType: object.StringObject,
			Value:      node.Value,
		}
	case *ast.Boolean:
		return &object.Boolean{
			ObjectType: object.BooleanObject,
			Value:      node.Value,
		}
	case *ast.InfixExpression:
		left := Eval(node.Left)
		right := Eval(node.Right)
		return evalInfixExpression(node.Operator, left, right)
	case *ast.PrefixExpression:
		right := Eval(node.Right)
		return evalPrefixExpression(node.Operator, right)
	default:
		fmt.Printf("node is %T type and node is have no eval function", node)
	}
	return nil
}

func evalStatements(statements []ast.Statement) object.Object {
	var result object.Object

	for _, st := range statements {
		result = Eval(st)
	}

	return result
}

func evalInfixExpression(operator string, left, right object.Object) object.Object {
	switch {
	case left.Type() == object.IntegerObject && right.Type() == object.IntegerObject:
		return evalIntegerInfixExpression(operator, left, right)
	case left.Type() == object.StringObject && right.Type() == object.StringObject:
		return evalStringInfixExpression(operator, left, right)
	default:
		return nil
	}
}

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch {
	case operator == "-" && right.Type() == object.IntegerObject:
		return evalIntegerPrefixExpression(operator, right)
	default:
		return nil
	}
}

func evalIntegerInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value

	switch operator {
	case "+":
		return &object.Integer{
			Value: leftVal + rightVal,
		}
	case "-":
		return &object.Integer{
			Value: leftVal - rightVal,
		}
	case "*":
		return &object.Integer{
			Value: leftVal * rightVal,
		}
	case "/":
		return &object.Integer{
			Value: leftVal / rightVal,
		}
	case "%":
		return &object.Integer{
			Value: leftVal % rightVal,
		}
	default:
		return nil
	}
}

func evalIntegerPrefixExpression(operator string, right object.Object) object.Object {
	rightVal := right.(*object.Integer).Value
	switch operator {
	case "-":
		return &object.Integer{
			ObjectType: object.IntegerObject,
			Value:      -rightVal,
		}
	default:
		return nil
	}
}

func evalStringInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.String).Value
	rightVal := right.(*object.String).Value

	switch operator {
	case "+":
		return &object.String{
			Value: leftVal + rightVal,
		}
	default:
		return nil
	}
}
