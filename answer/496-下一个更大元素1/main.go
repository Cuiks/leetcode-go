package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	result := make([]int, 0)
	if len(nums1) == 0 || len(nums2) == 0 {
		return result
	}
	stack := make([]int, 0)
	max := nums2[0]
	for _, v := range nums2 {
		if v > max {
			max = v
		}
		stack = append(stack, max)
	}
	dict := map[int]int{nums2[len(nums2)-1]: -1}
	for i := len(nums2) - 2; i >= 0; i-- {
		val := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if val > nums2[i] {
			dict[nums2[i]] = val
		} else {
			dict[nums2[i]] = -1
		}
	}

	for _, v := range nums1 {
		result = append(result, dict[v])
	}
	return result
}

func main() {

}
