package main

import "fmt"

type ListNode struct {
	Value int
	Next  *ListNode
}

func main() {
	head := &ListNode{
		Value: 1,
		Next:  nil,
	}

	head.Next = &ListNode{
		Value: 2,
		Next:  nil,
	}

	head.Next.Next = &ListNode{
		Value: 3,
		Next:  nil,
	}

	head.Next.Next.Next = &ListNode{
		Value: 4,
		Next:  nil,
	}

	head.Next.Next.Next.Next = &ListNode{
		Value: 5,
		Next:  nil,
	}

	h2 := swapPairs2(head, 2)
	fmt.Println(h2)
}

func swapPairs(head *ListNode) *ListNode {
	dummpy := &ListNode{
		Value: 0,
		Next:  head,
	}
	beforeF := dummpy
	beforeS := dummpy.Next

	//
	for beforeF.Next != nil && beforeS.Next != nil {
		// 保存尾巴
		tail := beforeS.Next.Next
		// 断开尾巴
		beforeS.Next.Next = nil

		f := beforeF.Next
		beforeF.Next = beforeS.Next
		beforeF.Next.Next = f

		// 接上维保
		beforeF.Next.Next.Next = tail

		beforeF = beforeF.Next.Next
		beforeS = beforeF.Next
	}

	return dummpy.Next
}

/**
利用数组的性质, 非 inplace,
*/
func swapPairs2(head *ListNode, n int) *ListNode {
	dummpy := &ListNode{
		Value: 0,
		Next:  head,
	}

	// 锚点
	pos := dummpy
	for {

		h, t := pos, pos
		sortStack := make([]*ListNode, n)
		for i := 0; i < n; i++ {
			if h.Next == nil {
				return dummpy.Next
			}
			h = h.Next
			t = h.Next
			sortStack[i] = &ListNode{Value: h.Value, Next: nil} // 断开连接
		}

		// reverse
		l, r := 0, n-1
		for l < r {
			sortStack[l], sortStack[r] = sortStack[r], sortStack[l]
			l++
			r--
		}

		// 反转
		pos.Next = sortStack[0]
		for i := 0; i < n-1; i++ {
			sortStack[i].Next = sortStack[i+1]
		}
		// 连上尾巴
		sortStack[n-1].Next = t


		// new cur is the last ele of latest stack
		pos = sortStack[n-1]
	}

	return dummpy.Next

}
