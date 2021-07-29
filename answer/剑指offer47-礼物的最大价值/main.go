package main

import "fmt"

func maxValue(grid [][]int) int {
	for i := 1; i < len(grid); i++ {
		grid[i][0] = grid[i-1][0] + grid[i][0]
	}
	for j := 1; j < len(grid[0]); j++ {
		grid[0][j] = grid[0][j-1] + grid[0][j]
	}
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			grid[i][j] = max(grid[i-1][j], grid[i][j-1]) + grid[i][j]
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func main() {
	fmt.Println(maxValue([][]int{{1, 2, 5}, {3, 2, 1}}))
}
