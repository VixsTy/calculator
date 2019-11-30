package fifo

import "github.com/VixsTy/calculator/pkg/type/node"

// Queue is a simple and not optimized fifo data structure
type Queue struct {
	queue []*node.Node
	n     int
}

// Push add a value at the end of the queue
func (s *Queue) Push(in *node.Node) {
	s.queue = append(s.queue, in)
	s.n++
}

// Pop consumme the value at the begin of the queue
func (s *Queue) Pop() *node.Node {
	if s.n > 0 {
		v := s.queue[0]
		s.queue[0] = nil
		s.queue = s.queue[1:]
		s.n--
		return v
	}
	return nil
}

// Len return the size of the queue
func (s *Queue) Len() int {
	return s.n
}
