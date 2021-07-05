package main

func lengthOfLongestSubstring(s string) int {
	if len(s)==0{
		return 0
	}
	win := map[byte]int{}
	left, right := 0, 0
	max := 1
	for right < len(s) {
		c := s[right]
		right++
		win[c]++
		for win[c] > 1 {
			d := s[left]
			left++
			win[d]--
		}
		if right-left > max {
			max = right - left
		}
	}
	return max
}

func main() {

}
