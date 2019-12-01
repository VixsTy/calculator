// Package operator have a bundle of code describing operator of arithmetics expression
package operator

import (
	"go/token"
	"math"
)

// Associativity is an artificial type to declare an Associativity constants
type Associativity uint

const (
	// L is an Associativity for the LEFT
	L Associativity = iota
	// R is an Associativity for the RIGHT
	R
)

// Operator describe a mathematical operator
type Operator struct {
	Lit           string
	Tok           token.Token
	Precedence    uint
	Associativity Associativity
	ArgsLen       int
	Operation     func(args []float64) float64
}

// Operators is a map of all supported operator
var Operators = map[token.Token]*Operator{
	token.ADD: {
		Lit:           "+",
		Tok:           token.ADD,
		Precedence:    2,
		Associativity: L,
		ArgsLen:       2,
		Operation: func(args []float64) float64 {
			return args[1] + args[0]
		},
	},
	token.SUB: {
		Lit:           "-",
		Tok:           token.SUB,
		Precedence:    2,
		Associativity: L,
		ArgsLen:       2,
		Operation: func(args []float64) float64 {
			return args[1] - args[0]
		},
	},
	token.MUL: {
		Lit:           "*",
		Tok:           token.MUL,
		Precedence:    3,
		Associativity: L,
		ArgsLen:       2,
		Operation: func(args []float64) float64 {
			return args[1] * args[0]
		},
	},
	token.QUO: {
		Lit:           "/",
		Tok:           token.QUO,
		Precedence:    3,
		Associativity: L,
		ArgsLen:       2,
		Operation: func(args []float64) float64 {
			return args[1] / args[0]
		},
	},
	token.REM: {
		Lit:           "%",
		Tok:           token.REM,
		Precedence:    3,
		Associativity: L,
		ArgsLen:       2,
		Operation: func(args []float64) float64 {
			return math.Mod(args[1], args[0])
		},
	},
	token.XOR: {
		Lit:           "^",
		Tok:           token.XOR,
		Precedence:    4,
		Associativity: R,
		ArgsLen:       2,
		Operation: func(args []float64) float64 {
			return math.Pow(args[1], args[0])
		},
	},
}
