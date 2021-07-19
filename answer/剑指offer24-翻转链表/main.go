package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var prev *ListNode
	for head != nil{
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

func main() {

}
