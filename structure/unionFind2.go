package structure

// quick find
// uion 效率低， 从而衍生出另外一种经典实现

// id[i] 代表的不是组号，而是 parent 的序号， 初始化的时候 id[i] 指向的就是 i,  寻找 i 对应的组号的过程一直向上找，找到父节点就是本身的节点
// 这个节点就是组号

type UnionFind2 struct {
	//data []int // 正常情况下我们只关心 connection 和 path 问题， 并不关心原始数据, 所以这里并不必要
	id    []int
	count int
}

func NewUnionFind2(n int) UnionFind2 {
	u := UnionFind2{
		id:    nil,
		count: n,
	}
	u.id = make([]int, n)
	u.count = n
	for i := 0; i < u.count; i++ {
		u.id[i] = i
	}
	return u
}

func (u *UnionFind2) Union(i, j int) error {
	//u.id[j] = u.id[i]
	// j 代表的群组合并到 i 代表的群组或者 另外一种
	iRoot, err := u.Find(i)
	if err != nil {
		return err
	}
	jRoot, err := u.Find(j)
	if err != nil {
		return err
	}

	if iRoot == jRoot {
		return nil
	}
	u.id[iRoot] = jRoot
	return nil
}

// 查找序数 i 元素对应的组号 id
func (u *UnionFind2) Find(i int) (int, error) {
	if i < 0 || i >= u.count {
		return 0, nil
	}
	cur := i
	for u.id[cur] != cur {
		cur = u.id[cur]
	}
	return cur, nil
}

func (u *UnionFind2) IsConnected(i, j int) (bool, error) {
	iId, err := u.Find(i)
	if err != nil {
		return false, err
	}

	jId, err := u.Find(j)
	if err != nil {
		return false, err
	}

	return iId == jId, nil
}
