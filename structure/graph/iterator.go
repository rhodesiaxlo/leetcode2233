package graph

type iIterator interface {
	begin() int
	end() bool
	next() int
}


type iWtIterator interface {
	begin() *edge
	end() bool
	next() *edge
}

func NewIterator(g interface{}, v int) iIterator {
	switch g.(type) {
	case *SparseGraph:
		sgPtr := g.(*SparseGraph)
		return NewSparseGraphIterator(sgPtr, v)
	case *DenseGraph:
		dgPtr := g.(*DenseGraph)
		return NewDenseGraphIterator(dgPtr, v)
	default:
		return nil
	}
}

func NewWtIterator(g interface{}, v int) iWtIterator {
	switch g.(type) {
	case *WtSparseGraph:
		wtsgPtr := g.(*WtSparseGraph)
		return NewWtSparseGraphIterator(wtsgPtr, v)
	case *WtDenseGraph:
		wtdgPtr := g.(*WtDenseGraph)
		return NewWtDenseGraphIterator(wtdgPtr, v)
	default:
		return nil
	}
}


