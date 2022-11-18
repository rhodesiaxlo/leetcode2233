package main

import "fmt"

func main() {
	head := &ListNode{
		Val:  0,
		Next: nil,
	}

	head.Next = &ListNode{
		Val:  1,
		Next: nil,
	}

	head.Next.Next = &ListNode{
		Val:  2,
		Next: nil,
	}
	head.Next.Next.Next = &ListNode{
		Val:  3,
		Next: nil,
	}

	head.Next.Next.Next.Next = &ListNode{
		Val:  4,
		Next: nil,
	}

	p:= reverseKGroup(head, 2)
	fmt.Println(p)
}
/*
1. 反转链表 head  rev  k
2. 多级反转 backtracking
 */
type ListNode struct {
	Val int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	cur,i:= head, 0
	for i < k {
		if cur == nil {
			return head
		}
		cur = cur.Next
		i++
	}

	rev := reverseKGroup(cur, k)

	for i > 0 {
		next:=head.Next
		head.Next = rev
		rev = head
		head = next
		i--
	}
	return rev
}
