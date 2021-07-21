package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	res := make([]int, 0)
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return res
	}
	l := len(matrix)
	r := len(matrix[0])
	x, y, dire := 0, 0, 0

	for len(res) < l*r {
		direction := dire % 4
		if direction == 0 {
			for y < r && matrix[x][y] != -1000 {
				res = append(res, matrix[x][y])
				matrix[x][y] = -1000
				y++
			}
			y--
			x++
		} else if direction == 1 {
			for x < l && matrix[x][y] != -1000 {
				res = append(res, matrix[x][y])
				matrix[x][y] = -1000
				x++
			}
			x--
			y--
		} else if direction == 2 {
			for y >= 0 && matrix[x][y] != -1000 {
				res = append(res, matrix[x][y])
				matrix[x][y] = -1000
				y--
			}
			y++
			x--
		} else {
			for x >= 0 && matrix[x][y] != -1000 {
				res = append(res, matrix[x][y])
				matrix[x][y] = -1000
				x--
			}
			x++
			y++
		}
		dire++
	}
	return res
}

func main() {
	fmt.Println(spiralOrder([][]int{{1,2,3},{4,5,6},{7,8,9}}))
}
