package main

func majorityElement(nums []int) []int {
	record := map[int]int{}
	for _, v := range nums {
		record[v]++
	}
	limit := len(nums) / 3

	result := make([]int, 0)
	for k, v := range record {
		if v > limit {
			result = append(result, k)
		}
	}
	return result
}

func main() {

}
