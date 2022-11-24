package graph

import (
	"fmt"
	"leetcode/leetcode/structure/stack"
	"math"
	"strconv"
)

/*
迪杰克斯拉最短路径算法
*/

type DijkStra struct {
	ipq    IndexMinHeapFloat64
	g      interface{}
	v      int
	src    int // source
	marked []bool
	from   []*edge
	disTo  []float64
}

func NewDijkStra(g interface{}, v int, src int) DijkStra {
	d := DijkStra{
		ipq:    NewIndexMinHeapFloat64(v),
		g:      g,
		v:      v,
		src:    src,
		marked: nil,
		from:   nil,
		disTo:  nil,
	}
	d.Init()
	return d
}

func (d *DijkStra) Init() {
	d.marked = make([]bool, d.v)
	d.disTo = make([]float64, d.v)
	// 这个体检很有必要
	for i := 0; i < len(d.disTo); i++ {
		d.disTo[i] = math.MaxFloat64
	}
	d.from = make([]*edge, d.v)

	// dijkstra
	d.marked[d.src] = true
	d.disTo[d.src] = 0

	d.ipq.Insert(d.src, 0)
	for !d.ipq.Empty() {
		tmpMinIndex, _ := d.ipq.ExtractMinIndex()
		tmpIter := NewWtIterator(d.g, tmpMinIndex)
		for v := tmpIter.begin(); !tmpIter.end(); v = tmpIter.next() {
			tmpNextVertex, _ := v.Other(tmpMinIndex)
			if !d.marked[tmpNextVertex] { // 没有成环
				// 如果之前没有计算 tmpNextVert 的距离(没有访问过)
				// 或者 relaxation 之后更小的
				if d.disTo[tmpNextVertex] == 0 || d.disTo[tmpMinIndex]+v.wt < d.disTo[tmpNextVertex] {
					d.disTo[tmpNextVertex] = d.disTo[tmpMinIndex] + v.wt
					d.from[tmpNextVertex] = v
					// 然后需要把新的 single source 距离加入到 ipq 中去参与下一轮
					if d.ipq.Contain(tmpNextVertex) {
						d.ipq.Change(tmpNextVertex, d.disTo[tmpNextVertex])
					} else {
						d.ipq.Insert(tmpNextVertex, d.disTo[tmpNextVertex])
					}
				}
			}
		}
	}
}

func (d *DijkStra) HasPathTo(w int) bool {
	return d.marked[w]
}

func (d *DijkStra) ShortestPahTo(x int) float64 {
	return d.disTo[x]
}

func (d *DijkStra) ShortestPath(x int) {
	// 在有权图里面 edge 的 v 是 from  w 是 to
	// 所以

	s := stack.NewStackStr(d.v + 10)
	e := d.from[x]
	s.Push(strconv.Itoa(x))
	for e.V() != d.src {
		s.Push(strconv.Itoa(e.V()))
		e = d.from[e.v]
	}
	s.Push(strconv.Itoa(d.src))

	fmt.Printf("shortest path to %d\n", x)
	for s.Size() > 1 {
		tmpPop := s.Pop()
		fmt.Print(tmpPop, "-->")
	}
	fmt.Println(s.Pop())
}
