package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	i, j := headA, headB
	for i != j {
		if i == nil {
			i = headA
		} else {
			i = i.Next
		}
		if j == nil {
			j = headB
		} else {
			j = j.Next
		}
	}
	return i
}

func main() {

}
