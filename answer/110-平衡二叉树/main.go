package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	result := helper(root)
	if result == -1 {
		return false
	}
	return true
}

func helper(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := helper(root.Left)
	right := helper(root.Right)

	if left == -1 || right == -1 || left-right > 1 || right-left > 1 {
		return -1
	}
	if left > right {
		return left + 1
	}
	return right + 1
}

func main() {

}
