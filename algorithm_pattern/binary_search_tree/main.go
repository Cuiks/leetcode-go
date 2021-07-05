package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 验证二叉搜索树
type ResultType struct {
	min   int
	max   int
	valid bool
}

func isValidBST(root *TreeNode) bool {
	return dfs(root).valid
}

func dfs(root *TreeNode) ResultType {
	result := ResultType{
		min:   math.MaxInt64,
		max:   math.MinInt64,
		valid: true,
	}
	if root == nil {
		return result
	}

	left := dfs(root.Left)
	right := dfs(root.Right)
	if left.valid && right.valid && left.max < root.Val && root.Val < right.min {
		result.valid = true
	}
	result.min = Min(Min(left.min, right.min), root.Val)
	result.max = Max(Max(left.max, right.max), root.Val)
	return result
}

func Min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 二叉搜索树插入值
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val:   val,
			Left:  nil,
			Right: nil,
		}
	}
	if val < root.Val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}

// 二叉搜索树删除节点
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			cur := root.Right
			for cur.Left != nil {
				cur = cur.Left
			}
			cur.Left = root.Left
			return root.Right
		}
	}
	return root
}

// 判断二叉平衡树
func isBalanced(root *TreeNode) bool {
	return dfs2(root).valid
}

type BalancedType struct {
	height int
	valid  bool
}

func dfs2(root *TreeNode) *BalancedType {
	if root == nil {
		return &BalancedType{
			height: 0,
			valid:  true,
		}
	}
	result := &BalancedType{}
	left := dfs2(root.Left)
	right := dfs2(root.Right)
	if left.valid && right.valid && (left.height-right.height <= 1 || right.height-left.height <= 1) {
		result.valid = true
	}
	result.height = Max(left.height, right.height) + 1
	return result
}

func main() {

}
