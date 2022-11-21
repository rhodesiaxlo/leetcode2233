package graph

type iIterator interface {
	begin() int
	end() bool
	next() int
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
