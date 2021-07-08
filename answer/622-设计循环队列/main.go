package main

import "fmt"

type MyCircularQueue struct {
	queue []int
	head  int
	tail  int
}

func Constructor(k int) MyCircularQueue {
	queue := make([]int, k)
	for i := 0; i < k; i++ {
		queue[i] = -1
	}
	return MyCircularQueue{
		queue: queue,
		head:  0,
		tail:  0,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	if !this.IsEmpty() {
		this.tail = (this.tail + 1) % len(this.queue)
	}
	this.queue[this.tail] = value

	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.queue[this.head] = -1
	if this.head == this.tail {
		this.tail = (this.head + 1) % len(this.queue)
	}
	this.head = (this.head + 1) % len(this.queue)
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[this.head]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[this.tail]
}

func (this *MyCircularQueue) IsEmpty() bool {
	if this.head == this.tail && this.queue[this.head] == -1 {
		return true
	}
	return false
}

func (this *MyCircularQueue) IsFull() bool {
	if (this.tail+1)%len(this.queue) == this.head && this.queue[this.head] != -1 {
		return true
	}
	return false
}

func main() {
	c := Constructor(1)
	fmt.Println(c.EnQueue(1))
	fmt.Println(c.EnQueue(2))
	//fmt.Println(c.Rear())
	//fmt.Println(c.Rear())
	//fmt.Println(c.DeQueue())
	//fmt.Println(c.EnQueue(5))
	//fmt.Println(c.Rear())
	//fmt.Println(c.DeQueue())
	//fmt.Println(c.Front())
	//fmt.Println(c.DeQueue())
	//fmt.Println(c.DeQueue())
	//fmt.Println(c.DeQueue())
	fmt.Println(c.queue)
}
