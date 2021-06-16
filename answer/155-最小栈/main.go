package main

type MinStack struct {
	stack []int
	min   []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack: make([]int, 0),
		min:   make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	min := this.GetMin()
	if val > min {
		this.min = append(this.min, min)
	} else {
		this.min = append(this.min, val)
	}
	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	this.stack = this.stack[:len(this.stack)-1]
	this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return 0
	}
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.min) == 0 {
		return 1 << 31
	}
	return this.min[len(this.min)-1]
}

func main() {

}
