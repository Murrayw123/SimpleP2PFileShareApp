package test

import (
	"fmt"
	"p2pfileshare/internal/chord"
	"testing"
)

func Nodes(t *testing.T) {
	t.Run("NewNode", func(t *testing.T) {
		node := chord.NewNode()
		fmt.Println(node)
		if node.ID == "" {
			t.Errorf("New node does not have a unique ID")
		}
	})
}
