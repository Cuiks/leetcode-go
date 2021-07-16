package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	head = dummy
	for head != nil {
		for head.Next != nil && head.Next.Val == val {
			head.Next = head.Next.Next
		}
		head = head.Next
	}
	return dummy.Next
}

func main() {

}
