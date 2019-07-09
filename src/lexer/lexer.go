// Package lexer lexer.go
package lexer

// Lexer type
type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

// New Lexer
func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}
