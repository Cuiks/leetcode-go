package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}
	headDummy := &ListNode{}
	headDummy.Next = head
	head = headDummy

	tailDummy := &ListNode{}
	tail := tailDummy

	for head.Next != nil {
		if head.Next.Val < x {
			head = head.Next
		} else {
			tail.Next = head.Next
			head.Next = head.Next.Next
			tail = tail.Next
		}
	}

	tail.Next = nil
	head.Next = tailDummy.Next
	return headDummy.Next
}

func main() {

}
