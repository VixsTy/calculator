package tokenizer

import (
	"go/token"

	"golang.org/x/xerrors"
)

// Tokenizer is an interface to define a way how parse a litteral string in a list of token.
type Tokenizer interface {
	// Init should initialize the tokenizer
	Init(string)
	// Scan forward in the token list and return the token literal, the token type and an error if it occur
	// end of list is notify by token.EOF token type
	Scan() (string, token.Token, error)
}

// UnrecognizedToken hold a general error
var UnrecognizedToken error = xerrors.Errorf("Unrecognized token")
