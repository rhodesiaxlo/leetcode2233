package graph

import (
	"fmt"
	"testing"
)

func TestPrimeMST_Init(t *testing.T) {
	wtsg, _ := NewWtSparseGraphFromFile("../../testG4")
	mst := NewPrimeMst(&wtsg, wtsg.v)
	mst.Show()
	fmt.Println(mst.mstWt)
}

func TestLazyPrimeMST_Init(t *testing.T) {
	wtsg, _ := NewWtSparseGraphFromFile("../../testG4")
	mst := newLazyPrimeMst(&wtsg, wtsg.v)
	mst.Show()
	fmt.Println(mst.mstWt)
}


func TestPer5ByLazyPrime(t *testing.T) {
	wtsg, _ := NewWtSparseGraphFromFile("../../testG5")
	mst := newLazyPrimeMst(&wtsg, wtsg.v)
	//mst.Show()
	fmt.Println(mst.mstWt)
}

func TestPer5ByPrime(t *testing.T) {
	wtsg, _ := NewWtSparseGraphFromFile("../../testG5")
	mst:=NewPrimeMst(&wtsg, wtsg.v)
	//mst.Show()
	fmt.Println(mst.mstWt)
}

func TestLazyPrimePrimeEqual5(t *testing.T) {
	wtsg1, _ := NewWtSparseGraphFromFile("../../testG5")
	mst1 := newLazyPrimeMst(&wtsg1, wtsg1.v)

	wtsg2, _ := NewWtSparseGraphFromFile("../../testG5")
	mst2:=NewPrimeMst(&wtsg2, wtsg2.v)

	if mst1.mstWt != mst2.mstWt {
		t.Fatal(fmt.Printf("prime lazy prime min weight not equal,  lazy:%3.2f, prime:%3.2f", mst1.mstWt, mst2.mstWt))
	}
}

func TestPerByLazyPrime6(t *testing.T) {
	wtsg, _ := NewWtSparseGraphFromFile("../../testG6")
	mst := newLazyPrimeMst(&wtsg, wtsg.v)
	//mst.Show()
	fmt.Println(mst.mstWt)
}

func TestPerByPrime6(t *testing.T) {
	wtsg, _ := NewWtSparseGraphFromFile("../../testG6")
	mst:=NewPrimeMst(&wtsg, wtsg.v)
	//mst.Show()
	fmt.Println(mst.mstWt)
}

func TestLazyPrimePrimeEqual6(t *testing.T) {
	wtsg, _ := NewWtSparseGraphFromFile("../../testG6")
	mst := newLazyPrimeMst(&wtsg, wtsg.v)

	wtsg2, _ := NewWtSparseGraphFromFile("../../testG6")
	mst2:=NewPrimeMst(&wtsg2, wtsg2.v)

	if mst.mstWt != mst2.mstWt {
		t.Fatal("prime lazy prime min weight not equal")
	}
}
