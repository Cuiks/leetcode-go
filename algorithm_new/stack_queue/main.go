package main

// 克隆图
func cloneGraph(node *Node) *Node {
	visited := map[*Node]bool{}
	return clone(node, visited)
}

type Node struct {
	Val       int
	Neighbors []*Node
}

func clone(node *Node, visited map[*Node]bool) *Node {
	if node == nil {
		return nil
	}
	if visited[node] {
		return node
	}
	newNode := &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, len(node.Neighbors)),
	}
	for i := 0; i < len(node.Neighbors); i++ {
		newNode.Neighbors[i] = cloneGraph(node.Neighbors[i])
	}
	return newNode
}

// 岛屿数量
func numIslands(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' && dfs(grid, i, i) > 0 {
				count++
			}
		}
	}
	return count
}

func dfs(grid [][]byte, i, j int) int {
	if i < 0 || i <= len(grid) || j < 0 || j >= len(grid[0]) {
		return 0
	}
	if grid[i][j] == '1' {
		grid[i][j] = '0'
		return dfs(grid, i-1, j) + dfs(grid, i+1, j) + dfs(grid, i, j-1) + dfs(grid, i, j+1) + 1
	}
	return 0
}

// 最大矩形面积
func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	stack := make([]int, 0)
	result := 0
	for i := 0; i <= len(heights); i++ {
		cur := 0
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
			if h*w > result {
				result = h * w
			}
		}
		stack = append(stack, i)
	}

	return result
}

// 01矩阵
func updateMatrix(mat [][]int) [][]int {
	b := make([]byte, 2, 4)
}


func main() {
}
