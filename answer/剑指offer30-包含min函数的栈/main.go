package main

import (
	"math"
)

type MinStack struct {
	stack    []int
	minStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack:    make([]int, 0),
		minStack: make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	currMin := this.Min()
	if x < currMin {
		this.minStack = append(this.minStack, x)
	} else {
		this.minStack = append(this.minStack, currMin)
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) > 0 {
		this.stack = this.stack[:len(this.stack)-1]
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.stack) > 0 {
		return this.stack[len(this.stack)-1]
	}
	return 0
}

func (this *MinStack) Min() int {
	if len(this.minStack) == 0 {
		return math.MaxInt64
	}
	return this.minStack[len(this.minStack)-1]
}

func main() {

}
