package graph

import (
	"fmt"
	"testing"
)

func TestBellmenFord_DetectNegativeCycle(t *testing.T) {
	wtsg, _ := NewWtSparseGraphFromFile("../../testG8", true)
	b := NewBellmenFord(&wtsg, wtsg.v, 0)
	fmt.Printf("bellman ford detect negative cycle is %t \n", b.hasNegativeCycle)
	if !b.hasNegativeCycle {
		b.ShortestPath(1)
		b.ShortestPath(2)
		b.ShortestPath(3)
		b.ShortestPath(4)
	}
}

func TestBellmenFord_DetectNegativeCycle2(t *testing.T) {
	wtsg, _ := NewWtSparseGraphFromFile("../../testG7", true)
	b := NewBellmenFord(&wtsg, wtsg.v, 0)
	fmt.Printf("bellman ford detect negative cycle is %t \n", b.hasNegativeCycle)
	if !b.hasNegativeCycle {
		b.ShortestPath(1)
		b.ShortestPath(2)
		b.ShortestPath(3)
		b.ShortestPath(4)
	}
}
