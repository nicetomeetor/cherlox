package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
)

type TokenType int

const (
    // Single-character tokens.
  	LEFT_PAREN TokenType  = iota 
  	RIGHT_PAREN 
	LEFT_BRACE
	RIGHT_BRACE
  	COMMA 
	DOT 
	MINUS 
	PLUS 
	SEMICOLON 
	SLASH 
	STAR

  	// One or two character tokens.
  	BANG 
	BANG_EQUAL
  	EQUAL 
	EQUAL_EQUAL
  	GREATER 
	GREATER_EQUAL
  	LESS 
	LESS_EQUAL

  	// Literals.
  	IDENTIFIER 
	STRING
	NUMBER

  	// Keywords.
  	AND
	CLASS 
	ELSE 
	FALSE 
	FUN 
	FOR 
	IF
	NIL 
	OR
  	PRINT 
	RETURN 
	SUPER 
	THIS
	TRUE 
	VAR 
	WHILE

	// End of file.
  	EOF
)

type Token struct {
	tokenType TokenType
	lexeme string
	literal int
	line int
}

func (t *Token) toString() {
	concatenated := fmt.Sprintf("%d %s %d", t.tokenType, t.lexeme, t.literal)

    fmt.Println(concatenated)
}

func main() {
    args := os.Args[1:]

	if (len(args) > 1) {
		fmt.Println("Usage: cher [script]")
		os.Exit(0)
 	} else if (len(args) == 1) {
		fmt.Println("Run: cher [file]")
  	} else {
		fmt.Println("Run: cher [prompt]")
		runPrompt()
  	}
}

func runPrompt() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		if (line == "exit\n") {
			os.Exit(0)
		}
	}
}
