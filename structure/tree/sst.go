package tree

type Sst struct {
	count int
	root  *SstNode
}

//作为和 bst 对比之用
func NewSst() Sst {
	return Sst{
		count: 0,
		root:  nil,
	}
}

func (b *Sst) Size() int {
	return b.count
}

func (b *Sst) Empty() bool {
	return b.count == 0
}

func (b *Sst) Insert(k string, v int) {
	cur := b.root
	for cur != nil {
		if cur.K == k {
			cur.V = v
			return
		}
		cur = cur.Next
	}

	// 如果没有找到， 插在头部
	newNode := &SstNode{
		K:    k,
		V:    v,
		Next: nil,
	}
	newNode.Next = b.root
	b.root = newNode
	b.count++
}

func (b *Sst) Contain(k string) bool {
	cur := b.root
	for cur != nil {
		if cur.K == k {
			return true
		}
		cur = cur.Next
	}
	return false
}

func (b *Sst) Search(k string) *SstNode {
	cur := b.root
	for cur != nil {
		if cur.K == k {
			return cur
		}
		cur = cur.Next
	}

	return nil
}
