package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	value := preorder[0]
	preorder = preorder[1:]
	root := &TreeNode{
		Val:   value,
		Left:  nil,
		Right: nil,
	}
	index := getIndex(inorder, value)
	root.Left = buildTree(preorder, inorder[0:index])
	root.Right = buildTree(preorder[index:], inorder[index+1:])
	return root
}

func getIndex(list []int, value int) int {
	for i := 0; i < len(list); i++ {
		if list[i] == value {
			return i
		}
	}
	return -1
}

func main() {
	root := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	fmt.Println(root.Val)
}
