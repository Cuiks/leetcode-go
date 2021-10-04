package main

import (
	"fmt"
	"strings"
)

func licenseKeyFormatting(s string, k int) string {
	newStr := strings.ToUpper(strings.Replace(s, "-", "", -1))
	length := len(newStr)
	mod := length % k

	res := newStr[:mod]
	i := mod
	for i < length {
		res += "-" + newStr[i:i+k]
		i += k
	}
	res = strings.Trim(res, "-")
	return res
}

func main() {
	fmt.Println(licenseKeyFormatting("5F3Z-2e-9-w", 4))
}
