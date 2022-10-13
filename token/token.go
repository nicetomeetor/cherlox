package token

import "fmt"

type TokenType string

const (
	// Single-character tokens.
	LEFT_PAREN  TokenType = "("
	RIGHT_PAREN           = ")"
	LEFT_BRACE            = "{"
	RIGHT_BRACE           = "}"
	COMMA                 = ","
	DOT                   = "."
	MINUS                 = "-"
	PLUS                  = "+"
	SEMICOLON             = ";"
	SLASH                 = "/"
	STAR                  = "*"

	// One or two character tokens.
	BANG          = "!"
	BANG_EQUAL    = "!="
	EQUAL         = "="
	EQUAL_EQUAL   = "=="
	GREATER       = ">"
	GREATER_EQUAL = ">="
	LESS          = "<"
	LESS_EQUAL    = "<="

	// Literals.
	IDENTIFIER = "identifier"
	STRING     = "string"
	NUMBER     = "number"

	// Keywords.
	AND    = "and"
	CLASS  = "class"
	ELSE   = "else"
	FALSE  = "false"
	FUN    = "fun"
	FOR    = "for"
	IF     = "if"
	NIL    = "nil"
	OR     = "or"
	PRINT  = "print"
	RETURN = "return"
	SUPER  = "super"
	THIS   = "this"
	TRUE   = "true"
	VAR    = "var"
	WHILE  = "while"

	// End of file.
	EOF = "eof"
)

//Token contains the lexeme read by the scanner
type Token struct {
	Type    TokenType
	Lexeme  string
	Literal interface{}
	Line    int
}

func (t *Token) String() string {
	return fmt.Sprintf("%s %s %v", t.Type, t.Lexeme, t.Literal)
}
