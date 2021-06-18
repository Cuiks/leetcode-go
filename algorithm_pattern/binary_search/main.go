package main

// 二分查找
func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start+1 < end {
		mid := start + (end-start)/2
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			end = mid
		} else {
			start = mid
		}
	}
	if nums[start] == target {
		return start
	}
	if nums[end] == target {
		return end
	}
	return -1
}

// 查找排序数组中target的起始索引和最后索引
func searchRange(A []int, target int) []int {
	if len(A) == 0 {
		return []int{-1, -1}
	}

	result := make([]int, 2)
	start := 0
	end := len(A) - 1
	for start+1 < end {
		mid := start + (end-start)/2
		if target <= A[mid] {
			end = mid
		} else {
			start = mid
		}
	}

	if A[start] == target {
		result[0] = start
	} else if A[end] == target {
		result[0] = end
	} else {
		result[0] = -1
		result[1] = -1
		return result
	}

	start = 0
	end = len(A) - 1
	for start+1 < end {
		mid := start + (end-start)/2
		if target < A[mid] {
			end = mid
		} else {
			start = mid
		}
	}

	if A[start] == target {
		result[1] = start
	} else if A[end] == target {
		result[1] = end
	} else {
		result[0] = -1
		result[1] = -1
		return result
	}
	return result
}

// 查找插入位置
func searchInsert(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start+1 < end {
		mid := start + (end-start)/2
		if nums[mid] >= target {
			end = mid
		} else {
			start = mid
		}
	}

	if target <= nums[start] {
		return start
	} else if target <= nums[end] {
		return end
	} else if target > nums[end] {
		return end + 1
	}
	return 0
}

// 判断 m x n 矩阵中，是否存在一个目标值
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	row := len(matrix)
	col := len(matrix[0])
	start := 0
	end := row*col - 1
	for start+1 < end {
		mid := start + (end-start)/2
		val := matrix[mid/col][mid%col]
		if target == val {
			return true
		} else if target < val {
			end = mid
		} else {
			start = mid
		}
	}

	if target == matrix[start/col][start%col] || target == matrix[end/col][end%col] {
		return true
	}
	return false
}

// first-bad-version
func isBadVersion(n int) bool {
	return true
}
func firstBadVersion(n int) int {
	start := 0
	end := n
	for start+1 < end {
		mid := start + (end-start)/2
		if isBadVersion(mid) {
			end = mid
		} else {
			start = mid
		}
	}
	if isBadVersion(start) {
		return start
	}
	return end
}

// 旋转数组找最小值
func findMin(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	start := 0
	end := len(nums) - 1
	for start+1 < end {
		mid := start + (end-start)/2
		if nums[mid] <= nums[end] {
			end = mid
		} else {
			end = mid
		}
	}
	if nums[start] < nums[end] {
		return nums[start]
	}
	return nums[end]
}

// 旋转数组找最小值2
func findMin2(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	start := 0
	end := len(nums) - 1
	for start+1 < end {
		for start < end && nums[end] == nums[end-1] {
			end--
		}
		for start < end && nums[start] == nums[start+1] {
			start++
		}
		mid := start + (end-start)/2
		if nums[mid] <= nums[end] {
			end = mid
		} else {
			start = mid
		}
	}
	if nums[start] < nums[end] {
		return nums[start]
	}
	return nums[end]
}

// 旋转数组查找元素
func search2(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	start := 0
	end := len(nums) - 1
	for start+1 < end {
		mid := start + (end-start)/2
		if target == nums[mid] {
			return mid
		}
		if nums[start] < nums[mid] {
			if nums[start] <= target && target < nums[mid] {
				end = mid
			} else {
				start = mid
			}
		} else if nums[end] > nums[mid] {
			if nums[mid] < target && target <= nums[end] {
				start = mid
			} else {
				end = mid
			}
		}
	}
	if nums[start] == target {
		return start
	} else if nums[end] == target {
		return end
	}
	return -1
}

// 旋转数组有重复数字，查找元素
func search3(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	start := 0
	end := len(nums) - 1
	for start+1 < end {
		for start < end && nums[end] == nums[end-1] {
			end--
		}
		for start < end && nums[start] == nums[start+1] {
			start++
		}

		mid := start + (end-start)/2
		if nums[mid] == target {
			return true
		}
		if nums[start] < nums[mid] {
			if nums[start] <= target && target < nums[mid] {
				end = mid
			} else {
				start = mid
			}
		} else if nums[end] > nums[mid] {
			if nums[mid] < target && target <= nums[end] {
				start = mid
			} else {
				end = mid
			}
		}
	}
	if nums[start] == target || nums[end] == target {
		return true
	}
	return false
}

func main() {

}
