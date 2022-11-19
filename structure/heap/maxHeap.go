package heap

import (
	"fmt"
	"leetcode/leetcode/structure"
)

// heap
// 1. complete tree
// 2. root is level 0 ,  level is the same as depth , height is the counter part
//    a leaf node will the the height of 0, so a heap of the level k can have the largest
//    number of nodes of 2*k +1
//    complete tree give us the advantage to use an array to represent a heap
//    parent index of k have left child of 2*k and right child index 2*k + 1
//    if index is 0-based, then node at level l index is 2*l
//    if index is 1-based, then node at level l index is 2*l+1
// 4. max node number is 2*l+1
// 5. a heap with nodes number of n, the last parent nodes index is n/2, if is zero-based  it is n/2-1

// basic operation
// 1. heapify
// 2. insert
// 3. delete
// 4. size
// 5. emtpy
// 6. extractMin
// 7. extractMax

// step1
// 1. shiftup
// 2. shiftdown
// 3. insert
// 4. heapify
// 5. extractMin / extractMax  displace / inplace
// 6. print

type MaxHeap struct {
	arr    []int
	eleNum int
}

// 插入的效率没有原地排序效率高
func NewMaxHeapByInsert(nums []int) MaxHeap {
	mh := MaxHeap{}
	sz := len(nums)
	for i := 0; i < sz; i++ {
		mh.Insert(nums[i])
	}

	return mh
}

func NewMaxHeapByHeapify(nums []int) MaxHeap {
	mh := MaxHeap{}
	mh.Heapify(nums)
	return mh
}

func (m *MaxHeap) Size() int {
	return m.eleNum
}

func (m *MaxHeap) Empty() bool {
	return m.eleNum == 0
}

func (m *MaxHeap) Insert(x int) {
	if len(m.arr) == 0 {
		// 插入一个空元素， 使其 1 based
		m.arr = append(m.arr, 0)
	}
	m.arr = append(m.arr, x)
	m.eleNum += 1

	m.shiftUp(len(m.arr) - 1)
}

func (m *MaxHeap) Max() int {
	return m.arr[1]
}

func (m *MaxHeap) shiftUp(nthNum int) {
	for nthNum > 1 && m.arr[nthNum] > m.arr[nthNum/2] {
		m.arr[nthNum], m.arr[nthNum/2] = m.arr[nthNum/2], m.arr[nthNum]
		nthNum /= 2
	}
}

func (m *MaxHeap) ExtractMax() (int, error) {
	if m.Empty() {
		return 0, fmt.Errorf("no more element")
	}
	max := m.arr[1]
	m.arr[1], m.arr[m.eleNum] = m.arr[m.eleNum], m.arr[1]
	m.eleNum -= 1
	m.shiftDown(1, m.eleNum)
	return max, nil
}

func (m *MaxHeap) shiftDown(nthNum, n int) {
	for 2*nthNum <= n {
		j := 2 * nthNum
		if j+1 <= n && m.arr[j+1] > m.arr[j] {
			j += 1
		}

		if m.arr[j] > m.arr[nthNum] {
			m.arr[j], m.arr[nthNum] = m.arr[nthNum], m.arr[j]
		}
		nthNum = j
	}
}

// 优化点1. heapify
func (m *MaxHeap) Heapify(nums []int) {
	sz := len(nums)
	m.arr = make([]int, sz+1)
	copy(m.arr[1:], nums)
	m.eleNum = sz

	for lastParent := m.eleNum / 2; lastParent >= 1; lastParent-- {
		m.shiftDown(lastParent, m.eleNum)
	}
}

func (m *MaxHeap) SortAsc() {
	n := m.eleNum
	for i := m.eleNum; i >= 1; i-- {
		m.arr[i], m.arr[1] = m.arr[1], m.arr[i]
		n -= 1
		m.shiftDown(1, n)
	}
}

func (m *MaxHeap) String() string {
	return "not implement yet"
}

func (m *MaxHeap) Display() {
	structure.PrintComTree(m.arr)
}

func (m *MaxHeap) GetArr() []int {
	return m.arr[1:]
}
