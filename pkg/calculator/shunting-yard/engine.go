package shuntingyard

import (
	"go/token"
	"strconv"

	"github.com/VixsTy/calculator/pkg/calculator"
	"github.com/VixsTy/calculator/pkg/operator"
	"github.com/VixsTy/calculator/pkg/tokenizer"
	"github.com/VixsTy/calculator/pkg/tokenizer/textscanner"
	"github.com/VixsTy/calculator/pkg/type/fifo"
	"github.com/VixsTy/calculator/pkg/type/lifo"
	"github.com/VixsTy/calculator/pkg/type/node"
	"golang.org/x/xerrors"
)

// ShuntingYard implement calculator.Calculator with the Shunting-yard algorithm : https://en.wikipedia.org/wiki/Shunting-yard_algorithm
type ShuntingYard struct {
	tokenizer tokenizer.Tokenizer
	operator  lifo.Stack
	output    fifo.Queue
}

// NewShuntingYard return a new calculator.Calculator using the Shunting yard algorithm
func NewShuntingYard() calculator.Calculator {
	return &ShuntingYard{}
}

// Calc resolve the operations pass in input string
func (s *ShuntingYard) Calc(input string) string {
	s.tokenizer = textscanner.NewTextScanner()
	s.tokenizer.Init(input)
	s.buildRPN()
	result := s.resolveRPN()
	return strconv.FormatFloat(result, 'G', -1, 64)
}

func (s *ShuntingYard) buildRPN() {
	// while there are tokens to be read do:
	//     read a token.
	lit, tok, err := s.tokenizer.Scan()
	for ; err == nil && tok != token.EOF; lit, tok, err = s.tokenizer.Scan() {
		//     if the token is a number, then:
		switch tok {
		case token.INT, token.FLOAT:
			//         push it to the output queue.
			s.output.Push(&node.Node{
				Tok: tok,
				Lit: lit,
			})
			//     if the token is a function then:
			//         push it onto the operator stack
			//     if the token is an operator, then:
		case token.ADD, token.SUB, token.MUL, token.QUO, token.REM, token.XOR:
			//         while ((there is a function at the top of the operator stack)
			//                or (there is an operator at the top of the operator stack with greater precedence)
			//                or (the operator at the top of the operator stack has equal precedence and is left associative))
			//               and (the operator at the top of the operator stack is not a left parenthesis):
			stackTok := s.operator.Pop()
			for ; stackTok != nil && isPriorOperator(tok, stackTok.Tok); stackTok = s.operator.Pop() {
				//             pop operators from the operator stack onto the output queue.
				s.output.Push(stackTok)
			}
			if stackTok != nil { // replace the operator at the top of the stack
				s.operator.Push(stackTok)
			}
			//         push it onto the operator stack.
			s.operator.Push(&node.Node{
				Tok: tok,
				Lit: lit,
			})
			//     if the token is a left paren (i.e. "("), then:
		case token.LPAREN:
			//         push it onto the operator stack.
			s.operator.Push(&node.Node{
				Tok: tok,
				Lit: lit,
			})
			//     if the token is a right paren (i.e. ")"), then:
		case token.RPAREN:
			//         while the operator at the top of the operator stack is not a left paren:
			stackTok := s.operator.Pop()
			for ; stackTok != nil && stackTok.Tok != token.LPAREN; stackTok = s.operator.Pop() {
				//             pop the operator from the operator stack onto the output queue.
				s.output.Push(stackTok)
			}
			//         /* if the stack runs out without finding a left paren, then there are mismatched parentheses. */
			if stackTok == nil {
				panic(xerrors.Errorf("parenthesis mismatch"))
			} else {
				//         if there is a left paren at the top of the operator stack, then:
				//             pop the operator from the operator stack and discard it
			}
		}
	}
	if err != nil {
		panic(err)
	}

	// after while loop, if operator stack not null, pop everything to output queue
	stackTok := s.operator.Pop()
	for ; stackTok != nil; stackTok = s.operator.Pop() {
		s.output.Push(stackTok)
	}
}

func isPriorOperator(tok, stackTok token.Token) bool {
	//         while ((there is a function at the top of the operator stack)
	//                or (there is an operator at the top of the operator stack with greater precedence)
	//                or (the operator at the top of the operator stack has equal precedence and is left associative))
	//               and (the operator at the top of the operator stack is not a left parenthesis):
	if stackTok == token.LPAREN {
		return false
	}

	evaluatedOp := operator.Operators[tok]
	stackOp := operator.Operators[stackTok]
	return (stackOp.Precedence > evaluatedOp.Precedence || (stackOp.Precedence == evaluatedOp.Precedence && stackOp.Associativity == operator.L))
}

func (s *ShuntingYard) resolveRPN() float64 {

	var stack lifo.Stack

	t := s.output.Pop()
	for ; t != nil; t = s.output.Pop() {
		switch t.Tok {
		case token.INT, token.FLOAT:
			value, err := strconv.ParseFloat(t.Lit, 64)
			if err != nil {
				panic(err)
			}
			t.Val = value
			stack.Push(t)
		case token.ADD, token.SUB, token.MUL, token.QUO, token.REM, token.XOR:
			op := operator.Operators[t.Tok]
			var args []float64
			for i := 0; i < op.ArgsLen; i++ {
				item := stack.Pop()
				if item == nil {
					panic(xerrors.Errorf("mismatch arguments"))
				}
				args = append(args, item.Val)
			}
			result := op.Operation(args)
			stack.Push(&node.Node{
				Tok: token.FLOAT,
				Val: result,
			})
		}
	}

	if stack.Len() > 1 {
		panic(xerrors.Errorf("error during the calcul"))
	}

	return stack.Pop().Val
}
