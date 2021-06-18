package main

func reverseBits(num uint32) uint32 {
	var result uint32
	pow := 31
	for num != 0 {
		result += (num & 1) << pow
		pow--
		num >>= 1
	}
	return result
}
func main() {

}
