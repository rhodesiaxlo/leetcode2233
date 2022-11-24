package stack

type StackStr struct {
	a []string
	cap int
	len int
}

func NewStackStr(c int) StackStr {
	s := StackStr{
		a:   nil,
		cap: c,
	}

	s.a = make([]string, c)
	s.len = 0
	return s

}
func (s *StackStr) IsEmtpy() bool {
	return s.len == 0
}

func (s *StackStr) Size() int {
	return s.len
}
func (s *StackStr) Push(x string) {
	s.a[s.len+1] = x
	s.len++
}

func (s *StackStr) Pop() string {
	ret := s.a[s.len]
	s.len--
	return ret
}

func (s *StackStr) Top() string {
	return s.a[s.len]
}

func (s *StackStr) Peek() string {
	return s.a[s.len]
}

func (s *StackStr) Clear()  {
	for !s.IsEmtpy() {
		s.Pop()
	}
}
