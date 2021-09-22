package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	length := 0
	dummy := head
	for dummy != nil {
		length++
		dummy = dummy.Next
	}

	avgLength := length / k
	mod := length % k
	result := make([]*ListNode, k)
	if head == nil {
		return result
	}
	result[0] = head
	dummy = head

	for i := 1; i < k; i++ {
		l := avgLength
		if mod > 0 {
			l++
			mod--
		}
		for l > 1 {
			dummy = dummy.Next
			l--
		}
		res := dummy.Next
		if res == nil {
			break
		}
		dummy.Next = nil
		result[i] = res
		dummy = res
	}
	return result
}

func main() {

}
