package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if helper(root) == -1 {
		return false
	}
	return true
}

func helper(root *TreeNode) int {
	// 返回深度。-1表示不平衡
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
