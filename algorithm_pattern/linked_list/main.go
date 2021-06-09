package main

type ListNode struct {
	Next *ListNode
	Val  int
}

// 给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

func deleteDuplicates(head *ListNode) *ListNode {
	current := head
	for current != nil {
		for current.Next != nil && current.Val == current.Next.Val {
			current.Next = current.Next.Next
		}
		current = current.Next
	}
	return head
}


// 排序链表，删除所有含有重复数字的节点，只保留没有重复出现的数字。
