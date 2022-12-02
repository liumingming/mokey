package ast

import "mokey/token"

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressNode() {

}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Value
}
