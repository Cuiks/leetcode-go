package main

type CQueue struct {
	stackHead []int
	stackTail []int
}

func Constructor() CQueue {
	return CQueue{
		stackHead: make([]int, 0),
		stackTail: make([]int, 0),
	}
}

func (this *CQueue) AppendTail(value int) {
	if len(this.stackHead) != 0 {
		for i := len(this.stackHead) - 1; i >= 0; i-- {
			v := this.stackHead[i]
			this.stackTail = append(this.stackTail, v)
			this.stackHead = this.stackHead[:i]
		}
	}
	this.stackTail = append(this.stackTail, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.stackTail) != 0 {
		for i := len(this.stackTail) - 1; i >= 0; i-- {
			v := this.stackTail[i]
			this.stackHead = append(this.stackHead, v)
			this.stackTail = this.stackTail[:i]
		}
	}
	if len(this.stackHead) == 0{
		return -1
	}
	l := len(this.stackHead)
	value := this.stackHead[l-1]
	this.stackHead = this.stackHead[:l-1]
	return value
}

func main() {

}
