package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargest(root *TreeNode, k int) int {
	res := 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil{
			return
		}
		if node.Right != nil {
			dfs(node.Right)
		}
		if k == 0 {
			return
		}
		k -= 1
		if k == 0 {
			res = node.Val
			return
		}
		dfs(node.Left)
	}

	dfs(root)
	return res
}

func main() {

}
