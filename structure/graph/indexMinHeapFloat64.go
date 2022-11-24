package graph

import "fmt"

type IndexMinHeapFloat64 struct {
	data   []float64
	index  []int
	eleNum int
}

func NewIndexMinHeapFloat64(cap int) IndexMinHeapFloat64 {
	mh := IndexMinHeapFloat64{}
	mh.Init(cap)
	return mh
}

// 插入的效率没有原地排序效率高
func NewIndexMinHeapFloat64ByInsert(items []float64) IndexMinHeapFloat64 {
	mh := NewIndexMinHeapFloat64(len(items))
	for i := 0; i < mh.eleNum; i++ {
		mh.Insert(i, items[i])
	}

	return mh
}

func NewIndexMinHeapFloat64ByHeapify(items []float64) IndexMinHeapFloat64 {
	mh := IndexMinHeapFloat64{}
	mh.Heapify(items)
	return mh
}

func (m *IndexMinHeapFloat64) Init(cap int) {
	m.data = make([]float64, cap+1) // 需要考虑是否要有顺序插入, 如果是， 需要 length  加上预分配空间
	m.index = make([]int, cap+1)
	//for i := 0; i < len(m.index); i++ {
	//	m.index[i] = i
	//}
	m.eleNum = 0
}

func (m *IndexMinHeapFloat64) Size() int {
	return m.eleNum
}

func (m *IndexMinHeapFloat64) Empty() bool {
	return m.eleNum == 0
}

// 注意 insert 可以可以是追加， 也可以是在某一个索引的位置增加一个元素
// 这里我们为了简便， 理解为在末尾增加一个元素, 而修改我们用 change 来实现
func (m *IndexMinHeapFloat64) Insert(i int, x float64) {
	//if len(m.data) == 0 {
	//	// 插入一个空元素， 使其 1 based
	//	m.data = append(m.data, nil)
	//	m.index = append(m.index, 0)
	//}
	m.data[i+1] = x             // 有意思， 这里稀疏插入，  index 密集整理
	m.index[m.eleNum+1] = i + 1 // 元素怎么变都没有关系， 主要是 index , 从而维护了堆的定义
	m.eleNum += 1
	m.shiftUp(m.eleNum)
}

func (m *IndexMinHeapFloat64) Min() float64 {
	return m.data[m.index[1]]
}

func (m *IndexMinHeapFloat64) shiftUp(nthNum int) {
	for nthNum > 1 && m.data[m.index[nthNum]] < m.data[m.index[nthNum/2]] {
		m.index[nthNum], m.index[nthNum/2] = m.index[nthNum/2], m.index[nthNum]
		nthNum /= 2
	}
}

func (m *IndexMinHeapFloat64) ExtractMin() (float64, error) {
	if m.Empty() {
		return 0, fmt.Errorf("no more element")
	}
	max := m.data[m.index[1]]
	m.index[1], m.index[m.eleNum] = m.index[m.eleNum], m.index[1]
	m.eleNum -= 1
	m.shiftDown(1, m.eleNum)
	return max, nil
}

func (m *IndexMinHeapFloat64) ExtractMinIndex() (int, error) {
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

func (m *IndexMinHeapFloat64) shiftDown(nthNum, n int) {
	for 2*nthNum <= n {
		j := 2 * nthNum
		if j+1 <= n && m.data[m.index[j+1]] < m.data[m.index[j]] {
			j += 1
		}

		if m.data[m.index[j]] < m.data[m.index[nthNum]] {
			m.index[j], m.index[nthNum] = m.index[nthNum], m.index[j]
		}
		nthNum = j
	}
}

func (m *IndexMinHeapFloat64) Heapify(items []float64) {
	sz := len(items)
	m.data = make([]float64, sz+1)
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

func (m *IndexMinHeapFloat64) SortAsc() {
	n := m.eleNum
	for i := m.eleNum; i >= 1; i-- {
		m.index[i], m.index[1] = m.index[1], m.index[i]
		n -= 1
		m.shiftDown(1, n)
	}
}

func (m *IndexMinHeapFloat64) String() string {
	return "not implement yet"
}

func (m *IndexMinHeapFloat64) Display() {
}


/**
i 代表数据索引号
*/
func (m *IndexMinHeapFloat64) Change(i int, item float64) error {
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

func (m *IndexMinHeapFloat64) Contain(k int) bool {
	for i:=0;i <= m.eleNum;i++ {
		if m.index[i] == k+1 {
			return true
		}
	}

	return false
}
