package graph

type WtDenseGraphIterator struct {
	dg    *WtDenseGraph
	v     int
	index int
}

func NewWtDenseGraphIterator(sg *WtDenseGraph, iterFrom int) *WtDenseGraphIterator {
	return &WtDenseGraphIterator{
		dg:    sg,
		v:     iterFrom,
		index: 0,
	}
}

func (s *WtDenseGraphIterator) begin() *edge {
	s.index = -1
	return s.next()
}

func (s *WtDenseGraphIterator) next() *edge {
	for s.index = s.index + 1; s.index < s.dg.V(); s.index++ {
		return s.dg.g[s.v][s.index]
	}

	return nil
}

func (s *WtDenseGraphIterator) end() bool {
	return s.index >= s.dg.v
}
