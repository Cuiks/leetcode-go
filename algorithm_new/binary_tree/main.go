package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序递归
func preorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	preorderTraversal(root.Left)
	preorderTraversal(root.Right)
}

// 前序非递归
func preorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			result = append(result, root.Val)
			stack = append(stack, root)
			root = root.Left
		}

		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	return result
}

// 中序非递归
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)
		root = node.Right
	}
	return result
}

// 后序非递归
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	var lastVisited *TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		if node.Right == nil || node.Right == lastVisited {
			stack = stack[:len(stack)-1]
			result = append(result, node.Val)
			lastVisited = node
		} else {
			root = node.Right
		}
	}
	return result
}

// 分治法前序遍历
func preorderTraversal3(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	return helperPreorder(root)
}

func helperPreorder(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	left := helperPreorder(root.Left)
	right := helperPreorder(root.Right)
	result = append(result, root.Val)
	result = append(result, left...)
	result = append(result, right...)
	return result
}

// MergeSort 分治法-归并排序
func MergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	left := MergeSort(nums[:mid])
	right := MergeSort(nums[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	i, j := 0, 0
	result := make([]int, 0)
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

// QuickSort 快速排序
func QuickSort(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, start, end int) {
	if start < end {
		point := partition(nums, start, end)
		quickSort(nums, start, point-1)
		quickSort(nums, point+1, end)
	}
}

func partition(nums []int, start, end int) int {
	p := nums[end]
	j := start
	for i := start; i < end; i++ {
		if nums[j] < p {
			swap(nums, i, j)
			j++
		}
	}
	swap(nums, j, end)
	return j
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

// 平衡二叉树
func isBalanced(root *TreeNode) bool {
	if maxDepth(root) == -1 {
		return false
	}
	return true
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left == -1 || right == -1 || left-right > 1 || right-left > 1 {
		return -1
	}
	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}

// 非空二叉树返回路径最大和
func maxPathSum(root *TreeNode) int {
	return helper(root).maxPath
}

type ResultType struct {
	maxPath    int
	singlePath int
}

func helper(root *TreeNode) ResultType {
	if root == nil {
		return ResultType{
			maxPath:    0,
			singlePath: math.MinInt64,
		}
	}

	left := helper(root.Left)
	right := helper(root.Right)
	result := ResultType{}
	if left.singlePath > right.singlePath {
		result.singlePath = max(left.singlePath+root.Val, 0)
	} else {
		result.singlePath = max(right.singlePath+root.Val, 0)
	}
	maxPath := max(left.maxPath, right.maxPath)
	result.maxPath = max(maxPath, left.singlePath+right.singlePath+root.Val)
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	return nil
}

// 判断是否二叉搜索树
func isValidBST(root *TreeNode) bool {
	return helperValidBST(root).isValid
}

type ResultValidBST struct {
	isValid bool
	min     *TreeNode
	max     *TreeNode
}

func helperValidBST(root *TreeNode) ResultValidBST {
	if root == nil {
		return ResultValidBST{isValid: true}
	}

	result := ResultValidBST{}
	left := helperValidBST(root.Left)
	right := helperValidBST(root.Right)
	if !left.isValid || !right.isValid {
		result.isValid = false
		return result
	}
	if left.max != nil && left.max.Val >= root.Val {
		result.isValid = false
		return result
	}
	if right.min != nil && left.min.Val <= root.Val {
		result.isValid = false
		return result
	}

	result.isValid = true
	if left.min != nil {
		result.min = left.min
	} else {
		result.min = root
	}
	if right.max != nil {
		result.max = right.max
	} else {
		result.max = root
	}
	return result
}

// 二叉搜索树插入新节点
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil{
		return &TreeNode{
			Val:   val,
		}
	}
	if val < root.Val{
		root.Left = insertIntoBST(root.Left, val)
	}else {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}

func main() {

}
