package main

func findRepeatNumber(nums []int) int {
	dict := map[int]int{}
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if _, ok := dict[num]; ok {
			return num
		} else {
			dict[num] = 1
		}
	}
	return -1
}

func main() {

}
