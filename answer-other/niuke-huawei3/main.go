package main

import "fmt"

var direction = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {-1, -1}, {1, 1}, {-1, 1}, {1, -1}}

func border(m, n int, record [][]int) int {
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if record[i][j] == 5 {
				flag := true
				for _, item := range direction {
					p, q := i+item[0], j+item[1]
					if p >= 0 && p < m && q >= 0 && q < n && !judgeNode(record, []int{p, q}) {
						flag = false
					}
				}
				if flag {
					res++
				}
				for _, item := range direction {
					p, q := i+item[0], j+item[1]
					if p >= 0 && p < m && q >= 0 && q < n {
						record[p][q] = 2
					}
				}
			}
		}
	}
	return res
}

func judgeNode(record [][]int, node []int) bool {
	flag := true
	for _, item := range direction {
		p, q := node[0]+item[0], node[1]+item[1]
		if p >= 0 && p < len(record) && q >= 0 && q < len(record[0]) {
			if record[p][q] == 2 {
				flag = false
				break
			}
		}
	}
	return flag
}

func main() {
	m, n := 0, 0

	fmt.Scan(&m, &n)
	record := make([][]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if len(record[i]) == 0 {
				record[i] = make([]int, n)
			}
			x := 0
			fmt.Scan(&x)
			record[i][j] = x
		}
	}
	fmt.Println(border(m, n, record))
}
