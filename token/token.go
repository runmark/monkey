package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF               = "EOF"

	// operator
	ASSIGN = "="
	PLUS   = "+"

	// identifier + literal
	IDNET = "IDENT"
	INT   = "INT"

	// keyword
	LET  = "LET"
	FUNC = "FUNC"

	// seperator
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
)
