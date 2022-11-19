package tree

/**
bst 解决了一个什么样的问题
综合场景下经常需要定位一个元素，修改他或者删除他
如果使用普通数组， 查找的复杂度 O(n), 插入和删除都是 O(n）
顺序数组提高了查找效率，利用 BS 可以在 O(logn) 的时间效率下，快速找到元素，但是插入和删除依然是 O(n) 级别的，在大量操作下依然和慢
BST 则在 search  ,insert , delete 上做了平衡， 使得插入和查找和删除的平均复杂度都在 O(logn) 级别


min,max, rank, select, floor, ceil

BST 定义：
每一个节点之多只有2个孩子节点
父亲节点大于左孩子，同时小于右孩子
以左右孩子为根的子树依然满足二分搜索树的性质

值得注意的是 BST 不是 complete tree, 所以用数组表示不是很方便，经典的数据结构是用链表来实现的

basic operation
1. size
2. emtpy
3. insert
4. contain search

*/

type Bst struct {
	count int
	root  *BstNode
}

func NewBst() Bst {
	return Bst{
		count: 0,
		root:  nil,
	}
}

func (b *Bst) Size() int {
	return b.count
}

func (b *Bst) Empty() bool {
	return b.count == 0
}

func (b *Bst) Insert(k string, v int) {
	b.root = b.insertNode(b.root, k, v)
}

func (b *Bst) insertNode(root *BstNode, k string, v int) *BstNode {
	if root == nil {
		b.count++
		return &BstNode{
			K: k,
			V: v,
			L: nil,
			R: nil,
		}
	}

	if root.K == k {
		root.V = v
	} else if root.K > k {
		root.L = b.insertNode(root.L, k, v)
	} else {
		root.R = b.insertNode(root.R, k, v)
	}
	return root
}

func (b *Bst) Contain(k string) bool {
	return b.containKey(b.root, k)
}

func (b *Bst) Agg() int {
	return b.agg(b.root)
}

func (b *Bst) agg(root *BstNode) int {
	if root == nil {
		return 0
	}

	return root.V + b.agg(root.L) + b.agg(root.R)

}

func (b *Bst) Search(k string) *BstNode {
	return b.searchKey(b.root, k)
}

func (b *Bst) insertRecursive() {
	// todo
	// 非递归实现
}

func (b *Bst) containKey(root *BstNode, k string) bool {
	if root == nil {
		return false
	}

	if root.K == k {
		return true
	} else if root.K > k {
		return b.containKey(root.L, k)
	} else {
		return b.containKey(root.R, k)
	}
}

func (b *Bst) searchKey(root *BstNode, k string) *BstNode {
	if root == nil {
		return nil
	}

	if root.K == k {
		return root
	} else if root.K > k {
		return b.searchKey(root.L, k)
	} else {
		return b.searchKey(root.R, k )
	}
}
