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

	p := reverseList(head)
	fmt.Println(p)
}
type ListNode struct {
	Val int
	Next *ListNode
}
func reverseList(head *ListNode) (rev *ListNode) {
	for head != nil {
		head.Next, rev, head = rev, head, head.Next
		//next := head.Next
		//head.Next = rev
		//rev = head
		//head = next
	}

	return
}
