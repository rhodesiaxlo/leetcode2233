package heap

import (
	"fmt"
)

// 背景：对于数据结构比较复杂， 或者需要保存现场的情况, 把源数据搞得乱七八槽肯定不满足要求
// indexMaxHeap
type IndexMaxHeap struct {
	data   []Item
	index  []int
	eleNum int

	// 为了减少内存扩充，需要增加一个
	// cap int
}

func NewIndexMaxHeap(cap int) IndexMaxHeap {
	// 构造一个容量为 n 的 IndexMaxHeap
	mh := IndexMaxHeap{}
	mh.Init(cap)
	return mh
}

// 插入的效率没有原地排序效率高
func NewIndexMaxHeapByInsert(items []Item) IndexMaxHeap {
	mh := IndexMaxHeap{}
	sz := len(items)
	for i := 0; i < sz; i++ {
		// 保留备份， 能否用指针呢？ 不能
		mh.Insert(items[i])
	}

	return mh
}

func NewIndexMaxHeapByHeapify(items []Item) IndexMaxHeap {
	mh := IndexMaxHeap{}
	mh.Heapify(items)
	return mh
}

func (m *IndexMaxHeap) Init(cap int) {
	m.data = make([]Item, cap+1)
	m.index = make([]int, cap+1)
	m.eleNum = 0
}

func (m *IndexMaxHeap) Size() int {
	return m.eleNum
}

func (m *IndexMaxHeap) Empty() bool {
	return m.eleNum == 0
}

// 注意 insert 可以可以是追加， 也可以是在某一个索引的位置增加一个元素
// 这里我们为了简便， 理解为在末尾增加一个元素, 而修改我们用 change 来实现
func (m *IndexMaxHeap) Insert(x Item) {
	if len(m.data) == 0 {
		// 插入一个空元素， 使其 1 based
		m.data = append(m.data, Item{})
		m.index = append(m.index, 0)
	}
	m.data = append(m.data, x)
	m.eleNum += 1
	m.index = append(m.index, m.eleNum)
	m.shiftUp(m.eleNum)
}

func (m *IndexMaxHeap) Max() Item {
	return m.data[m.index[1]]
}

func (m *IndexMaxHeap) shiftUp(nthNum int) {
	for nthNum > 1 && m.data[m.index[nthNum]].Gt(m.data[m.index[nthNum/2]]) {
		m.index[nthNum], m.index[nthNum/2] = m.index[nthNum/2], m.index[nthNum]
		nthNum /= 2
	}
}

func (m *IndexMaxHeap) ExtractMax() (Item, error) {
	if m.Empty() {
		return Item{}, fmt.Errorf("no more element")
	}
	max := m.data[m.index[1]]
	m.index[1], m.index[m.eleNum] = m.index[m.eleNum], m.index[1]
	m.eleNum -= 1
	m.shiftDown(1, m.eleNum)
	return max, nil
}

func (m *IndexMaxHeap) shiftDown(nthNum, n int) {
	for 2*nthNum <= n {
		j := 2 * nthNum
		if j+1 <= n && m.data[m.index[j+1]].Gt(m.data[m.index[j]]) {
			j += 1
		}

		if m.data[m.index[j]].Gt(m.data[m.index[nthNum]]) {
			m.index[j], m.index[nthNum] = m.index[nthNum], m.index[j]
		}
		nthNum = j
	}
}

func (m *IndexMaxHeap) Heapify(items []Item) {
	sz := len(items)
	m.data = make([]Item, sz+1)
	m.index = make([]int, sz+1)
	copy(m.data[1:], items)
	for i := 1; i <= sz; i++ {
		m.index[i] = i
	}
	m.eleNum = sz

	for lastParent := m.eleNum / 2; lastParent >= 1; lastParent-- {
		m.shiftDown(lastParent, m.eleNum)
	}
}

func (m *IndexMaxHeap) SortAsc() {
	n := m.eleNum
	for i := m.eleNum; i >= 1; i-- {
		m.index[i], m.index[1] = m.index[1], m.index[i]
		n -= 1
		m.shiftDown(1, n)
	}
}

func (m *IndexMaxHeap) String() string {
	return "not implement yet"
}

func (m *IndexMaxHeap) Display() {
}

func (m *IndexMaxHeap) GetArr() []Item {
	ans := make([]Item, m.eleNum)
	for i := 0; i < m.eleNum; i++ {
		ans[i] = m.data[m.index[i+1]]
	}
	return ans
}

func (m *IndexMaxHeap) Change(i int, item Item) error {
	// i 是外部看起来的 index 是， 0-based
	i += 1
	if i < 1 || i > m.eleNum {
		return fmt.Errorf("index out of bound")
	}
	m.data[i] = item
	// 注意这里的找到元素
	for k := 1; k < m.eleNum; k++ {
		if m.index[k] == i {
			m.shiftUp(k)
			m.shiftDown(k, m.eleNum)
			break
		}
	}
	return nil
}

func (m *IndexMaxHeap) GetItem(i int) (Item, error) {
	// i 是外部看起来的 index 是， 0-based
	i += 1
	if i < 1 || i > m.eleNum {
		return Item{}, fmt.Errorf("index out of bound")
	}

	//return m.data[m.index[i]], nil
	return m.data[i], nil
}
