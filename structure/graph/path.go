package graph

import (
	"fmt"
	stack2 "leetcode/leetcode/structure/stack"
)

type Path struct {
	src     int
	visited []bool
	from    []int
	g       interface{}
	v       int
}

func NewPath(graph interface{}, src int, vNumber int) Path {
	p := Path{
		src:     src,
		visited: nil,
		from:    nil,
		g:       graph,
		v:       vNumber,
	}
	p.Init()
	return p
}

func (p *Path) Init() {
	p.visited = make([]bool, p.v)
	p.from = make([]int, p.v)
	for i := 0; i < p.v; i++ {
		p.from[i] = -1
	}
}

func (p *Path) CalcPath() {
	p.dfs(p.src)
}

func (p *Path) HasPath(des int) bool {
	return p.visited[des]
}

func (p *Path) GetPath(des int) []rune {
	if !p.HasPath(des) {
		return nil
	}

	stack := stack2.Stack{}
	cur := des
	for cur != -1 {
		stack.Push(rune(cur))
		cur = p.from[cur]
	}

	ans := []rune(nil)
	for !stack.IsEmtpy() {
		ans = append(ans, stack.Pop())
	}
	return ans
}

func (p *Path) ShowPath(in int) {
	fmt.Println(p.GetPath(in))
}

func (p *Path) dfs(in int) {
	iter := NewIterator(p.g, in)
	p.visited[in] = true
	for v := iter.begin(); !iter.end(); v = iter.next() {
		if !p.visited[v] {
			p.from[v] = in
			p.dfs(v)
		}
	}
}
