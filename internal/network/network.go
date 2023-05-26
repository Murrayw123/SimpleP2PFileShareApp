package network

import (
	"p2pfileshare/internal/chord"
)

type Network struct {
	Nodes []*chord.Node
}

// NewNetwork initializes a new Network
func NewNetwork() *Network {
	return &Network{}
}

// AddNode adds a new node to the network
func (n *Network) AddNode(node *chord.Node) {
	n.Nodes = append(n.Nodes, node)
}
