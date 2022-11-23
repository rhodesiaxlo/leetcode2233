package graph

import "testing"

func TestCreateWtDenseGraphFromAdd(t *testing.T) {
	// add
	// edge +1
	// remove edge -1
}

func TestCreateWtDenseGraphFromFile(t *testing.T) {
	wtdg, _ := NewWtDenseGraphFromFile("../../testG4")
	wtdg.show()
}

func TestIteratorWtDenseGraph(t *testing.T) {

}

func TestShowWtDenseGraph(t *testing.T) {

}


