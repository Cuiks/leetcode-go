package main

type MaxQueue struct {
	queue    []int
	maxQueue []int
}

func Constructor() MaxQueue {
	return MaxQueue{
		queue:    make([]int, 0),
		maxQueue: make([]int, 0),
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.maxQueue) == 0 {
		return -1
	}
	return this.maxQueue[0]
}

func (this *MaxQueue) Push_back(value int) {
	for len(this.maxQueue) > 0 && value > this.maxQueue[len(this.maxQueue)-1] {
		this.maxQueue = this.maxQueue[:len(this.maxQueue)-1]
	}
	this.maxQueue = append(this.maxQueue, value)
	this.queue = append(this.queue, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.queue) == 0 {
		return -1
	}
	val := this.queue[0]
	this.queue = this.queue[1:]
	if val == this.maxQueue[0] {
		this.maxQueue = this.maxQueue[1:]
	}
	return val
}
func main() {

}
