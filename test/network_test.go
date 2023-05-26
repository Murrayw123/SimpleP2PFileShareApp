package test

import (
	"p2pfileshare/internal/chord"
	"p2pfileshare/internal/network"
	"testing"
)

func TestNodeNetwork(t *testing.T) {
	t.Run("NewNetwork", func(t *testing.T) {
		n := network.NewNetwork()
		node := chord.NewNode()
		n.AddNode(node)

		if len(n.Nodes) != 1 {
			t.Errorf("Expected network node length of 1 but got %d", len(n.Nodes))
		}

		if n.Nodes[0].ID != node.ID {
			t.Errorf("The node ID in the network doesn't match with the newly created node")
		}
	})
}
