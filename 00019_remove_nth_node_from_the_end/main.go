package main

type Node struct {
	Val int
	Next *Node
}
func main() {

}

func removeNthFromTheEnd(head *Node, n int) *Node {
	// 删除倒数第 N 个元素， 所以我们要定位到倒数第n + 1 个元素
	// 如果元素的个数小于n+1个，
	// 最后一个元素是 nil
	// 第一个元素有一个虚拟的 head, Next 是代表单链， 单链是可以为Nil  的
	dummy := &Node{  // dummy 引入后 head 不需要特殊对待拉
		Val:0,
		Next:head,
	}
	front := dummy
	back := dummy

	// front 和 back 相隔 n+1 距离
	for i:= 0; i < n+1; i++  {
		back = back.Next
	}

	// back == nil 代表是最后一个元素， 滑动到最后一个元素位置,front 一起滑动并覆盖更新
	for back != nil {
		front = front.Next
		back = back.Next
	}

	// 如果删除的元素是最后一个, 则不需要关联
	if front.Next != nil {
		front.Next = front.Next.Next
	}
	return dummy.Next
}
