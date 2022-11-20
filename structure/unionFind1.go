package structure

/*
UnionFind 在排序和搜索的基础上， 探索了另外一个问题， 就是 i,j 是否相连(connectivity problem)， 在i,j 是否连接的基础上，
进而衍生出了路径问题(path)

和 path 问题相比, connectivity problem 比 path 回答更少的问题， 对比序列n 中第 k 的元素
对比顺序查找(O(n)) 和 bst （O(logn)) 相比， 顺序查找多差了钱 k-1 个元素的问题， 所以比二分慢， 再进一步, heap 则更进一步， 只关心 min,max



应用场景
init
union
find
isConnectioned
*/

// 第一种最糙的实现，就是数组
// 数组的需要代表元素的序数， 值代表 group number, i 和 i 连接， 0 代表该节点并不存在
type UnionFind1 struct {
	//data []int // 正常情况下我们只关心 connection 和 path 问题， 并不关心原始数据, 所以这里并不必要
	id    []int
	count int
}

func NewUnionFind1(n int) UnionFind1 {
	u := UnionFind1{
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

func (u *UnionFind1) Union(i, j int) error {
	//u.id[j] = u.id[i]
	// j 代表的群组合并到 i 代表的群组或者 另外一种
	iGroup, err := u.Find(i)
	if err != nil {
		return err
	}
	jGroup, err := u.Find(j)
	if err != nil {
		return err
	}
	for k := 0; k < u.count; k++ {
		if u.id[k] == jGroup {
			u.id[k] = iGroup
		}
	}
	return nil
}

// 查找序数 i 元素对应的组号 id
func (u *UnionFind1) Find(i int) (int, error) {
	if i < 0 || i >= u.count {
		return 0, nil
	}
	return u.id[i], nil
}

func (u *UnionFind1) IsConnected(i, j int) (bool, error) {
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
