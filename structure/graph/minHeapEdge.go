package graph

import (
	"fmt"
)

type MinHeapEdge struct {
	arr    []*edge
	eleNum int
}

func NewMinHeapEdge(cap int) MinHeapEdge {
	mhp := MinHeapEdge{
		arr:    nil,
		eleNum: 0,
	}

	mhp.arr = make([]*edge, cap+1)
	return mhp
}

// 插入的效率没有原地排序效率高
func NewMinHeapEdgeByInsert(nums []*edge) MinHeapEdge {
	pq := MinHeapEdge{}
	sz := len(nums)
	for i := 0; i < sz; i++ {
		pq.Insert(nums[i])
	}

	return pq
}

func NewMinHeapEdgeByHeapify(nums []*edge) MinHeapEdge {
	mh := MinHeapEdge{}
	mh.Heapify(nums)
	return mh
}

func (m *MinHeapEdge) Size() int {
	return m.eleNum
}

func (m *MinHeapEdge) Empty() bool {
	return m.eleNum == 0
}

func (m *MinHeapEdge) Insert(x *edge) {
	if len(m.arr) == 0 {
		// 插入一个空元素， 使其 1 based
		m.arr = append(m.arr, nil)
	}

	//m.arr = append(m.arr, x) // 因为最后一个元素
	//m.eleNum += 1
	m.arr[m.eleNum+1] = x
	m.eleNum++

	//m.shiftUp(len(m.arr) - 1) // 定义完整的重要性
	m.shiftUp(m.eleNum)

}

func (m *MinHeapEdge) Min() *edge {
	return m.arr[1]
}

func (m *MinHeapEdge) shiftUp(nthNum int) {
	for nthNum > 1 && m.arr[nthNum].wt < m.arr[nthNum/2].wt {
		m.arr[nthNum], m.arr[nthNum/2] = m.arr[nthNum/2], m.arr[nthNum]
		nthNum /= 2
	}
}

func (m *MinHeapEdge) ExtractMin() (*edge, error) {
	if m.Empty() {
		return nil, fmt.Errorf("no more element")
	}
	min := m.arr[1]
	m.arr[1], m.arr[m.eleNum] = m.arr[m.eleNum], m.arr[1]
	m.eleNum -= 1
	//m.arr = append([]*edge(nil), m.arr[:m.eleNum+1]...) // 还是删除比较好
	m.shiftDown(1, m.eleNum)
	return min, nil
}

func (m *MinHeapEdge) shiftDown(nthNum, n int) {
	for 2*nthNum <= n {
		j := 2 * nthNum
		if j+1 <= n && m.arr[j+1].wt < m.arr[j].wt {
			j += 1
		}

		if m.arr[j].wt < m.arr[nthNum].wt {
			m.arr[j], m.arr[nthNum] = m.arr[nthNum], m.arr[j]
		}
		nthNum = j
	}
}

// 优化点1. heapify
func (m *MinHeapEdge) Heapify(nums []*edge) {
	sz := len(nums)
	m.arr = make([]*edge, sz+1)
	copy(m.arr[1:], nums)
	m.eleNum = sz

	for lastParent := m.eleNum / 2; lastParent >= 1; lastParent-- {
		m.shiftDown(lastParent, m.eleNum)
	}
}

func (m *MinHeapEdge) SortDesc() {
	n := m.eleNum
	for i := m.eleNum; i >= 1; i-- {
		m.arr[i], m.arr[1] = m.arr[1], m.arr[i]
		n -= 1
		m.shiftDown(1, n)
	}
}

func (m *MinHeapEdge) String() string {
	return "not implement yet"
}

func (m *MinHeapEdge) Display() {

}

func (m *MinHeapEdge) GetArr() []*edge {
	return m.arr[1:]
}
