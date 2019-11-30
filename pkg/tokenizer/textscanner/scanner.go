package textscanner

import (
	"go/token"
	"strings"
	"text/scanner"

	"github.com/VixsTy/calculator/pkg/tokenizer"
)

// TextScanner implement the tokenizer.Tokenizer interface with the 'text/scanner' pkg : https://godoc.org/text/scanner
type TextScanner struct {
	tokenizer *scanner.Scanner
}

// NewTextScanner generate a TextScanner object and force to check implementation of the tokenizer.Tokenizer interface
func NewTextScanner() tokenizer.Tokenizer {
	return &TextScanner{}
}

// Init init a scanner.Scanner of the go/scanner package
func (s *TextScanner) Init(input string) {
	s.tokenizer = &scanner.Scanner{}
	s.tokenizer.Init(strings.NewReader(input))
}

// Scan forward in the token list and return the token literal, the token type and an error if it occur
// end of list is notify by token.EOF token type
func (s *TextScanner) Scan() (string, token.Token, error) {
	r := s.tokenizer.Scan()
	if r == scanner.EOF {
		return "", token.EOF, nil
	} else if lit := strings.TrimSpace(s.tokenizer.TokenText()); len(lit) > 0 {
		// create the token type from literal value
		tok := parse(lit)
		if tok == token.ILLEGAL {
			return "", tok, tokenizer.UnrecognizedToken
		}
		return lit, tok, nil
	} else {
		return "", token.ILLEGAL, tokenizer.UnrecognizedToken
	}
}
