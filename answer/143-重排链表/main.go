package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	mid := findMiddle(head)

	tail := reverse(mid.Next)
	mid.Next = nil

	dummy := &ListNode{}
	cur := dummy

	toggle := true
	for head != nil && tail != nil {
		if toggle {
			cur.Next = head
			head = head.Next
		} else {
			cur.Next = tail
			tail = tail.Next
		}
		cur = cur.Next
		toggle = !toggle
	}

	for head != nil{
		cur.Next = head
		head = head.Next
		cur = cur.Next
	}

	for tail != nil{
		cur.Next = tail
		tail = tail.Next
		cur = cur.Next
	}

	head = dummy.Next
}

func findMiddle(head *ListNode) *ListNode {
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverse(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var prev *ListNode
	for head != nil {
		temp := head.Next
		head.Next = prev
		prev = head
		head = temp
	}
	return prev
}

func main() {

}
