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
func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	dummy := &ListNode{}
	dummy.Next = head
	head = dummy

	for head.Next != nil && head.Next.Next != nil {
		if head.Next.Val == head.Next.Next.Val {
			rmVal := head.Next.Val
			for head.Next != nil && head.Next.Val == rmVal {
				head.Next = head.Next.Next
			}
		} else {
			head = head.Next
		}
	}
	return dummy.Next
}

// 反转一个单链表
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		temp := head.Next
		head.Next = prev
		prev = head
		head = temp
	}
	return prev
}

// 反转从位置 m 到 n 的链表
func reverseBetween(head *ListNode, m, n int) *ListNode {
	if head == nil {
		return head
	}

	dummy := head
	dummy.Next = head
	head = dummy

	mid := head
	for i := 0; i < m; i++ {
		mid = head
		head = head.Next
	}

	var prev *ListNode
	for j := 0; head != nil && j <= m-n; j++ {
		temp := head.Next
		head.Next = prev
		prev = head
		head = temp
	}
	mid.Next.Next = head
	mid.Next = prev

	return dummy.Next
}

// 合并两个升序链表
func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
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

	for l1 != nil {
		head.Next = l1
		l1 = l1.Next
		head = head.Next
	}
	for l2 != nil {
		head.Next = l2
		l2 = l2.Next
		head = head.Next
	}
	return dummy.Next
}

// 对链表根据值x进行分割
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}
	headDummy := &ListNode{Val: 0}
	tailDummy := &ListNode{Val: 0}

	headDummy.Next = head
	head = headDummy

	tail := tailDummy
	for head.Next != nil {
		if head.Next.Val < x {
			head = head.Next
		} else {
			t := head.Next
			head.Next = head.Next.Next
			tail.Next = t
			tail = tail.Next
		}
	}
	tail.Next = nil
	head.Next = tailDummy.Next
	return headDummy.Next
}

// O(nlogn)时间复杂度 O(1)空间复杂度下 排序
func sortList(head *ListNode) *ListNode {
	return mergeSort(head)
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

func findMiddle(head *ListNode) *ListNode {
	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
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

	for l1 != nil {
		head.Next = l1
		l1 = l1.Next
		head = head.Next
	}
	for l2 != nil {
		head.Next = l2
		l2 = l2.Next
		head = head.Next
	}
	return dummy.Next
}

// 找到中点，翻转后面部分
func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	mid := findMiddle(head)
	tail := reverseList(mid.Next)
	mid.Next = nil

	head = mergeTwoList2(head, tail)
}

func mergeTwoList2(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	head := dummy

	toggle := true
	for l1 != nil && l2 != nil {
		if toggle {
			head.Next = l1
			l1 = l1.Next
		} else {
			head.Next = l2
			l2 = l2.Next
		}
		head = head.Next
		toggle = !toggle
	}

	for l1 != nil {
		head.Next = l1
		head = head.Next
		l1 = l1.Next
	}
	for l2 != nil {
		head.Next = l2
		head = head.Next
		l2 = l2.Next
	}
	return dummy.Next
}

// 判断链表是否有环
func detectCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}

// 返回链表开始入环的第一个节点，无环返回nil
func detectCycle2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		if fast == slow {
			fast = head
			slow = slow.Next
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return fast
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return nil
}

// 判断是否回文链表
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	mid := findMiddle(head)
	tail := reverseList(mid.Next)
	mid.Next = nil

	for head != nil && tail != nil {
		if head != tail {
			return false
		}
		head = head.Next
		tail = tail.Next
	}
	return true
}

// 复制带随机指针的链表
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}

	cur := head
	for cur != nil {
		cloned := &Node{Val: cur.Val, Next: cur.Next}
		temp := cur.Next
		cur.Next = cloned
		cur = temp
	}

	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next
	}

	cur = head
	clonedHead := cur.Next
	for cur != nil && cur.Next != nil {
		temp := cur.Next
		cur.Next = cur.Next.Next
		cur = temp
	}
	return clonedHead
}
