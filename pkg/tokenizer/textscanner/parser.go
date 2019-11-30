package textscanner

import (
	"go/token"

	"github.com/asaskevich/govalidator"
)

//nolint:gocyclo // TODO: rework the cyclomatic complexity
func parse(lit string) token.Token {
	switch {
	case govalidator.IsInt(lit):
		return token.INT
	case govalidator.IsFloat(lit):
		return token.FLOAT
	case len(lit) == 1:
		r := []rune(lit)
		switch r[0] {
		case '+':
			return token.ADD
		case '-':
			return token.SUB
		case '/':
			return token.QUO
		case '*':
			return token.MUL
		case '(':
			return token.LPAREN
		case ')':
			return token.RPAREN
		case '%':
			return token.REM
		case '^':
			return token.XOR
		default:
			return token.ILLEGAL
		}
	default:
		return token.ILLEGAL
	}
}
