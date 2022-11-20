package structure

// union 不应该是固定顺序, 而应该是根据 size

type UnionFind5 struct {
	//data []int // 正常情况下我们只关心 connection 和 path 问题， 并不关心原始数据, 所以这里并不必要
	id    []int
	rank  []int
	count int
}

func NewUnionFind5(n int) UnionFind5 {
	u := UnionFind5{
		id:    nil,
		rank:  nil,
		count: n,
	}

	u.id = make([]int, n)
	u.rank = make([]int, n)
	u.count = n
	for i := 0; i < u.count; i++ {
		u.id[i] = i
		u.rank[i] = 1
	}
	return u
}

func (u *UnionFind5) Union(i, j int) error {
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

	if u.rank[iRoot] < u.rank[jRoot] {
		u.id[iRoot] = jRoot
	} else if u.rank[iRoot] > u.rank[jRoot] {
		u.id[jRoot] = iRoot
	} else {
		u.id[jRoot] = iRoot
		u.rank[iRoot]+=1
	}
	return nil
}

// 查找序数 i 元素对应的组号 id
func (u *UnionFind5) Find(i int) (int, error) {
	if i < 0 || i >= u.count {
		return 0, nil
	}
	cur := i
	for u.id[cur] != cur {
		/*
		 		0               0                        0
			   /			   /					1   2   3   4
			  1				  1
			 /				 /
		    2				2
		   /			   / \
		  3				  3  4
		 /
		4
		 */
		u.id[cur] = u.id[u.id[cur]]
		cur = u.id[cur]
	}
	return cur, nil
}

func (u *UnionFind5) IsConnected(i, j int) (bool, error) {
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
