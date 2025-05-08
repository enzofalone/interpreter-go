package lexer

import (
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

	stats, err := f.Stat()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting file description: %v\n", err)
		os.Exit(1)
	}

	var status = status.SUCCESS
	if stats.Size() > 0 {
		status = l.readFile(f)
	} else {
		fmt.Println("EOF null")
	}

	return status
}

// readFile traverses through file one byte at a time scanning and printing for different tokens found
func (l *Lexer) readFile(f *os.File) status.ReturnCode {
	line := 1
	lexicalError := false

	for {
		_, err := f.Seek(0, io.SeekCurrent)
		if err != nil {
			return status.UNKNOWN_ERROR
		}

		buffer := make([]byte, 1)
		n, _ := f.Read(buffer)
		if n == 0 {
			fmt.Println("EOF null")
			break
		}

		// raw char read
		char := string(buffer[:n])

		// identifier that translates char(s) to
		ident, err := token.LookupIdent(char)
		if err != nil {
			fmt.Printf("[line %d] Error: %s\n", line, err)
			lexicalError = false

			continue
		}

		if ident == token.NEWLINE {
			line++
		}

		fmt.Printf("%s %s null\n", ident, char)
	}

	if lexicalError {
		return status.LEXICAL_ERROR
	}

	return status.SUCCESS
}
