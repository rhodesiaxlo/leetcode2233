package heap

import "fmt"

/**
增加一个 rev 数组, 用于在修改元素的数组的时候，能够快速定位到 data[i] 在 index 中的位置
这样，如果修改的次数非常大的时候， 时间复杂度就可以从 m*n*log(n) 降到 nlog(n)
rev 数组和 index 数组具备这样的性质
index[i] = j
rev[j] = i
index[rev[i]] = i
rev[index[i]] = i

需要修改的有3个地方
1. 初始化， 初始化的事后 rev 全部为0 根据定义 index 代表 在第 i 个位置存的 data 的索引
   如果没有排序，
2. shiftUp
3. shiftDown
4. extractMax
	m.rev[m.index[1]] = 1
	m.rev[m.index[m.eleNum]] = 0 // 如何理解这段
5. change
*/
type IndexMaxHeapWithRev struct {
	data   []Item
	index  []int
	rev    []int
	eleNum int

	// 为了减少内存扩充，需要增加一个
	// cap int
}

func NewIndexMaxHeapWithRevInit(cap int) IndexMaxHeapWithRev {
	// 构造一个容量为 n 的 IndexMaxHeap
	mh := IndexMaxHeapWithRev{}
	mh.Init(cap)
	return mh
}

// 插入的效率没有原地排序效率高
func NewIndexMaxHeapWithRevInsert(items []Item) IndexMaxHeapWithRev {
	mh := IndexMaxHeapWithRev{}
	sz := len(items)
	for i := 0; i < sz; i++ {
		// 保留备份， 能否用指针呢？ 不能
		mh.Insert(items[i])
	}

	return mh
}

func NewIndexMaxHeapWithRevHeapify(items []Item) IndexMaxHeapWithRev {
	mh := IndexMaxHeapWithRev{}
	mh.Heapify(items)
	return mh
}

func (m *IndexMaxHeapWithRev) Init(cap int) {
	m.data = make([]Item, cap+1)
	m.index = make([]int, cap+1)
	m.rev = make([]int, cap+1)
	m.eleNum = 0
}

func (m *IndexMaxHeapWithRev) Size() int {
	return m.eleNum
}

func (m *IndexMaxHeapWithRev) Empty() bool {
	return m.eleNum == 0
}

// 注意 insert 可以可以是追加， 也可以是在某一个索引的位置增加一个元素
// 这里我们为了简便， 理解为在末尾增加一个元素, 而修改我们用 change 来实现
func (m *IndexMaxHeapWithRev) Insert(x Item) {
	if len(m.data) == 0 {
		// 插入一个空元素， 使其 1 based
		m.data = append(m.data, Item{})
		m.index = append(m.index, 0)
		m.rev = append(m.rev, 0)
	}
	m.data = append(m.data, x)
	m.eleNum += 1
	// index[eleNum] = eleNum
	m.index = append(m.index, m.eleNum)
	m.rev = append(m.rev, m.eleNum)
	m.shiftUp(m.eleNum)
}

func (m *IndexMaxHeapWithRev) Max() Item {
	return m.data[m.index[1]]
}

func (m *IndexMaxHeapWithRev) shiftUp(nthNum int) {
	for nthNum > 1 && m.data[m.index[nthNum]].Gt(m.data[m.index[nthNum/2]]) {
		m.index[nthNum], m.index[nthNum/2] = m.index[nthNum/2], m.index[nthNum]

		// 维护定义
		m.rev[m.index[nthNum]] = nthNum
		m.rev[m.index[nthNum/2]] = nthNum / 2

		nthNum /= 2
	}
}

func (m *IndexMaxHeapWithRev) ExtractMax() (Item, error) {
	if m.Empty() {
		return Item{}, fmt.Errorf("no more element")
	}
	max := m.data[m.index[1]]
	m.index[1], m.index[m.eleNum] = m.index[m.eleNum], m.index[1]

	m.rev[m.index[1]] = 1
	m.rev[m.index[m.eleNum]] = 0 // 如何理解这段

	m.eleNum -= 1
	m.shiftDown(1, m.eleNum)
	return max, nil
}

func (m *IndexMaxHeapWithRev) shiftDown(nthNum, n int) {
	for 2*nthNum <= n {
		j := 2 * nthNum
		if j+1 <= n && m.data[m.index[j+1]].Gt(m.data[m.index[j]]) {
			j += 1
		}

		if m.data[m.index[j]].Gt(m.data[m.index[nthNum]]) {
			m.index[j], m.index[nthNum] = m.index[nthNum], m.index[j]

			// 维护定义
			m.rev[m.index[j]] = j
			m.rev[m.index[nthNum]] = nthNum

		}
		nthNum = j
	}
}

func (m *IndexMaxHeapWithRev) Heapify(items []Item) {
	sz := len(items)
	m.data = make([]Item, sz+1)
	m.index = make([]int, sz+1)
	m.rev = make([]int, sz+1)
	copy(m.data[1:], items)
	for i := 1; i <= sz; i++ {
		m.index[i] = i
		m.rev[i] = i
	}
	m.eleNum = sz

	for lastParent := m.eleNum / 2; lastParent >= 1; lastParent-- {
		m.shiftDown(lastParent, m.eleNum)
	}
}

func (m *IndexMaxHeapWithRev) SortAsc() {
	n := m.eleNum
	for i := m.eleNum; i >= 1; i-- {
		m.index[i], m.index[1] = m.index[1], m.index[i]
		n -= 1
		m.shiftDown(1, n)
	}
}

func (m *IndexMaxHeapWithRev) String() string {
	return "not implement yet"
}

func (m *IndexMaxHeapWithRev) Display() {
}

func (m *IndexMaxHeapWithRev) GetArr() []Item {
	ans := make([]Item, m.eleNum)
	for i := 0; i < m.eleNum; i++ {
		ans[i] = m.data[m.index[i+1]]
	}
	return ans
}

func (m *IndexMaxHeapWithRev) Change(i int, item Item) error {
	// i 是外部看起来的 index 是， 0-based
	i += 1
	if i < 1 || i > m.eleNum {
		return fmt.Errorf("index out of bound")
	}
	m.data[i] = item
	// 注意这里的找到元素
	//for k := 1; k < m.eleNum; k++ {
	//	if m.index[k] == i {
	//		m.shiftUp(k)
	//		m.shiftDown(k, m.eleNum)
	//		break
	//	}
	//}
	m.shiftUp(m.rev[i])
	m.shiftDown(m.rev[i], m.eleNum)
	return nil
}

func (m *IndexMaxHeapWithRev) GetItem(i int) (Item, error) {
	// i 是外部看起来的 index 是， 0-based
	i += 1
	if i < 1 || i > m.eleNum {
		return Item{}, fmt.Errorf("index out of bound")
	}

	//return m.data[m.index[i]], nil
	return m.data[i], nil
}
