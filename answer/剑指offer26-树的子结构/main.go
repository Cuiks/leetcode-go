package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	if compare(A, B) && compare(A.Left, B.Left) && compare(A.Right, B.Right) {
		return true
	}
	left := isSubStructure(A.Left, B)
	right := isSubStructure(A.Right, B)
	if left || right {
		return true
	}
	return false
}

func compare(node1 *TreeNode, node2 *TreeNode) bool {
	if node2 == nil {
		return true
	}
	if node1 == nil || (node1.Val != node2.Val) {
		return false
	}
	left := compare(node1.Left, node2.Left)
	right := compare(node1.Right, node2.Right)
	return left && right
}

func main() {
	t1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	t2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}
	fmt.Println(compare(t1, t2))

}
