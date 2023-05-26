package chord

import (
	"crypto/rand"
	"encoding/hex"
)

type Node struct {
	ID string
	//... Other properties will be added as needed
}

func NewNode() *Node {
	node := &Node{
		ID: generateNodeID(),
	}
	return node
}

// generateNodeID generates a unique ID for a node.
// Here we are using a random number but it could also be a hash of the node's address
func generateNodeID() string {
	b := make([]byte, 16) // generates a random 128-bit number
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(b)
}
