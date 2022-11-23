package graph

import "fmt"

type IndexMixHeapEdge struct {
	data   []*edge
	index  []int
	eleNum int
}

func NewIndexMixHeapEdge(cap int) IndexMixHeapEdge {
	mh := IndexMixHeapEdge{}
	mh.Init(cap)
	return mh
}

// 插入的效率没有原地排序效率高
func NewIndexMaxHeapByInsert(items []*edge) IndexMixHeapEdge {
	mh := IndexMixHeapEdge{}
	sz := len(items)
	for i := 0; i < sz; i++ {
		// 保留备份， 能否用指针呢？ 不能
		mh.Insert(mh.eleNum, items[i])
	}

	return mh
}

func NewIndexMaxHeapByHeapify(items []*edge) IndexMixHeapEdge {
	mh := IndexMixHeapEdge{}
	mh.Heapify(items)
	return mh
}

func (m *IndexMixHeapEdge) Init(cap int) {
	m.data = make([]*edge, cap+1) // 需要考虑是否要有顺序插入, 如果是， 需要 length  加上预分配空间
	m.index = make([]int, cap+1)
	//for i := 0; i < len(m.index); i++ {
	//	m.index[i] = i
	//}
	m.eleNum = 0
}

func (m *IndexMixHeapEdge) Size() int {
	return m.eleNum
}

func (m *IndexMixHeapEdge) Empty() bool {
	return m.eleNum == 0
}

// 注意 insert 可以可以是追加， 也可以是在某一个索引的位置增加一个元素
// 这里我们为了简便， 理解为在末尾增加一个元素, 而修改我们用 change 来实现
func (m *IndexMixHeapEdge) Insert(i int, x *edge) {
	//if len(m.data) == 0 {
	//	// 插入一个空元素， 使其 1 based
	//	m.data = append(m.data, nil)
	//	m.index = append(m.index, 0)
	//}
	m.data[i+1] = x // 有意思， 这里稀疏插入，  index 密集整理
	m.index[m.eleNum+1] = i + 1 // 元素怎么变都没有关系， 主要是 index , 从而维护了堆的定义
	m.eleNum += 1
	m.shiftUp(m.eleNum)
}

func (m *IndexMixHeapEdge) Min() *edge {
	return m.data[m.index[1]]
}

func (m *IndexMixHeapEdge) shiftUp(nthNum int) {
	for nthNum > 1 && (m.data[m.index[nthNum/2]] == nil || m.data[m.index[nthNum]].wt < (m.data[m.index[nthNum/2]].wt)) {
		m.index[nthNum], m.index[nthNum/2] = m.index[nthNum/2], m.index[nthNum]
		nthNum /= 2
	}
}

func (m *IndexMixHeapEdge) ExtractMin() (*edge, error) {
	if m.Empty() {
		return nil, fmt.Errorf("no more element")
	}
	max := m.data[m.index[1]]
	m.index[1], m.index[m.eleNum] = m.index[m.eleNum], m.index[1]
	m.eleNum -= 1
	m.shiftDown(1, m.eleNum)
	return max, nil
}

func (m *IndexMixHeapEdge) ExtractMinIndex() (int, error) {
	if m.Empty() {
		return 0, fmt.Errorf("no more element")
	}
	//min := m.data[m.index[1]]
	tmpMinIndex := m.index[1] // 数据的索引是应该不为空的
	m.index[1], m.index[m.eleNum] = m.index[m.eleNum], m.index[1]
	m.eleNum -= 1
	m.shiftDown(1, m.eleNum)
	return tmpMinIndex - 1, nil
}

func (m *IndexMixHeapEdge) shiftDown(nthNum, n int) {
	for 2*nthNum <= n {
		j := 2 * nthNum
		if j+1 <= n &&
			m.data[m.index[j]] != nil &&
			m.data[m.index[j+1]] != nil &&
			m.data[m.index[j+1]].wt < (m.data[m.index[j]].wt) {
			j += 1
		}

		if m.data[m.index[nthNum]] == nil ||
			(m.data[m.index[j]] != nil && m.data[m.index[j]].wt < m.data[m.index[nthNum]].wt) {
			m.index[j], m.index[nthNum] = m.index[nthNum], m.index[j]
		}
		nthNum = j
	}
}

func (m *IndexMixHeapEdge) Heapify(items []*edge) {
	sz := len(items)
	m.data = make([]*edge, sz+1)
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

func (m *IndexMixHeapEdge) SortAsc() {
	n := m.eleNum
	for i := m.eleNum; i >= 1; i-- {
		m.index[i], m.index[1] = m.index[1], m.index[i]
		n -= 1
		m.shiftDown(1, n)
	}
}

func (m *IndexMixHeapEdge) String() string {
	return "not implement yet"
}

func (m *IndexMixHeapEdge) Display() {
}

func (m *IndexMixHeapEdge) GetArr() []*edge {
	ans := make([]*edge, m.eleNum)
	for i := 0; i < m.eleNum; i++ {
		ans[i] = m.data[m.index[i+1]]
	}
	return ans
}

/**
i 代表数据索引号
*/
func (m *IndexMixHeapEdge) Change(i int, item *edge) error {
	// i 是外部看起来的 index 是， 0-based
	i += 1
	m.data[i] = item
	// 注意这里的找到元素
	for k := 1; k < len(m.index); k++ {
		if m.index[k] == i {
			m.shiftUp(k)
			m.shiftDown(k, m.eleNum)
			break
		}
	}
	return nil
}

func (m *IndexMixHeapEdge) GetItem(i int) (*edge, error) {
	// i 是外部看起来的 index 是， 0-based
	i += 1
	if i < 1 || i > m.eleNum {
		return nil, fmt.Errorf("index out of bound")
	}

	//return m.data[m.index[i]], nil
	return m.data[i], nil
}
