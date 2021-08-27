package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 判断列表有环
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast := head.Next
	slow := head
	for fast != nil && fast.Next != nil {
		if fast == slow {
			return true
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return false
}



func main() {

}
