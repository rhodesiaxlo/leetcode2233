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
type MaxHeap struct {
	arr []int
}

func NewMaxHeap(nums []int) MaxHeap {
	mh := MaxHeap{}
	mh.Heapify(nums)
	return mh
}
/**
返回元素个数，而不是低层数组的长度
 */
func (m *MaxHeap) Size() int {
	return len(m.arr) - 1
}

/**
判断是否还有元素
 */
func (m *MaxHeap) Empty() bool {
	return m.Size() == 1
}

func (m *MaxHeap) Heapify(nums []int) {
	eleSize := len(nums)
	m.arr = make([]int, eleSize+1)
	for i:=1;i< eleSize+1;i++ {
		m.arr[i] = nums[i-1]
	}

	// shift down
	for lastParent := eleSize/2;lastParent >=1;lastParent-- {
		m.shiftDown(lastParent)
	}
}

func (m *MaxHeap) Insert(x int) {
	m.arr = append(m.arr, x)
	m.shiftUp(len(m.arr)-1)
}

func (m *MaxHeap) Max() int {
	return m.arr[1]
}

func (m *MaxHeap) ExtractMax() (int, error) {
	if m.Empty() {
		return 0, fmt.Errorf("no ele any more")
	}
	max := m.arr[1]
	m.arr[1], m.arr[len(m.arr)-1] = m.arr[len(m.arr)-1], m.arr[1]
	m.arr = append([]int{}, m.arr[:len(m.arr)-1]...)
	m.shiftDown(1)
	return max, nil
}

func (m *MaxHeap) String() string {
	// todo
	return ""
}

func (m *MaxHeap) Display() {
	structure.PrintComTree(m.arr)
}

// 假设 maxHeap, 在末尾增加了一个元素，则需要移动节点使其重新有序， 这个过程
// 我们称为  shiftUp
// index 代表第几个元素，不是 array 的 index
func (m *MaxHeap) shiftUp(nthNum int) {
	for nthNum > 1 && m.arr[nthNum] > m.arr[nthNum/2] {
		m.arr[nthNum], m.arr[nthNum/2] = m.arr[nthNum/2], m.arr[nthNum]
		nthNum /= 2
	}
}

// 假设 maxHeap, 则 arr[1] 一定是最大的元素，那么，如果拿走了这个元素之后，如何才能是 maxHeap 重新回归到
// 有序的状态呢
func (m *MaxHeap) shiftDown(nthNum int) {
	sz := len(m.arr)
	maxEleIndex := sz -1
	for 2* nthNum <=maxEleIndex {
		j := 2*nthNum
		if j+1 <=maxEleIndex && m.arr[j+1] > m.arr[j]{
			j+=1
		}

		if m.arr[j] > m.arr[nthNum] {
			m.arr[j], m.arr[nthNum] = m.arr[nthNum], m.arr[j]
		}
		nthNum = j
	}
}

func (m *MaxHeap) Arr() []int {
	return m.arr
}



// ------------------- min heap ------------------
type MinHeap struct {
	arr []int
}

func NewMinHeap(nums []int) MinHeap {
	mh:= MinHeap{}
	mh.Heapify(nums)
	return mh
}

/**
返回元素个数，而不是低层数组的长度
*/
func (m *MinHeap) Size() int {
	return len(m.arr) - 1
}

/**
判断是否还有元素
*/
func (m *MinHeap) Empty() bool {
	return m.Size() == 1
}

func (m *MinHeap) Heapify(nums []int) {
	eleSize := len(nums)
	m.arr = make([]int, eleSize+1)
	for i:=1;i< eleSize+1;i++ {
		m.arr[i] = nums[i-1]
	}

	// shift down
	for lastParent := eleSize/2;lastParent >=1;lastParent-- {
		m.shiftDown(lastParent)
	}
}

func (m *MinHeap) Insert(x int) {
	m.arr = append(m.arr, x)
	m.shiftUp(len(m.arr)-1)
}

func (m *MinHeap) Min() int {
	return m.arr[1]
}

func (m *MinHeap) ExtractMin() (int, error) {
	if m.Empty() {
		return 0, fmt.Errorf("no ele any more")
	}
	max := m.arr[1]
	m.arr[1], m.arr[len(m.arr)-1] = m.arr[len(m.arr)-1], m.arr[1]
	m.arr = append([]int{}, m.arr[:len(m.arr)-1]...)
	m.shiftDown(1)
	return max, nil
}

func (m *MinHeap) String() string {
	// todo
	return ""
}

func (m *MinHeap) Display() {
	structure.PrintComTree(m.arr)
}

// 假设 maxHeap, 在末尾增加了一个元素，则需要移动节点使其重新有序， 这个过程
// 我们称为  shiftUp
// index 代表第几个元素，不是 array 的 index
func (m *MinHeap) shiftUp(nthNum int) {
	for nthNum > 1 && m.arr[nthNum] < m.arr[nthNum/2] {
		m.arr[nthNum], m.arr[nthNum/2] = m.arr[nthNum/2], m.arr[nthNum]
		nthNum /= 2
	}
}

// 假设 maxHeap, 则 arr[1] 一定是最大的元素，那么，如果拿走了这个元素之后，如何才能是 maxHeap 重新回归到
// 有序的状态呢
func (m *MinHeap) shiftDown(nthNum int) {
	sz := len(m.arr)
	maxEleIndex := sz -1
	for 2* nthNum <=maxEleIndex {
		j := 2*nthNum
		if j+1 <=maxEleIndex && m.arr[j+1] < m.arr[j]{
			j+=1
		}

		if m.arr[j] < m.arr[nthNum] {
			m.arr[j], m.arr[nthNum] = m.arr[nthNum], m.arr[j]
		}
		nthNum = j
	}
}

func (m *MinHeap) Arr() []int {
	return m.arr
}

