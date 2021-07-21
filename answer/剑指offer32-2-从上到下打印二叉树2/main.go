package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) != 0 {
		l := len(stack)
		temp := make([]int, 0)
		for i := 0; i < l; i++ {
			node := stack[0]
			stack = stack[1:]
			temp = append(temp, node.Val)
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
		result = append(result, temp)
	}
	return result
}

func main() {

}
