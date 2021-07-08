package main

import (
	"fmt"
	"strconv"
)

var min = 1<<31 - 1

func openLock(deadends []string, target string) int {
	deadMap := map[string]bool{}
	for i:=0;i<len(deadends);i++{
		deadMap[deadends[i]] = true
	}
	var dfs func(int, string)
	dfs = func(step int, value string) {
		fmt.Println(step, value, strCom(value, "9999"))
		if strCom(value, "9999") {
			return
		}
		if hit(deadMap, value) {
			return
		}
		if step < min && value == target {
			min = step
			return
		}
		dfs(step+1, addStr(value, 1))
		dfs(step+1, addStr(value, 10))
		dfs(step+1, addStr(value, 100))
		dfs(step+1, addStr(value, 1000))
		return
	}
	dfs(0, "0000")
	return min
}

func strCom(a, b string) bool {
	aInt, _ := strconv.Atoi(a)
	bInt, _ := strconv.Atoi(b)
	if aInt > bInt {
		return true
	}
	return false
}

func hit(deadMap map[string]bool, value string) bool {
	_,ok := deadMap[value]
	return ok
}

func addStr(value string, add int) string {
	v, _ := strconv.Atoi(value)
	v += add
	res := ""
	vStr := strconv.Itoa(v)
	for i := 0; i < len(value)-len(vStr); i++ {
		res += "0"
	}
	return res + vStr
}

func main() {
	dead := []string{"0201", "0101", "0102", "1212", "2002"}
	target := "0202"
	fmt.Println(openLock(dead, target))
}
