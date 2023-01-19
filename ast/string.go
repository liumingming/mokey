package ast

import "monkey/token"

type StringLiteral struct {
	Token token.Token
	Value string
}

func (i *StringLiteral) TokenLiteral() string {
	return i.Token.Literal
}

func (i *StringLiteral) String() string {
	return i.Token.Literal
}

func (i *StringLiteral) expressNode() {
}
