// Package node describe an entity to be manipulate in other data structure
package node

import "go/token"

// Node is the basic entity which will be manipulate in the LIFO Stack and FIFO Queue
type Node struct {
	Tok token.Token
	Lit string
	Val float64
}
