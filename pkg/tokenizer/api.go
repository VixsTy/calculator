package tokenizer

// Tokenizer is an interface to define a way how parse a litteral string in a list of token.
type Tokenizer interface {
	// Init should initialize the tokenizer
	Init(string)
	// Scan is expected to scan a new token and return it in first and else return the error io.EOF if there is no more token
	Scan() (string, error)
}
