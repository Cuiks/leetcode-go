package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, target int) [][]int {
	result := make([][]int, 0)
	var dfs func([]int, int, *TreeNode)
	dfs = func(path []int, limit int, root *TreeNode) {
		if root == nil {
			return
		}
		path = append(path, root.Val)
		limit = limit - root.Val
		if limit == 0 && root.Left == nil && root.Right == nil {
			copyPath := make([]int, len(path))
			copy(copyPath, path)
			result = append(result, copyPath)
			return
		}
		dfs(path, limit, root.Left)
		dfs(path, limit, root.Right)
	}
	path := make([]int, 0)
	dfs(path, target, root)
	return result
}

func main() {

}
