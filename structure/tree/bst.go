package tree

import (
	"fmt"
	"leetcode/leetcode/structure/queue"
)

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
另外 bst 的结构不唯一
   2           3
  / \		 /
1	 3     2
		  /
		 1

basic operation
1. size
2. emtpy
3. insert
4. contain search
5. preOrder, inorder, postorder, levelorder
6. delete // 最难的操作
7. minimum maximum
8. removeMin, removeMax
9. hi

		41
		   \
           58
		 /    \
       50     60
      / \    /  \
     42 53  59  63

remove 58, hibbard deletion  swap (58, 59)
		41
		   \
           59
		 /    \
       50     60
      / \    /  \
     42 53      63

删除有左右节点的节点 d
s = min(d->r)
s->l = d->l
s->r = delMin(d->r)

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

func (b *Bst) PreOrder() {
	b.preOrderNode(b.root)
}

func (b *Bst) InOrder() {
	b.inOrderNode(b.root)
}

func (b *Bst) PostOrder() {
	b.postOrderNode(b.root)
}

func (b *Bst) LevelOrder() {
	qu := queue.NewQueue()
	qu.PushBack(b.root)
	for qu.Len() != 0 {
		node := (qu.Front()).(*BstNode)
		qu.PopFront()
		fmt.Print(node.V, " ")
		if node.L != nil {
			qu.PushBack(node.L)
		}

		if node.R != nil {
			qu.PushBack(node.R)
		}
	}
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
		return b.searchKey(root.R, k)
	}
}

func (b *Bst) Minimum() *BstNode {
	return b.getMinNode(b.root)
}

func (b *Bst) Maximum() *BstNode {
	return b.getMaxNode(b.root)
}

func (b *Bst) RemoveMin() *BstNode {
	min := b.Minimum()
	b.root = b.removeMinNode(b.root)
	return min
}

func (b *Bst) RemoveMax() *BstNode {
	max := b.Maximum()
	b.root = b.removeMaxNode(b.root)
	return max
}

func (b *Bst) Remove(k string) {
	b.root = b.removeNode(b.root, k)
}

func (b *Bst) Agg() int {
	return b.agg(b.root)
}

// -----------------------------------------
func (b *Bst) agg(root *BstNode) int {
	if root == nil {
		return 0
	}

	return root.V + b.agg(root.L) + b.agg(root.R)
}

func (b *Bst) preOrderNode(node *BstNode) {
	if node == nil {
		return
	}

	fmt.Print(node.V, " ")
	b.preOrderNode(node.L)
	b.preOrderNode(node.R)
}

func (b *Bst) inOrderNode(node *BstNode) {
	if node == nil {
		return
	}
	b.inOrderNode(node.L)
	fmt.Print(node.V, " ")
	b.inOrderNode(node.R)
}

func (b *Bst) postOrderNode(node *BstNode) {
	if node == nil {
		return
	}

	b.postOrderNode(node.L)
	b.postOrderNode(node.R)
	fmt.Print(node.V, " ")
}

func (b *Bst) getMinNode(root *BstNode) *BstNode {
	if root == nil {
		return nil
	}

	if root.L == nil {
		return root
	}

	return b.getMinNode(root.L)
}

func (b *Bst) getMaxNode(root *BstNode) *BstNode {
	if root == nil {
		return nil
	}

	if root.R == nil {
		return root
	}

	return b.getMaxNode(root.R)
}

func (b *Bst) removeMinNode(root *BstNode) *BstNode {
	if root.L == nil {
		r := root.R
		b.count--
		return r
	}

	root.L = b.removeMinNode(root.L)
	return root
}

func (b *Bst) removeMaxNode(root *BstNode) *BstNode {
	if root.R == nil {
		l := root.L
		b.count--
		return l
	}

	root.R = b.removeMaxNode(root.R)
	return root
}

func (b *Bst) removeNode(root *BstNode, k string) *BstNode {
	if root == nil {
		return nil
	}

	if k < root.K {
		root.L = b.removeNode(root.L, k)
	} else if k > root.K {
		root.R = b.removeNode(root.R, k)
	} else {
		// 没有左孩子， 直接返回右孩子
		if root.L == nil {
			b.count--
			return root.R
		}

		// 没有右孩子， 直接返回左孩子
		if root.R == nil{
			b.count--
			return root.L
		}

		// hibbard deletion
		tmp:= b.getMinNode(root.R)
		s := &BstNode{
			K: tmp.K,
			V: tmp.V,
			L: nil,
			R: nil,
		}
		b.count++

		s.R = b.removeMinNode(root.R)

		s.L = root.L
		b.count--
		return s
	}
	return root
}
