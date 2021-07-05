package main

func findAnagrams(s string, p string) []int {
	need := map[byte]int{}
	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}

	result := make([]int, 0)
	win := map[byte]int{}
	left, right := 0, 0
	match := 0
	for right < len(s) {
		c := s[right]
		right++
		if need[c] != 0 {
			win[c]++
			if need[c] == win[c] {
				match++
			}
		}
		for right-left >= len(p) {
			if right-left == len(p) && match == len(need) {
				result = append(result, left)
			}
			c = s[left]
			left++
			if need[c] != 0 {
				if need[c] == win[c] {
					match--
				}
				win[c]--
			}
		}
	}
	return result
}

func main() {

}
