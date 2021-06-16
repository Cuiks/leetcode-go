package main

import "fmt"

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

	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for len(q) != 0 {
		point := q[0]
		q = q[1:]
		for _, v := range directions {
			x := point[0] + v[0]
			y := point[1] + v[1]
			if x >= 0 && x < len(mat) && y >= 0 && y < len(mat[0]) && mat[x][y] == -1 {
				mat[x][y] = mat[point[0]][point[1]] + 1
				q = append(q, []int{x, y})
			}
		}
	}
	return mat
}

func main() {
	mat := [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}
	fmt.Println(updateMatrix(mat))
}
