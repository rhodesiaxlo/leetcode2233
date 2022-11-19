package heap

import (
	"fmt"
	"leetcode/leetcode/structure"
)

type MinHeap struct {
	arr    []int
	eleNum int
}

// 插入的效率没有原地排序效率高
func NewMinHeapByInsert(nums []int) MinHeap {
	mh := MinHeap{}
	sz := len(nums)
	for i := 0; i < sz; i++ {
		mh.Insert(nums[i])
	}

	return mh
}

func NewMinHeapByHeapify(nums []int) MinHeap {
	mh := MinHeap{}
	mh.Heapify(nums)
	return mh
}

func (m *MinHeap) Size() int {
	return m.eleNum
}

func (m *MinHeap) Empty() bool {
	return m.eleNum == 0
}

func (m *MinHeap) Insert(x int) {
	if len(m.arr) == 0 {
		// 插入一个空元素， 使其 1 based
		m.arr = append(m.arr, 0)
	}
	m.arr = append(m.arr, x)
	m.eleNum += 1

	m.shiftUp(len(m.arr) - 1)
}

func (m *MinHeap) Mix() int {
	return m.arr[1]
}

func (m *MinHeap) shiftUp(nthNum int) {
	for nthNum > 1 && m.arr[nthNum] < m.arr[nthNum/2] {
		m.arr[nthNum], m.arr[nthNum/2] = m.arr[nthNum/2], m.arr[nthNum]
		nthNum /= 2
	}
}

func (m *MinHeap) ExtractMin() (int, error) {
	if m.Empty() {
		return 0, fmt.Errorf("no more element")
	}
	min := m.arr[1]
	m.arr[1], m.arr[m.eleNum] = m.arr[m.eleNum], m.arr[1]
	m.eleNum -= 1
	m.shiftDown(1, m.eleNum)
	return min, nil
}

func (m *MinHeap) shiftDown(nthNum, n int) {
	for 2*nthNum <= n {
		j := 2 * nthNum
		if j+1 <= n && m.arr[j+1] < m.arr[j] {
			j += 1
		}

		if m.arr[j] < m.arr[nthNum] {
			m.arr[j], m.arr[nthNum] = m.arr[nthNum], m.arr[j]
		}
		nthNum = j
	}
}

// 优化点1. heapify
func (m *MinHeap) Heapify(nums []int) {
	sz := len(nums)
	m.arr = make([]int, sz+1)
	copy(m.arr[1:], nums)
	m.eleNum = sz

	for lastParent := m.eleNum / 2; lastParent >= 1; lastParent-- {
		m.shiftDown(lastParent, m.eleNum)
	}
}

func (m *MinHeap) SortDesc() {
	n := m.eleNum
	for i := m.eleNum; i >= 1; i-- {
		m.arr[i], m.arr[1] = m.arr[1], m.arr[i]
		n -= 1
		m.shiftDown(1, n)
	}
}

func (m *MinHeap) String() string {
	return "not implement yet"
}

func (m *MinHeap) Display() {
	structure.PrintComTree(m.arr)
}

func (m *MinHeap) GetArr() []int {
	return m.arr[1:]
}

