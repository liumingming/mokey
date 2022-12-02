package ast

import "mokey/token"

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
	Statement  Statement
}

func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) expressNode() {
}

func (es *ExpressionStatement) statementNode() {

}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
