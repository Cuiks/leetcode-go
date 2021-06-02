package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序递归遍历
func preorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	preorderTraversal(root.Left)
	preorderTraversal(root.Right)
}

// 前序非递归
func preorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)
	stack := make([]*TreeNode, 0)

	for root != nil || len(stack) != 0 {
		for root != nil {
			result = append(result, root.Val)
			stack = append(stack, root) // record root
			root = root.Left
		}
		// pop root get right
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	return result
}

// 中序非递归
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)
	stack := make([]*TreeNode, 0)

	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		node := stack[len(stack)-1]
		result = append(result, node.Val)
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	return result
}

// 后序非递归
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	var lastVisit *TreeNode

	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
	}
	// root pop after node
	node := stack[len(stack)-1]
	if node.Right == nil || node.Right == lastVisit {
		result = append(result, node.Val)
		stack = stack[:len(stack)-1]
		lastVisit = node
	} else {
		root = node.Right
	}
	return result
}

// DFS 深度搜索  从上到下
func dfs(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	*result = append(*result, root.Val)
	dfs(root.Left, result)
	dfs(root.Right, result)
}

// DFS 深度搜索  从下到上 （分治法）
func divideAndConquer(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return nil
	}
	// 分治
	left := divideAndConquer(root.Left)
	right := divideAndConquer(root.Right)
	// 合并结果
	result = append(result, root.Val)
	result = append(result, left...)
	result = append(result, right...)
	return result
}

// BFS层次遍历
func levelOrder(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			// 遍历queue获取val
			node := queue[0]
			queue = queue[1:]
			result = append(result, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return result
}

func main() {

}
