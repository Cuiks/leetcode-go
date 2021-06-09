package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type ResultType struct {
	SinglePath int
	MaxPath    int
}

func maxPathSum(root *TreeNode) int {
	result := helper(root)
	return result.MaxPath
}

func helper(root *TreeNode) ResultType {
	result := ResultType{
		SinglePath: 0,
		MaxPath:    -(1 << 32),
	}
	if root == nil {
		return result
	}

	left := helper(root.Left)
	right := helper(root.Right)

	if left.SinglePath > right.SinglePath {
		result.SinglePath = max(left.SinglePath+root.Val, 0)
	} else {
		result.SinglePath = max(right.SinglePath+root.Val, 0)
	}

	maxPath := max(left.MaxPath, right.MaxPath)
	result.MaxPath = max(maxPath, left.SinglePath+right.SinglePath+root.Val)
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

}
