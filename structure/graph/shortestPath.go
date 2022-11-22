package graph

import (
	"fmt"
	"leetcode/leetcode/structure/queue"
	stack2 "leetcode/leetcode/structure/stack"
)

type ShortestPath struct {
	src     int
	visited []bool
	from    []int
	ord     []int
	g       interface{}
	v       int
}

func NewShortestPath(graph interface{}, src int, vNumber int) ShortestPath {
	p := ShortestPath{
		src:     src,
		visited: nil,
		from:    nil,
		g:       graph,
		v:       vNumber,
	}
	p.Init()
	return p
}

func (p *ShortestPath) Init() {
	p.visited = make([]bool, p.v)
	p.from = make([]int, p.v)
	p.ord = make([]int, p.v)
	for i := 0; i < p.v; i++ {
		p.from[i] = -1
		p.ord[i] = -1
	}

	// 求解最短路径
	q := queue.NewQueue()
	q.PushBack(p.src) // back 和 front 的概念
	p.ord[p.src] = 0
	p.visited[p.src] = true
	for q.Len() > 0 {
		front := q.Front().(int)
		iterator := NewIterator(p.g, front)
		for v := iterator.begin(); !iterator.end(); v = iterator.next() {
			if !p.visited[v] {
				p.ord[v] = p.ord[front] + 1
				p.from[v] = front
				p.visited[v] = true
				q.PushBack(v)
			}
		}
		q.PopFront()
	}
}

func (p *ShortestPath) CalcPath() {
	p.dfs(p.src)
}

func (p *ShortestPath) HasPath(des int) bool {
	return p.visited[des]
}

func (p *ShortestPath) GetPath(des int) []rune {
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

func (p *ShortestPath) ShowPath(in int) {
	fmt.Println(p.GetPath(in))
}

func (p *ShortestPath) dfs(in int) {
	iter := NewIterator(p.g, in)
	p.visited[in] = true
	for v := iter.begin(); !iter.end(); v = iter.next() {
		if !p.visited[v] {
			p.from[v] = in
			p.dfs(v)
		}
	}
}

func (p *ShortestPath) Length(j int) int {
	return p.ord[j]
}
