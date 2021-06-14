package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	return mergeSort(head)
}

func findMiddle(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func mergeSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := findMiddle(head)
	tail := mid.Next
	mid.Next = nil

	left := mergeSort(head)
	right := mergeSort(tail)
	result := mergeTwoList(left, right)
	return result
}

func mergeTwoList(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	head := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			head.Next = l1
			l1 = l1.Next
		} else {
			head.Next = l2
			l2 = l2.Next
		}
		head = head.Next
	}

	if l1 != nil {
		head.Next = l1
		l1 = l1.Next
		head = head.Next
	}
	if l2 != nil {
		head.Next = l2
		l2 = l2.Next
		head = head.Next
	}
	return dummy.Next
}

func main() {

}
