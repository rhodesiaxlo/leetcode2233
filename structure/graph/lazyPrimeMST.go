package graph

import "fmt"

/*
在从无权图, 遍历算法 -> 有权图 -> 寻找带权图的最小路径 (minimum span tree)
目前有 prime 算法
*/

type LazyPrimeMST struct {
	g      interface{}
	pq     MinHeapEdge
	marked []bool
	v      int
	ans    []*edge
	mstWt  float64
}

func newLazyPrimeMst(g interface{}, v int) LazyPrimeMST {
	mst := LazyPrimeMST{
		g:      g,
		pq:     MinHeapEdge{},
		marked: nil,
		v:      v,
	}
	mst.Init()
	return mst
}

/**
算法复杂度分析 for 循环 e
2个操作 extractMin logE
visit eloge
 */
func (l *LazyPrimeMST) Init() {
	l.marked = make([]bool, l.v)

	l.Visit(0)
	for !l.pq.Empty() {
		tmpMin, _ := l.pq.ExtractMin()
		//for l.marked[tmpMin.v] && l.marked[tmpMin.w] {
		if l.marked[tmpMin.v] == l.marked[tmpMin.w] {
			//tmpMin, _ = l.pq.ExtractMin()
			continue
		}

		l.ans = append(l.ans, tmpMin)
		l.mstWt += tmpMin.wt
		if !l.marked[tmpMin.v] {
			l.Visit(tmpMin.v)
		} else {
			l.Visit(tmpMin.w)
		}
	}

	// 最小生成树的边一定等于顶点数-1
	// todo..
}

func (l *LazyPrimeMST) MsgEdges() []*edge {
	return l.ans
}

func (l *LazyPrimeMST) Show() {
	for i := 0; i < len(l.ans); i++ {
		fmt.Printf("%d to %d, wt :%f \n", l.ans[i].v, l.ans[i].w, l.ans[i].wt)
	}
}

func (l *LazyPrimeMST) MsgWt() float64 {
	return l.mstWt
}

func (l *LazyPrimeMST) Visit(i int) {
	if !l.marked[i] {
		l.marked[i] = true
	}

	iter := NewWtIterator(l.g, i)
	for v := iter.begin(); !iter.end(); v = iter.next() {
		otherV, _ := v.Other(i)
		if !l.marked[otherV] {
			l.pq.Insert(v)
		}
	}
}
