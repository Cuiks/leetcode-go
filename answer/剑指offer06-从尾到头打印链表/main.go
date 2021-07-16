package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	result := make([]int, 0)
	if head == nil {
		return result
	}

	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}

	reverse(result)
	return result
}

func reverse(list []int) {
	for i := 0; i < len(list)/2; i++ {
		list[i], list[len(list)-i-1] = list[len(list)-i-1], list[i]
	}
}

func main() {

}
