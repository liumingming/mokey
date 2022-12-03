package parser

import (
	"mokey/ast"
	"mokey/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	if len(program.Statements) != 3 {
		t.Fatalf("program does not contain three let statement. got %d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got = %q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not ast.Statement got %T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got =%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.TokenLiteral() not '%s'. got =%s", name, letStmt.Name.TokenLiteral())
		return false
	}
	return true
}

func TestReturnStatements(t *testing.T) {
	input := `
return 5;
return 10;
return 993  322;
`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	if len(program.Statements) != 3 {
		t.Fatalf("program does not contain three let statement. got %d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement")
			continue
		}

		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q", returnStmt.TokenLiteral())
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar"

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statement.got %d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statement[0] is not  ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("exp not *ast.Identifier. got = %T", stmt.Expression)
	}

	if ident.Value != "foobar" {
		t.Errorf("ident.Value not %s. got=%s", "foobar", ident.Value)
	}

	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral not %s. got=%s", "foobar", ident.TokenLiteral())
	}
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "5;"

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statement.got %d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statement[0] is not  ast.ExpressionStatement. got=%T", program.Statements[0])
	}
	ident, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("exp not *ast.IntegerLiteral. got = %T", stmt.Expression)
	}

	if ident.Value != 5 {
		t.Errorf("ident.Value not %s. got=%d", "foobar", ident.Value)
	}

	if ident.TokenLiteral() != "5" {
		t.Errorf("ident.TokenLiteral not %s. got=%s", "5", ident.TokenLiteral())
	}
}

func TestBooleanLiteralExpression(t *testing.T) {
	input := "true;"

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statement.got %d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statement[0] is not  ast.ExpressionStatement. got=%T", program.Statements[0])
	}
	boolean, ok := stmt.Expression.(*ast.Boolean)
	if !ok {
		t.Fatalf("exp not *ast.IntegerLiteral. got = %T", stmt.Expression)
	}

	if boolean.Value != true {
		t.Errorf("ident.Value not %s. got=%v", "true", boolean.Value)
	}

	if boolean.TokenLiteral() != "true" {
		t.Errorf("ident.TokenLiteral not %s. got=%s", "5", boolean.TokenLiteral())
	}
}
