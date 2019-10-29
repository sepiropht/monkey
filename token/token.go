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

	// MINUS "-"
	MINUS = "-"

	// BANG operators
	BANG = "!"

	// ASTERISK operators
	ASTERISK = "*"

	// SLASH operators
	SLASH = "/"

	// LT operators
	LT = "<"

	// GT operators
	GT = ">"

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

	// TRUE bool
	TRUE = "TRUE"

	// FALSE bool
	FALSE = "FALSE"

	// IF if
	IF = "IF"

	// ELSE else
	ELSE = "ELSE"

	// RETURN return
	RETURN = "RETURN"

	// EQ double eq
	EQ = "=="

	// NOTEQ negation
	NOT_EQ = "!="

	STRING = "STRING"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdent check
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
