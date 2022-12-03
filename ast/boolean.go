package ast

import "mokey/token"

type Boolean struct {
	Token token.Token
	Value bool
}

func (i *Boolean) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Boolean) String() string {
	return i.Token.Literal
}

func (i *Boolean) expressNode() {
}
