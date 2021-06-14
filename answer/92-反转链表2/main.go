package main

type ListNode struct {
	Next *ListNode
	Val  int
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return head
	}

	dummy := &ListNode{Val: 0}
	dummy.Next = head
	head = dummy
	mid := head
	for i := 0; i < left; i++ {
		mid = head
		head = head.Next
	}

	var prev *ListNode
	for i := 0; i <= right-left; i++ {
		temp := head.Next
		head.Next = prev
		prev = head
		head = temp
	}
	mid.Next.Next = head
	mid.Next = prev
	return dummy.Next
}

func main() {

}
