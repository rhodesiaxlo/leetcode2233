package graph

type SparseGraphIterator struct {
	sg *SparseGraph
	v int
	index int
}

func NewSparseGraphIterator(sg *SparseGraph, iterFrom int) *SparseGraphIterator {
	return &SparseGraphIterator{
		sg:    sg,
		v:     iterFrom,
		index: 0,
	}
}

func (s *SparseGraphIterator) begin() int {
	s.index = 0
	if s.index < len(s.sg.g[s.v]) {
		return s.sg.g[s.v][s.index]
	}
	return -1
}

func (s *SparseGraphIterator) next() int {
	s.index++
	if s.index < len(s.sg.g[s.v]) {
		return s.sg.g[s.v][s.index]
	}
	return -1
}

func (s *SparseGraphIterator) end() bool {
	return s.index >= len(s.sg.g[s.v])
}


