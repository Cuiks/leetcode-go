package main

func countBits(n int) []int {
	result := make([]int, n+1)
	for i := 0; i <= n; i++ {
		result[i] = count(i)
	}
	return result
}

func count(num int) int {
	res := 0
	for num != 0 {
		num &= num - 1
		res++
	}
	return res
}
func main() {

}
