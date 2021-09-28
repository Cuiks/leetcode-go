package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	var dfs func(node *TreeNode, sum int)
	res := 0
	dfs = func(node *TreeNode, sum int) {
		sum += node.Val
		if sum == targetSum {
			res++
		}
		if node.Left != nil {
			dfs(node.Left, sum)
		}
		if node.Right != nil {
			dfs(node.Right, sum)
		}
		sum -= node.Val
		return
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			dfs(node, 0)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return res
}

func main() {
	root := &TreeNode{Val: 10}
	node1 := &TreeNode{Val: 5}
	node2 := &TreeNode{Val: -3}
	node3 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 2}
	node5 := &TreeNode{Val: 11}
	node6 := &TreeNode{Val: 3}
	node7 := &TreeNode{Val: -2}
	node8 := &TreeNode{Val: 1}
	root.Left = node1
	root.Right = node2
	node1.Left = node3
	node1.Right = node4
	node2.Right = node5
	node3.Left = node6
	node3.Right = node7
	node4.Right = node8
	fmt.Println(pathSum(root, 8))
}
