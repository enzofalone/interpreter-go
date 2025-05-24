package token

import "fmt"

type TokenType string

type Token struct {
	TokenType TokenType
	Lexeme    string
	Literal   interface{}
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
	NUMBER                = "NUMBER"

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

	// numbers
	"1": NUMBER,
	"2": NUMBER,
	"3": NUMBER,
	"4": NUMBER,
	"5": NUMBER,
	"6": NUMBER,
	"7": NUMBER,
	"8": NUMBER,
	"9": NUMBER,
	"0": NUMBER,

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
	switch v := t.Literal.(type) {
	case float64:
		if int(t.Literal.(float64)) == t.Literal {
			return fmt.Sprintf("%s %s %.1f", t.TokenType, t.Lexeme, t.Literal)
		}
		return fmt.Sprintf("%s %s %f", t.TokenType, t.Lexeme, t.Literal)
	case string:
		return fmt.Sprintf("%s %s %s", t.TokenType, t.Lexeme, t.Literal)
	default:
		return fmt.Sprintf("Unexpected type %T", v)
	}

}
