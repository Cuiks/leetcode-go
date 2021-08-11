package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	valid := true

	record := map[int][]int{}
	for _, info := range prerequisites {
		record[info[1]] = append(record[info[1]], info[0])
	}

	visited := map[int]int{}

	var dfs func(int)
	dfs = func(i int) {
		visited[i] = 1
		for _, node := range record[i] {
			if visited[node] == 0 {
				dfs(node)
				if !valid {
					return
				}
			} else if visited[node] == 1 {
				valid = false
				return
			}
		}
		visited[i] = 2
	}

	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	return valid
}

func main() {

}
