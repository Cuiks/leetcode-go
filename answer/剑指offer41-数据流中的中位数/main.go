package main

import "fmt"

type MedianFinder struct {
	queue []float64
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{queue: make([]float64, 0)}
}

func (this *MedianFinder) AddNum(num int) {
	this.queue = append(this.queue, float64(num))
	index := len(this.queue) - 1
	for index >= 1 && this.queue[index] < this.queue[index-1] {
		this.queue[index], this.queue[index-1] = this.queue[index-1], this.queue[index]
		index--
	}
}

func (this *MedianFinder) FindMedian() float64 {
	length := len(this.queue)
	if length == 0 {
		return 0
	}
	if length%2 == 0 {
		return (this.queue[length/2] + this.queue[length/2-1]) / 2
	} else {
		return this.queue[length/2]
	}
}

func main() {
	c := Constructor()
	fmt.Println(c.FindMedian())
	c.AddNum(-1)
	fmt.Println(c.FindMedian())
	c.AddNum(-2)
	fmt.Println(c.FindMedian())
	c.AddNum(-3)
	fmt.Println(c.FindMedian())
	c.AddNum(-4)
	fmt.Println(c.FindMedian())
	c.AddNum(-5)

}
