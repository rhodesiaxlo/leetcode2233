package structure

// union 不应该是固定顺序, 而应该是根据 size

type UnionFind3 struct {
	//data []int // 正常情况下我们只关心 connection 和 path 问题， 并不关心原始数据, 所以这里并不必要
	id    []int
	sz    []int
	count int
}

func NewUnionFind3(n int) UnionFind3 {
	u := UnionFind3{
		id:    nil,
		sz:    nil,
		count: n,
	}
	u.id = make([]int, n)
	u.sz = make([]int, n)
	u.count = n
	for i := 0; i < u.count; i++ {
		u.id[i] = i
		u.sz[i] = 1
	}
	return u
}

func (u *UnionFind3) Union(i, j int) error {
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

	if u.sz[iRoot] < u.sz[jRoot] {
		u.id[iRoot] = jRoot
		u.sz[jRoot] += u.sz[iRoot]
	} else {
		u.id[jRoot] = iRoot
		u.sz[iRoot] += u.sz[jRoot]
	}
	return nil
}

// 查找序数 i 元素对应的组号 id
func (u *UnionFind3) Find(i int) (int, error) {
	if i < 0 || i >= u.count {
		return 0, nil
	}
	cur := i
	for u.id[cur] != cur {
		cur = u.id[cur]
	}
	return cur, nil
}

func (u *UnionFind3) IsConnected(i, j int) (bool, error) {
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
