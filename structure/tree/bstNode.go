package tree

type BstNode struct {
	K    string
	V    int
	L, R *BstNode
}

func NewBstNode(k string, v int) BstNode {
	return BstNode{
		K: k,
		V: v,
		L: nil,
		R: nil,
	}
}
