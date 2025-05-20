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
	STAR                  = "STAR"
	DOT                   = "DOT"
	COMMA                 = "COMMA"
	PLUS                  = "PLUS"
	MINUS                 = "MINUS"
	SEMICOLON             = "SEMICOLON"
	SLASH                 = "SLASH"
	EQUAL                 = "EQUAL"
	BANG                  = "BANG"
	GREATER               = "GREATER"
	LESS                  = "LESS"
	STRING                = "STRING"

	// operators
	EQUAL_EQUAL   = "EQUAL_EQUAL"
	BANG_EQUAL    = "BANG_EQUAL"
	LESS_EQUAL    = "LESS_EQUAL"
	GREATER_EQUAL = "GREATER_EQUAL"

	// whitespace
	NEWLINE = "NEWLINE"
	SPACE   = "SPACE"
	TAB     = "TAB"
)

var tokens = map[string]TokenType{
	"(":  LEFT_PAREN,
	")":  RIGHT_PAREN,
	"{":  LEFT_BRACE,
	"}":  RIGHT_BRACE,
	"*":  STAR,
	",":  COMMA,
	".":  DOT,
	"+":  PLUS,
	"-":  MINUS,
	";":  SEMICOLON,
	"/":  SLASH,
	"=":  EQUAL,
	"!":  BANG,
	">":  GREATER,
	"<":  LESS,
	"\"": STRING,

	// operators
	"==": EQUAL_EQUAL,
	"!=": BANG_EQUAL,
	"<=": LESS_EQUAL,
	">=": GREATER_EQUAL,

	"\n": NEWLINE,
	" ":  SPACE,
	"\t": TAB,
}

func LookupIdent(ident string) (TokenType, error) {
	if tok, ok := tokens[ident]; ok {
		return tok, nil
	}
	return ERROR, fmt.Errorf("Unexpected character: %s", ident)
}

func (t *Token) String() string {
	return fmt.Sprintf("%s %s %s", t.TokenType, t.Lexeme, t.Literal)
}
