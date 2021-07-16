package main

import "fmt"

func movingCount(m int, n int, k int) int {
	result := 0
	visited := map[string]bool{}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || sum(i, j) > k {
			return
		}
		if v, ok := visited[fmt.Sprintf("%d:%d", i, j)]; ok && v {
			return
		}
		result += 1
		visited[fmt.Sprintf("%d:%d", i, j)] = true
		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
	}
	dfs(0, 0)
	return result
}

func sum(m, n int) int {
	res := 0
	for m > 0 {
		res += m % 10
		m = m / 10
	}
	for n > 0 {
		res += n % 10
		n = n / 10
	}
	return res
}

func main() {
	fmt.Println(movingCount(2, 3, 1))
}
