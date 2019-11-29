package textscanner

import (
	"io"
	"strings"
	"text/scanner"

	"github.com/VixsTy/calculator/pkg/tokenizer"
)

// TextScanner implement the tokenizer.Tokenizer interface with the 'text/scanner' pkg : https://godoc.org/text/scanner
type TextScanner struct {
	parser *scanner.Scanner
}

// NewTextScanner generate a TextScanner object and force to check implementation of the tokenizer.Tokenizer interface
func NewTextScanner() tokenizer.Tokenizer {
	return &TextScanner{}
}

// Init init a scanner.Scanner of the go/scanner package
func (s *TextScanner) Init(input string) {
	s.parser = &scanner.Scanner{}
	s.parser.Init(strings.NewReader(input))
}

// Scan forward in the token list and return the token or and io.EOF if it's the end of the string
func (s *TextScanner) Scan() (string, error) {
	tok := s.parser.Scan()
	if tok == scanner.EOF {
		return "", io.EOF
	}
	if value := strings.TrimSpace(s.parser.TokenText()); len(value) > 0 {
		return value, nil
	}
	return "", io.EOF
}
