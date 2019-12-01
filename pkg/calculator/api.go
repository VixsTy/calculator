// Package calculator have a bundle of code helping in the resolution of arithmetics expression
package calculator

// Calculator describe an interface which have to resolve mathematical expressions
type Calculator interface {
	// Calc resolve the operations pass in input string
	Calc(input string) string
}
