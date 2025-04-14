package token

import "fmt"

type TokenType string

type Token struct {
	TokenType TokenType
	Lexeme    string
	Literal   string
	Line      int
}

const (
	LEFT_PAREN  TokenType = "LEFT_PAREN"
	RIGHT_PAREN           = "RIGHT_PAREN"
	ERROR                 = "ERROR"
	EOF                   = "EOF"
	LEFT_BRACE            = "LEFT_BRACE"
	RIGHT_BRACE           = "RIGHT_BRACE"
)

var tokens = map[string]TokenType{
	"(": LEFT_PAREN,
	")": RIGHT_PAREN,
	"{": LEFT_BRACE,
	"}": RIGHT_BRACE,
}

func LookupIdent(ident string) (TokenType, error) {
	if tok, ok := tokens[ident]; ok {
		return tok, nil
	}
	return ERROR, fmt.Errorf("LookupIdent: Could not identify ident. Unexpected\n")
}

func (t *Token) String() string {
	return fmt.Sprintf("%s %s %s", t.TokenType, t.Lexeme, t.Literal)
}
