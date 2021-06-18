package main

func isBadVersion(version int) bool {
	return false
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
func main() {

}
