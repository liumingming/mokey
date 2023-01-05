package repl

import (
	"mokey/eval"
	"mokey/lexer"
	"mokey/parser"
	"testing"
)

func TestRepl(t *testing.T)  {
	line := "1+1;"
	l := lexer.New(line)
	p := parser.New(l)

	program := p.ParseProgram()

	evaluated := eval.Eval(program)

	if evaluated == nil {
		t.Errorf("eval errors")
	}

	t.Logf("result is %s", evaluated.Inspect())
}
