package graph

import "testing"

func TestDijkStra_Init(t *testing.T) {
	wtsg, _ := NewWtSparseGraphFromFile("../../testG7", true)
	d := NewDijkStra(&wtsg, wtsg.v, 0)
	d.HasPathTo(3)

	d.ShortestPath(3)
	d.ShortestPath(1)
	d.ShortestPath(2)
	d.ShortestPath(4)
}
