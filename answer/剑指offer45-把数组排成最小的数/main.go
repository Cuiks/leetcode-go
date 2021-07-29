package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func minNumber(nums []int) string {
	sort.Sort(sortNum(nums))
	numsStr := make([]string, len(nums))
	for i, v := range nums {
		numsStr[i] = strconv.Itoa(v)
	}
	return strings.Join(numsStr, "")
}

type sortNum []int

func (s sortNum) Len() int {
	return len(s)
}

func (s sortNum) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortNum) Less(i, j int) bool {
	iStr := strconv.Itoa(s[i])
	jStr := strconv.Itoa(s[j])
	if iStr+jStr < jStr+iStr {
		return true
	}
	return false
}

func main() {
	fmt.Println(minNumber([]int{3, 30, 34, 5, 9}))
}
