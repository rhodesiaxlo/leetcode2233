package graph

import (
	"fmt"
	"leetcode/leetcode/structure/stack"
	"strconv"
)

// dijkstra 只能处理没有负权边的, 新加入负权边， 会把整个 dijkstra 算法思想搞乱， 所以应该负权边的情况，我们有另外一个算法
// bellmenford, 想法是，对于任意一个顶点 des  和源点 src ，最多只需要进行 v-1 轮松弛操作就可以找到最短路径的那一个点，如果说
// 图中有负权换 1 <-(1)-  4  1-(-4)-> 4 等的，第 v 轮依然可以进行松弛
type BellmenFord struct {
	ipq IndexMinHeapFloat64
	g   interface{}
	v   int
	src int // source
	//marked []bool
	from             []*edge
	disTo            []float64
	hasNegativeCycle bool
}

func NewBellmenFord(g interface{}, v int, src int) BellmenFord {
	d := BellmenFord{
		ipq:              NewIndexMinHeapFloat64(v),
		g:                g,
		v:                v,
		src:              src,
		from:             nil,
		disTo:            nil,
		hasNegativeCycle: false,
	}
	d.Init()
	return d
}

func (d *BellmenFord) Init() {
	//d.marked = make([]bool, d.v)
	d.disTo = make([]float64, d.v)
	// 这个体检很有必要
	//for i := 0; i < len(d.disTo); i++ {
	//	d.disTo[i] = math.MaxFloat64
	//}
	d.from = make([]*edge, d.v)

	//bellmen-ford
	for pass := 1; pass < d.v; pass++ { // 进行 v-1 次
		for i := 0; i < d.v; i++ { // 每次对所有顶点进行松弛 relaxation 操作
			tmpIter := NewWtIterator(d.g, i)
			for e := tmpIter.begin(); !tmpIter.end(); e = tmpIter.next() {
				// 注意这里在有向图中 e.V() e.W() 是有具体含义的 e.V() 代表 from  e.W() 代表 to
				if d.from[e.W()] == nil || d.disTo[e.V()]+e.wt < d.disTo[e.W()] { // 如果没有走过， 或者绕道后更短
					d.from[e.W()] = e
					d.disTo[e.W()] = d.disTo[e.V()] + e.wt
				}
			}
		}
	}

	// 在进行一轮
	d.DetectNegativeCycle()

}

func (d *BellmenFord) DetectNegativeCycle() bool {
	for i := 0; i < d.v; i++ {
		tmpIter := NewWtIterator(d.g, i)
		for e := tmpIter.begin(); !tmpIter.end(); e = tmpIter.next() {
			if d.from[e.W()] == nil || d.disTo[e.V()]+e.wt < d.disTo[e.W()] { // 如果没有走过， 或者绕道后更短
				d.hasNegativeCycle = true
				return true
			}
		}
	}
	return false
}
func (d *BellmenFord) ShortestPahTo(x int) float64 {
	return d.disTo[x]
}

func (d *BellmenFord) ShortestPath(x int) {
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
