package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
栈
*/

// 设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈
type MinStack struct {
	min   []int
	stack []int
}

func Constructor() MinStack {
	return MinStack{
		min:   make([]int, 0),
		stack: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	min := this.GetMin()
	if min < val {
		this.min = append(this.min, min)
	} else {
		this.min = append(this.min, val)
	}
	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	this.min = this.min[:len(this.min)-1]
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return 0
	}
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.min) == 0 {
		return 1 << 31
	}
	return this.min[len(this.min)-1]
}

// 逆波兰表达式求值
func evalRPN(tokens []string) int {
	if len(tokens) == 0 {
		return 0
	}

	stack := make([]int, 0)
	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "+", "-", "*", "/":
			if len(stack) < 2 {
				return -1
			}

			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-1]

			var result int
			switch tokens[i] {
			case "+":
				return a + b
			case "-":
				return a - b
			case "*":
				return a * b
			case "/":
				return a / b
			}
			stack = append(stack, result)
		default:
			val, _ := strconv.Atoi(tokens[i])
			stack = append(stack, val)
		}
	}
	return stack[0]
}

// 字符串解码
func decodeString(s string) string {
	if len(s) == 0 {
		return ""
	}
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ']':
			temp := make([]byte, 0)
			for len(stack) != 0 && stack[len(stack)-1] != '[' {
				v := stack[len(stack)-1]
				temp = append(stack, v)
				stack = stack[:len(stack)-1]
			}
			// pop '['
			stack = stack[:len(stack)-1]
			// pop num
			idx := 1
			for len(stack) >= idx && stack[len(stack)-idx] > '0' && stack[len(stack)-idx] < '9' {
				idx++
			}
			num := stack[len(stack)-idx+1:]
			stack = stack[:len(stack)-idx+1]
			count, _ := strconv.Atoi(string(num))
			for j := 0; j < count; j++ {
				for k := len(temp) - 1; k >= 0; k-- {
					stack = append(stack, temp[k])
				}
			}
		default:
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}

/*
利用栈进行 DFS 递归搜索模板
boolean DFS(int root, int target) {
    Set<Node> visited;
    Stack<Node> s;
    add root to s;
    while (s is not empty) {
        Node cur = the top element in s;
        return true if cur is target;
        for (Node next : the neighbors of cur) {
            if (next is not in visited) {
                add next to s;
                add next to visited;
            }
        }
        remove cur from s;
    }
    return false;
}
*/

// 给定一个二叉树，返回它的中序遍历。
func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	stack := make([]*TreeNode, 0)
	for len(stack) != 0 || root != nil {
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

// 克隆无向连通图
type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	visited := make(map[*Node]*Node)
	return clone(node, visited)
}

func clone(node *Node, visited map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}

	if v, ok := visited[node]; ok {
		return v
	}

	newNode := &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, len(node.Neighbors)),
	}
	visited[node] = newNode

	for i := 0; i < len(node.Neighbors); i++ {
		newNode.Neighbors[i] = clone(node.Neighbors[i], visited)
	}
	return newNode
}

// 岛屿数量
func numIslands(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' && dfs(grid, i, j) >= 1 {
				count++
			}
		}
	}
	return count
}

func dfs(grid [][]byte, i, j int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return 0
	}
	if grid[i][j] == '1' {
		grid[i][j] = 0
		return dfs(grid, i-1, j) + dfs(grid, i, j-1) + dfs(grid, i+1, j) + dfs(grid, i, j+1) + 1
	}
	return 0
}

// 柱状图中最大的矩形
func largestRectangleArea(heights []int) int {
	max := 0
	if len(heights) == 0 {
		return max
	}
	stack := make([]int, 0)
	for i := 0; i <= len(heights); i++ {
		var cur int
		if i == len(heights) {
			cur = 0
		} else {
			cur = heights[i]
		}

		for len(stack) != 0 && cur <= heights[stack[len(stack)-1]] {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			h := heights[pop]
			w := i
			if len(stack) != 0 {
				peek := stack[len(stack)-1]
				w = i - peek - 1
			}
			max = Max(max, h*w)
		}
		stack = append(stack, i)
	}
	return max

}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
Queue 队列
*/
// 使用栈实现队列
type MyQueue struct {
	stack []int
	back  []int
}

/** Initialize your data structure here. */
func ConstructorQueue() MyQueue {
	return MyQueue{
		stack: make([]int, 0),
		back:  make([]int, 0),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	for len(this.back) != 0 {
		val := this.back[len(this.back)-1]
		this.stack = append(this.stack, val)
		this.back = this.back[:len(this.back)-1]
	}
	this.stack = append(this.stack, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	for len(this.stack) != 0 {
		val := this.stack[len(this.stack)-1]
		this.back = append(this.back, val)
		this.stack = this.stack[:len(this.stack)-1]
	}
	if len(this.back) == 0 {
		return 0
	}
	pop := this.back[len(this.back)-1]
	this.back = this.back[:len(this.back)-1]
	return pop
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	for len(this.stack) != 0 {
		val := this.stack[len(this.stack)-1]
		this.back = append(this.back, val)
		this.stack = this.stack[:len(this.stack)-1]
	}
	if len(this.back) == 0 {
		return 0
	}
	return this.back[len(this.back)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.stack) == 0 && len(this.back) == 0
}

// 01矩阵
func updateMatrix(mat [][]int) [][]int {
	q := make([][]int, 0)
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 0 {
				q = append(q, []int{i, j})
			} else {
				mat[i][j] = -1
			}
		}
	}

	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for len(q) != 0 {
		point := q[0]
		q = q[1:]

		for _, v := range directions {
			x := point[0] + v[0]
			y := point[1] + v[1]
			if x >= 0 && x < len(mat) && y >= 0 && y < len(mat) && mat[x][y] == -1 {
				mat[x][y] = mat[point[0]][point[1]] + 1
				q = append(q, []int{x, y})
			}
		}
	}
	return mat
}

func main() {
	heights := []int{2, 1, 5, 6, 2, 3}
	fmt.Println(largestRectangleArea(heights))
}
