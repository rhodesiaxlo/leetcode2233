package graph

import "testing"

func TestCreateWtSparseGraphFromFile(t *testing.T) {
	wtdg, _ := NewWtSparseGraphFromFile("../../testG4", false)
	wtdg.show()
}
