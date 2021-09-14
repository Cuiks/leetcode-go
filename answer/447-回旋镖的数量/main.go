package main

func numberOfBoomerangs(points [][]int) int {
	res := 0
	for _, point := range points {
		record := map[int]int{}
		for _, node := range points {
			record[(node[0]-point[0])*(node[0]-point[0])+(node[1]-point[1])*(node[1]-point[1])]++
		}
		for _, v := range record {
			res += v * (v - 1)
		}
	}
	return res
}

func main() {

}
