package main

/*
递归思维
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 翻转字符串
func reverseString(s []byte) {
	res := make([]byte, len(s))
	reverse(s, 0, &res)
	for i := 0; i < len(s); i++ {
		s[i] = res[i]
	}

}

func reverse(s []byte, i int, res *[]byte) {
	if i == len(s) {
		return
	}
	reverse(s, i+1, res)
	*res = append(*res, s[i])
}

// 两两交换链表节点
func swapPairs(head *ListNode) *ListNode {
	return helper(head)
}

func helper(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	headNext := head.Next
	head.Next = helper(headNext.Next)
	headNext.Next = head
	return headNext
}

// 生成所有不同的二叉搜索树
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return generate(1, n)
}

func generate(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	res := make([]*TreeNode, 0)
	for i := start; i <= end; i++ {
		lefts := generate(start, i-1)
		rights := generate(i+1, end)
		for j := 0; j < len(lefts); j++ {
			for k := 0; k < len(rights); k++ {
				root := &TreeNode{
					Val: i,
				}
				root.Left = lefts[j]
				root.Right = rights[k]
				res = append(res, root)
			}
		}
	}
	return res
}

// 斐波那契数
func fib(N int) int {
	return dfs(N)
}

var cache = make(map[int]int)

func dfs(n int) int {
	if n <=1 {
		return 1
	}
	if _, ok := cache[n]; ok {
		return cache[n]
	}
	res := dfs(n-1) + dfs(n-2)
	cache[n] = res
	return res
}

func main() {
}
