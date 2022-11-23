package graph

import (
	"fmt"
	"testing"
)

func TestLazyPrimeMst(t *testing.T) {
	wtsg, _ := NewWtSparseGraphFromFile("../../testG4")
	mst := newLazyPrimeMst(&wtsg, wtsg.v)
	//fmt.Println(mst.MsgEdges())
	mst.Show()
	fmt.Println(mst.mstWt)
}
