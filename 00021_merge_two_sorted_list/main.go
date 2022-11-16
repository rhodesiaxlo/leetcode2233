package main

import "fmt"

/**
1. 链表 nep 问题， 需要提前终止循环
 */
func main() {
	l1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	l1.Next = &ListNode{
		Val:  2,
		Next: nil,
	}

	l1.Next.Next = &ListNode{
		Val:  4,
		Next: nil,
	}

	l2 := &ListNode{
		Val:  1,
		Next: nil,
	}
	l2.Next = &ListNode{
		Val:  3,
		Next: nil,
	}
	l2.Next.Next = &ListNode{
		Val:  4,
		Next: nil,
	}

	p := mergeTwoLists2(l1, l2)
	fmt.Println(p)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummpy := &ListNode{
		Val:  0,
		Next: nil,
	}

	cur := dummpy
	i, j := list1, list2
	for {
		if i == nil && j == nil {
			return dummpy.Next
		} else if i == nil {
			cur.Next = j
			j = j.Next
		} else if j == nil {
			cur.Next = i
			i = i.Next
		} else if i.Val < j.Val {
			cur.Next = i
			i = i.Next
		} else {
			cur.Next = j
			j = j.Next
		}

		cur = cur.Next
	}
}

/**
循环的终止条件应为是 l， r 有一个位空就不再往下了
*/
func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	l, r := list1, list2
	for l != nil && r != nil {
		if l.Val < r.Val {
			cur.Next = l
			l = l.Next
		} else {
			cur.Next = r
			r = r.Next
		}
		cur = cur.Next
	}

	if l == nil {
		cur.Next = r
	} else {
		cur.Next = l
	}
	return dummy.Next
}
