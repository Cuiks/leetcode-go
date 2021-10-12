package main

import (
	"fmt"
)

/*
二叉树
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序递归遍历
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

	for root != nil || len(stack) != 0 {
		for root != nil {
			result = append(result, root.Val)
			stack = append(stack, root) // record root
			root = root.Left
		}
		// pop root get right
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

	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		node := stack[len(stack)-1]
		result = append(result, node.Val)
		stack = stack[:len(stack)-1]
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
	var lastVisit *TreeNode

	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// root after pop node
		node := stack[len(stack)-1]
		if node.Right == nil || node.Right == lastVisit {
			result = append(result, node.Val)
			stack = stack[:len(stack)-1]
			lastVisit = node
		} else {
			root = node.Right
		}
	}
	return result
}

// DFS 深度搜索  从上到下
func dfs(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	*result = append(*result, root.Val)
	dfs(root.Left, result)
	dfs(root.Right, result)
}

// DFS 深度搜索  从下到上 （分治法）
func divideAndConquer(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	// 分治
	left := divideAndConquer(root.Left)
	right := divideAndConquer(root.Right)
	// 合并结果
	result = append(result, root.Val)
	result = append(result, left...)
	result = append(result, right...)
	return result
}

// BFS层次遍历
func levelOrder(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			// 遍历queue获取val
			node := queue[0]
			queue = queue[1:]
			result = append(result, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return result
}

/* 分治法
应用场景
	快速排序
	归并排序
	二叉树相关问题

分治法模板
	递归返回条件
	分段处理
	合并结果
*/

// 归并排序(分治法)
func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])
	result := merge(left, right)
	return result
}

func merge(left, right []int) []int {
	result := make([]int, 0)
	l := 0
	r := 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
		} else {
			result = append(result, right[r])
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return result
}

// 快速排序(分治法)
func QuickSort(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, start, end int) {
	if start < end {
		pivot := partition(nums, start, end)
		quickSort(nums, 0, pivot-1)
		quickSort(nums, pivot+1, end)
	}
}

func partition(nums []int, start, end int) int {
	p := nums[end]
	i := start
	for j := start; j < end; j++ {
		if nums[j] < p {
			swap(nums, i, j)
			i++
		}
	}
	swap(nums, end, i)
	return i
}

func swap(nums []int, i, j int) {
	t := nums[i]
	nums[i] = nums[j]
	nums[j] = t
}

// 给一个二叉树，找出其最大深度(分治法)
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}

// 判断是否高度平衡的二叉树(分治法)
func isBalanced(root *TreeNode) bool {
	if maxDepth2(root) == -1 {
		return false
	}
	return true
}

func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth2(root.Left)
	right := maxDepth2(root.Right)

	if left == -1 || right == -1 || left-right > 1 || right-left > 1 {
		return -1
	}

	if left > right {
		return left + 1
	}
	return right + 1
}

// 计算非空二叉树最大路径和(分治法)
type ResultType struct {
	SinglePath int
	MaxPath    int
}

func maxPathSum(root *TreeNode) int {
	result := helper(root)
	return result.MaxPath
}

func helper(root *TreeNode) ResultType {
	if root == nil {
		return ResultType{
			SinglePath: 0,
			MaxPath:    -(1 << 31),
		}
	}

	left := helper(root.Left)
	right := helper(root.Right)

	result := ResultType{}
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

// 寻找两个指定节点的最近公共祖先
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

/*
BFS层次应用
*/
// 二叉树自底向上的层次遍历。层次遍历完成后逆序即可

// 二叉树 Z字形遍历
func zigzagLeveOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}

	queue := make([]*TreeNode, 0)
	toggle := false

	for len(queue) > 0 {
		list := make([]int, 0)
		l := len(queue)
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			list = append(list, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if toggle {
			reverse(list)
		}
		result = append(result, list)
		toggle = !toggle
	}
	return result
}

func reverse(list []int) {
	for i := 0; i < len(list)/2; i++ {
		list[i], list[len(list)-1-i] = list[len(list)-1-i], list[i]
	}
}

/*
二叉搜索树的应用
*/

// 判断是否二叉搜索树
// v1. 中序遍历  然后判断是否递增数列
func isValidBST(root *TreeNode) bool {
	result := make([]int, 0)
	inOrder(root, &result)
	for i := 0; i < len(result)-1; i++ {
		if result[i] >= result[i+1] {
			return false
		}
	}
	return true

}

func inOrder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, result)
	*result = append(*result, root.Val)
	inOrder(root.Right, result)
}

// v2. 分治。判断左MAX<根<右MIN
type ResultType2 struct {
	IsValid bool
	Max     *TreeNode
	Min     *TreeNode
}

func isValidBST2(root *TreeNode) bool {
	result := helper2(root)
	return result.IsValid
}

func helper2(root *TreeNode) ResultType2 {
	result := ResultType2{}
	if root == nil {
		result.IsValid = true
		return result
	}

	left := helper2(root.Left)
	right := helper2(root.Right)

	// 判断isValid
	if !left.IsValid || !right.IsValid {
		result.IsValid = false
		return result
	}
	if left.Max != nil && left.Max.Val >= root.Val {
		result.IsValid = false
		return result
	}
	if right.Min != nil && right.Min.Val <= root.Val {
		result.IsValid = false
		return result
	}
	result.IsValid = true

	// update result max/min
	result.Min = root
	if left.Min != nil {
		result.Min = left.Min
	}
	result.Max = root
	if right.Max != nil {
		result.Max = right.Max
	}
	return result
}

// 更新二叉搜索树(插入)
func insertBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		node := &TreeNode{
			Val: val,
		}
		return node
	}

	if val > root.Val {
		root.Right = insertBST(root.Right, val)
	} else {
		root.Left = insertBST(root.Left, val)
	}
	return root
}

func main() {
	root := TreeNode{
		Val: 5,
	}
	left := TreeNode{
		Val: -2,
	}
	right := TreeNode{
		Val: 6,
	}
	root.Left = &left
	root.Right = &right
	fmt.Println(maxPathSum(&root))
}
