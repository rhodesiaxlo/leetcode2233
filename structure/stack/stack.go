package stack

type Stack struct {
	a []rune
}

func (s *Stack) IsEmtpy() bool {
	return len(s.a) == 0
}

func (s *Stack) Size() int {
	return len(s.a)
}
func (s *Stack) Push(x rune) {
	s.a = append(s.a, x)
}

func (s *Stack) Pop() rune {
	sz := len(s.a)
	ret := s.a[sz-1]
	s.a = append([]rune{}, s.a[:sz-1]...)
	return ret
}

func (s *Stack) Top() rune {
	return s.a[len(s.a)-1]
}

func (s *Stack) Peek() rune {
	return s.a[0]
}

func (s *Stack) Clear()  {
	for !s.IsEmtpy() {
		s.Pop()
	}
}

