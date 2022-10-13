package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"github.com/nicetomeetor/cherlox/token"
)

func error(line int, message string) {
	report(line, "", message)
}

func report(line int, where string, message string) {
	concatenated := fmt.Sprintf("[line %d] Error %s: %s", line, where, message)

	fmt.Println(concatenated)
}

// type Scanner struct {
// 	source  string
// 	tokens  []Token
// 	start   int //0
// 	current int // 0
// 	line    int //= 1;
// }

// func (s *Scanner) isAtEnd() bool {
// 	return s.current >= len(s.source)
// }

// func (s *Scanner) scanToken() {
// 	c := s.advance()

// 	switch c {
// 	case "(":
// 		s.addToken(LEFT_PAREN)
// 	case ")":
// 		s.addToken(RIGTH_PAREN)
// 	case "}":
// 		s.addToken(LEFT_BRACE)
// 	case "{":
// 		s.addToken(RIGHT_BRACE)
// 	case ",":
// 		s.addToken(COMMA)
// 	case ".":
// 		s.addToken(DOT)
// 	case "-":
// 		s.addToken(MINUS)
// 	case "+":
// 		s.addToken(PLUS)
// 	case "*":
// 		s.addToken(STAR)
// 	case ";":
// 		s.addToken(SEMICOLON)

// 	case '!':
// 		if s.match('=') {
// 			s.addToken(BANG_EQUAL)
// 		} else {
// 			s.addToken(BANG)
// 		}
// 	case '=':
// 		if s.match('=') {
// 			s.addToken(EQUAL_EQUAL)
// 		} else {
// 			s.addToken(EQUAL)
// 		}
// 	case '<':
// 		if s.match('=') {
// 			s.addToken(LESS_EQUAL)
// 		} else {
// 			s.addToken(LESS)
// 		}
// 	case '>':
// 		if s.match('=') {
// 			s.addToken(GREATER_EQUAL)
// 		} else {
// 			s.addToken(GREATER)
// 		}

// 	case '/':
// 		if match('/') {
// 			// A comment goes until the end of the line.

// 			for s.peek() != '\n' && !s.isAtEnd() {
// 				s.advance()
// 			}
// 		} else {
// 			s.addToken(SLASH)
// 		}

// 	case ' ':
// 	case '\r':
// 	case '\t':
// 		// Ignore whitespace.

// 	case '\n':
// 		s.line++

// 	default:
// 		error(s.line, "Unexpected character.")
// 	}
// }

// func (s *Scanner) peek() string {
// 	if s.isAtEnd() {
// 		return "\\0"
// 	}

// 	return s.source[s.current]
// }

// func (s *Scanner) match(expected string) bool {
// 	if s.isAtEnd() {
// 		return false
// 	}

// 	if s.source[s.current] != expected {
// 		return false
// 	}

// 	s.current++
// 	return true
// }

// func (s *Scanner) advance() string {
// 	value := s.current
// 	s.current = s.current + 1
// 	return s.source[value]
// }

// func (s *Scanner) addToken(tokenType TokenType) {
// 	//addToken(type, null);
// 	text := s.source[s.start:s.current]
// 	append(s.tokens, TokensToken{tokenType, text, s.literal, s.line})
// }

// func (s *Scanner) scanTokens() {
// 	for !s.isAtEnd() {
// 		s.start = s.current
// 		scanToken()
// 	}

// 	token := Token{EOF, "", nil, line}

// 	append(s.tokens, token)

// 	return s.tokens
// }

func main() {
	args := os.Args[1:]

	if len(args) > 1 {
		fmt.Println("Usage: cher [script]")
		os.Exit(0)
	} else if len(args) == 1 {
		fmt.Println("Run: cher [file]")
		runFile(args[0])
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

		if line == "exit\n" {
			os.Exit(0)
		}
	}
}

func runFile(path string) {

}
