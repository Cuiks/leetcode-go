package main

// 快速排序
func QuickSort(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, start, end int) {
	if start < end {
		pivot := partition(nums, start, end)
		partition(nums, start, pivot-1)
		partition(nums, pivot+1, end)
	}
}

func partition(nums []int, start, end int) int {
	i := start
	p := nums[end]
	for j := 0; j < end; j++ {
		if nums[j] < p {
			swap(nums, i, j)
			i++
		}
	}
	swap(nums, i, end)
	return i
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

// 归并排序
func MergeSort(nums []int) []int {
	return mergeSort(nums)
}

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2
	left := mergeSort(nums[0:mid])
	right := mergeSort(nums[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	i := 0
	j := 0
	result := make([]int, 0)
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

// 堆排序
func HeapSort(nums []int) []int {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		sink(nums, i, len(nums))
	}

	for i := len(nums) - 1; i >= 1; i-- {
		swap(nums, 0, i)
		sink(nums, 0, i)
	}
	return nums
}

func sink(nums []int, i, length int) {
	for {
		l := i*2 + 1
		r := i*2 + 2
		idx := i
		if l < length && nums[l] > nums[idx] {
			idx = l
		}
		if r < length && nums[r] > nums[idx] {
			idx = r
		}
		if idx == i {
			break
		}
		swap(nums, i, idx)
		i = idx
	}
}

func main() {
}
