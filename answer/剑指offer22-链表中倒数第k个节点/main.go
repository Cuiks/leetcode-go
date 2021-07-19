package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	dummy := head
	for i := 0; i < k; i++ {
		dummy = dummy.Next
	}
	for dummy != nil {
		head = head.Next
		dummy = dummy.Next
	}
	return head
}

func main() {

}
