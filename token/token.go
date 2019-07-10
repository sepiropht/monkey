// token/token.go

package token

// TokenType string
type TokenType string

// Token type is
type Token struct {
	Type    TokenType
	Literal string
}

const (
	// ILLEGAL illegal
	ILLEGAL = "ILLEGAL"
	// EOF EOF
	EOF = "EOF"

	// Identifiers + literals

	//IDENT add, foobar, x, u ...
	IDENT = "IDENT"
	// INT 123456
	INT = "INT"

	// ASSIGN =
	ASSIGN = "="

	// PLUS "+"
	PLUS = "+"

	// COMMA Delimiters"
	COMMA = ","

	// SEMICOLON Delimiters
	SEMICOLON = ";"

	//LPAREN Function
	LPAREN = "("

	//RPAREN Function
	RPAREN = ")"

	//LBRACE Function
	LBRACE = "{"

	//RBRACE Function
	RBRACE = "}"

	// FUNCTION keywords
	FUNCTION = "FUNCTION"

	// LET keyword
	LET = "LET"
)
