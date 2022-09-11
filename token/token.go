package token

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	//标识符 + 字面量
	IDENT = "IDENT"
	INT = "INT"

	//运算符
	ASSIGN = "="
	PLUS = "+"
	MINUS = "-"
	BANG = "!"
	ASTERISK = "*"
	SLASH = "/"
	LT = "<"
	GT = ">"

	EQ  = "=="
	NOT_EQ = "!="

	//分隔符
	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//关键字
	LET = "let"
	FUNCTION = "fn"
	TRUE = "true"
	FALSE = "false"
	IF  = "if"
	ELSE = "else"
	RETURN = "return"
)

var keywords = map[string]TokenType{
	"fn" : FUNCTION,
	"let" : LET,
	"true": TRUE,
	"false":FALSE,
	"if":IF,
	"else":ELSE,
	"return":RETURN,
}

type TokenType  string

type Token struct {
	Type TokenType
	Literal string
}


func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}