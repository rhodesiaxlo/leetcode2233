package graph

type WtSparseGraphIterator struct {
	sg    *WtSparseGraph
	v     int
	index int
}

func NewWtSparseGraphIterator(sg *WtSparseGraph, iterFrom int) *WtSparseGraphIterator {
	return &WtSparseGraphIterator{
		sg:    sg,
		v:     iterFrom,
		index: 0,
	}
}

func (s *WtSparseGraphIterator) begin() *edge {
	s.index = 0
	if s.index < len(s.sg.g[s.v]) {
		return s.sg.g[s.v][s.index]
	}
	return nil
}

func (s *WtSparseGraphIterator) next() *edge {
	s.index++
	if s.index < len(s.sg.g[s.v]) {
		return s.sg.g[s.v][s.index]
	}
	return nil
}

func (s *WtSparseGraphIterator) end() bool {
	return s.index >= len(s.sg.g[s.v])
}
