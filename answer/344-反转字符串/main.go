package main

func reverseString(s []byte) {
	if len(s) == 0 {
		return
	}
	res := make([]byte, 0)
	reverse(s, 0, &res)
	for i := 0; i < len(s); i++ {
		s[i] = res[i]
	}
}

func reverse(s []byte, i int, res *[]byte) {
	if i >= len(s) {
		return
	}
	reverse(s, i+1, res)
	*res = append(*res, s[i])
}

func main() {
	s := []byte{'h','e','l','l','o'}
	reverseString(s)
}
