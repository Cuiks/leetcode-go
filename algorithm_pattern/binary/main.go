package main

/*
a = a^0 = 0^a
0 = a^a
a = a^b^b

// 交换两个数
a = a^b
b = a^b
a = a^b

// 移除最后一个1
a = n&(n-1)

// 获取最后一个1
diff=(n&(n-1))^n
*/

// 除了某个元素只出现一次以外，其余每个元素均出现两次。找出只出现一次的数字
func singleNumber(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		result = result ^ nums[i]
	}
	return result
}

// 除了某个元素只出现一次以外，其余每个元素均出现三次。找出只出现一次的数字
func singleNumber2(nums []int) int {
	result := 0
	for i := 0; i < 64; i++ {
		sum := 0
		for j := 0; j < len(nums); j++ {
			sum += (nums[j] >> i) & 1
		}
		result ^= (sum % 3) << i
	}
	return result
}

// 有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素
func singleNumber3(nums []int) []int {
	diff := 0
	for i := 0; i < len(nums); i++ {
		diff ^= nums[i]
	}
	result := []int{diff, diff}

	diff = (diff & (diff - 1)) ^ diff
	for i := 0; i < len(nums); i++ {
		if diff&nums[i] == 0 {
			result[0] ^= nums[i]
		} else {
			result[1] ^= nums[i]
		}
	}
	return result
}

// 汉明重量
func hammingWeight(num uint32) int {
	res := 0
	for num != 0 {
		num = num & (num - 1)
		res++
	}
	return res
}

// 计算0~nums中数的二进制的1的数量
func countBits(num int) []int {
	result := make([]int, num+1)
	for i := 0; i <= num; i++ {
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

// 动态规划解法
func countBits2(num int) []int {
	result := make([]int, num+1)
	for i := 1; i <= num; i++ {
		result[i] = result[i&(i-1)] + 1
	}
	return result
}

// 颠倒给定的 32 位无符号整数的二进制位。
func reverseBits(num uint32) uint32 {
	var res uint32
	pow := 31

	for num != 0 {
		res += (num & 1) << pow
		num >>= 1
		pow--
	}
	return res
}

// 返回[m, n]范围内所有数字的按位与
func rangeBitwiseAnd(m int, n int) int {
	counts := 0
	for m != n {
		m >>= 1
		n >>= 1
		counts++
	}
	return m << counts
}

func main() {

}
