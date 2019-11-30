package goscanner

import (
	"go/scanner"
	"go/token"

	"github.com/VixsTy/calculator/pkg/tokenizer"
)

// GoScanner implement the tokenizer.Tokenizer interface with the 'go/scanner' pkg : https://godoc.org/go/scanner
type GoScanner struct {
	parser *scanner.Scanner
}

// NewGoScanner generate a GoScanner object and force to check implementation the tokenizer.Tokenizer interface
func NewGoScanner() tokenizer.Tokenizer {
	return &GoScanner{}
}

// Init init a scanner.Scanner of the go/scanner package
func (s *GoScanner) Init(input string) {
	s.parser = &scanner.Scanner{}
	src := []byte(input)
	fset := token.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(src))
	s.parser.Init(file, src, nil, 0)
}

// Scan forward in the token list and return the token literal, the token type and an error if it occur
// end of list is notify by token.EOF token type
func (s *GoScanner) Scan() (string, token.Token, error) {
	_, tok, lit := s.parser.Scan()
	switch tok {
	case token.INT, token.FLOAT:
		return lit, tok, nil
	case token.EOF, token.SEMICOLON:
		return "", token.EOF, nil
	case token.LPAREN, token.RPAREN, token.ADD, token.SUB, token.MUL, token.QUO, token.REM, token.XOR:
		return tok.String(), tok, nil
	default:
		return "", token.ILLEGAL, tokenizer.UnrecognizedToken
	}
}
