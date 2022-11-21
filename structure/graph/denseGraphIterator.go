package graph

type DenseGraphIterator struct {
	dg    *DenseGraph
	v     int
	index int
}

func NewDenseGraphIterator(dg *DenseGraph, i int) *DenseGraphIterator {
	return &DenseGraphIterator{
		dg:    dg,
		v:     i,
		index: 0,
	}
}

func (s *DenseGraphIterator) begin() int {
	s.index = -1
	return s.next()
}

func (s *DenseGraphIterator) next() int {
	for s.index = s.index + 1; s.index < s.dg.V(); s.index++ {
		if s.dg.g[s.v][s.index] == true {
			return s.index
		}
	}

	return s.index
}

func (s *DenseGraphIterator) end() bool {
	return s.index >= s.dg.v
}
