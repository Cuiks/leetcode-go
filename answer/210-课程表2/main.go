package main

import "fmt"

func findOrder(numCourses int, prerequisites [][]int) []int {
	record := map[int][]int{}
	for _, info := range prerequisites {
		record[info[0]] = append(record[info[0]], info[1])
	}

	result := make([]int, 0)
	visited := map[int]int{}
	valid := true

	var dfs func(u int)
	dfs = func(u int) {
		visited[u] = 1
		for _, v := range record[u] {
			if visited[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
			} else if visited[v] == 1 {
				valid = false
				return
			}
		}
		visited[u] = 2
		result = append(result, u)
	}

	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	if !valid {
		return []int{}
	}
	return result
}

func main() {
	fmt.Println(findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
}
