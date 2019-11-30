package lifo

import (
	"github.com/VixsTy/calculator/pkg/type/node"
)

// Stack is a simple and not optimized lifo data structure
type Stack struct {
	stack []*node.Node
	n     int
}

// Push add a value at the end of the stack
func (s *Stack) Push(in *node.Node) {
	s.stack = append(s.stack, in)
	s.n++
}

// Pop consumme the value at the end of the stack
func (s *Stack) Pop() *node.Node {
	if s.n > 0 {
		s.n--
		v := s.stack[s.n]
		s.stack[s.n] = nil
		s.stack = s.stack[:s.n]
		return v
	}
	return nil
}

// Len return the size of the stack
func (s *Stack) Len() int {
	return s.n
}
