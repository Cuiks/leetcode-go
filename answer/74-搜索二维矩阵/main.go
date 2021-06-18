package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	row := len(matrix)
	col := len(matrix[0])
	start := 0
	end := row*col - 1
	for start+1 < end {
		mid := start + (end-start)/2
		if target == matrix[mid/col][mid%col] {
			return true
		} else if target < matrix[mid/col][mid%col] {
			end = mid
		} else {
			start = mid
		}
	}
	if target == matrix[start/col][start%col] || target == matrix[end/col][end%col] {
		return true
	}
	return false
}
func main() {
	matrix := [][]int{{1, 1}}
	target := 13
	fmt.Println(searchMatrix(matrix, target))
}
