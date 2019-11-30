package node

import "go/token"

// Node is the basic entity which will be manipulate in the LIFO stack
type Node struct {
	Tok token.Token
	Lit string
	Val float64
}
