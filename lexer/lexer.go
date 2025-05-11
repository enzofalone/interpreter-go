package lexer

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/codecrafters-io/interpreter-starter-go/status"
	"github.com/codecrafters-io/interpreter-starter-go/token"
)

type Lexer struct {
	tokens []token.Token
}

func (l *Lexer) Scan(f *os.File) status.ReturnCode {
	status := l.readFile(f)
	return status
}

// readFile traverses through file one byte at a time scanning and printing for different tokens found
func (l *Lexer) readFile(f *os.File) status.ReturnCode {
	line := 1
	lexicalError := false

	for {
		char, err := next(f)
		if errors.Is(err, io.EOF) {
			break
		}

		// create output token
		ident, err := token.LookupIdent(char)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[line %d] Error: %s\n", line, err)
			lexicalError = true

			continue
		}
		if ident == token.NEWLINE {
			line++
		}

		// parse operators
		switch ident {
		case token.EQUAL:
			if match := match(f, "="); match {
				l.addToken(token.EQUAL_EQUAL, "==", "null", line)
			} else {
				l.addToken(token.EQUAL, char, "null", line)
			}
		case token.BANG:
			if match := match(f, "="); match {
				l.addToken(token.BANG_EQUAL, "!=", "null", line)
			} else {
				l.addToken(token.BANG, char, "null", line)
			}
		case token.GREATER:
			if match := match(f, "="); match {
				l.addToken(token.GREATER_EQUAL, ">=", "null", line)
			} else {
				l.addToken(token.GREATER, char, "null", line)
			}
		case token.LESS:
			if match := match(f, "="); match {
				l.addToken(token.LESS_EQUAL, "<=", "null", line)
			} else {
				l.addToken(token.LESS, char, "null", line)
			}
		case token.SLASH:
			if match := match(f, "/"); match {
				for {
					// read until newline
					c, err := next(f)
					if c == "\n" || err != nil {
						break
					}
				}
				continue
			} else {
				l.addToken(token.SLASH, char, "null", line)
			}
		default:
			l.addToken(ident, char, "null", line)
		}
	}
	l.addToken(token.EOF, "", "null", line)

	// print
	for i := range l.tokens {
		fmt.Println(l.tokens[i].String())
	}

	if lexicalError {
		return status.LEXICAL_ERROR
	}

	return status.SUCCESS
}

// func (l *Lexer) scanToken(char string) {
// 	char
// }

// Advance returns the next character in f's cursor. advance returns an error if EOF is found
func next(f *os.File) (string, error) {
	buffer := make([]byte, 1)
	n, err := f.Read(buffer)

	// raw char read
	char := string(buffer[:n])
	return char, err
}

// prev returns the previous characters in f's cursor.
func prev(f *os.File) (string, error) {
	// os.File.Read moves cursor by bytes read, so we account for it by doing n-1
	_, err := f.Seek(-2, io.SeekCurrent)
	if err != nil {
		return "", err
	}

	buffer := make([]byte, 1)
	n, err := f.Read(buffer)

	char := string(buffer[:n])
	return char, err
}

// match moves forward in f's cursor only if the expected character is met
//
// match is mostly used to identify operators (==, !=, >=)
func match(f *os.File, expected string) bool {
	buf := make([]byte, 1)
	n, err := f.Read(buf)
	if err != nil {
		return false
	}
	if n == 0 {
		return false
	}

	if string(buf[0]) == expected {
		return true
	}

	prev(f)
	return false
}

func matchString(f *os.File) (string, error) {
	result := ""

	for {
		buf := make([]byte, 1)

		n, err := f.Read(buf)
		if err != nil {
			return "", err
		}
		if n == 0 {
			break
		}

		char := string(buf)
		result += char

		if char == "\"" {
			break
		}
	}
	return result, nil
}

func (l *Lexer) addToken(tokenType token.TokenType, lexeme string, literal string, line int) {
	l.tokens = append(l.tokens, token.Token{
		TokenType: tokenType,
		Lexeme:    lexeme,
		Literal:   literal,
		Line:      line,
	})
}

func printError(line int, err error) {
	fmt.Fprintf(os.Stderr, "[line %d] Error: %s\n", line, err)
}
