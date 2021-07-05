package main

var ans = int(0)

func numWays(n int, relation [][]int, k int) int {
	ways := map[int][]int{}
	for _, v := range relation {
		a, b := v[0], v[1]
		ways[a] = append(ways[a], b)
	}
	dfs(n, k, 0, 0, ways)
	return ans
}

func dfs(n, k, step, v int, ways map[int][]int) {
	if step == k {
		if v == n {
			ans++
		}
		return
	}
	for _, road := range ways[v] {
		dfs(n, k, step+1, road, ways)
	}

}

func main() {

}
