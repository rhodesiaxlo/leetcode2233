package graph

import "fmt"

/*
lazyprime 算法的时间复杂度是ElogE级别的，我们可以通过一直寻找每个顶点最短的那条边来寻找横切片的方式，来生成最下生成树
使用 indexMinHeap

已0 为起点， 遍历0 的临边，加入到 indexMinHeap
在 最小生出树的节点 < n - 1
	找出最短边， change
	遍历最短边的领边，如果新的横切边比现有的小，则更新权值从而更新最小堆
*/

type PrimeMST struct {
	g      interface{}
	ipq    IndexMixHeapEdge
	marked []bool
	v      int
	ans    []*edge
	mstWt  float64
	toEdge []*edge // 源数据
}

/**
注意 g 必须是实现interface的对象
*/
func NewPrimeMst(g interface{}, v int) PrimeMST {
	iqp := NewIndexMixHeapEdge(v)

	pmst := PrimeMST{
		g:      g,
		ipq:    iqp,
		marked: nil,
		v:      v,
		ans:    nil,
		mstWt:  0,
		toEdge: nil,
	}

	pmst.Init()
	return pmst
}

func (p *PrimeMST) Init() {
	p.marked = make([]bool, p.v)
	p.toEdge = make([]*edge, p.v)

	p.Visit(0)
	for !p.ipq.Empty() {
		minIndex, _ := p.ipq.ExtractMinIndex()
		if p.toEdge[minIndex] != nil {
			// 确实是存在这条横切边的
			p.ans = append(p.ans, p.toEdge[minIndex])
			p.mstWt += p.toEdge[minIndex].wt
			p.Visit(minIndex)
		} else {
			continue
		}
	}
	// 最小生成树的边一定等于顶点数-1
	// todo..
}

func (p *PrimeMST) Visit(i int) {
	// 避免成环
	if !p.marked[i] {
		p.marked[i] = true

		iter := NewWtIterator(p.g, i)
		for v := iter.begin(); !iter.end(); v = iter.next() {
			// 如果没有加入到横切边的数组中
			otherV, _ := v.Other(i)

			if !p.marked[otherV] { // 避免成环
				if p.toEdge[otherV] == nil {
					p.toEdge[otherV] = v
					p.ipq.Insert(otherV, v) // 原来的数字是0开始排列的， 且数字内容和序号是重合的， 进入 ipq 后， 需要从1开始排，最小为1
				} else if p.toEdge[otherV].wt > v.wt {
					p.toEdge[otherV] = v
					p.ipq.Change(otherV, v)
				}
			}
		}
	}
}

func (p *PrimeMST) Show() {
	for i := 0; i < len(p.ans); i++ {
		fmt.Printf("%d to %d, wt :%3.2f \n", p.ans[i].v, p.ans[i].w, p.ans[i].wt)
	}
}
