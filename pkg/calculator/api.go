package calculator

// Calculator describe an interface which have to resolve mathematical expressions
type Calculator interface {
	// Calc resolve the operations pass in input string
	Calc(input string) string
}
