package main

type Iterator struct {
}

func (this *Iterator) hasNext() bool {
	// Returns true if the iteration has more elements.
}

func (this *Iterator) next() int {
	// Returns the next element in the iteration.
}

type PeekingIterator struct {
	nextVal int
	iter    *Iterator
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{
		nextVal: 0,
		iter:    iter,
	}
}

func (this *PeekingIterator) hasNext() bool {
	if this.nextVal != 0 || this.iter.hasNext() {
		return true
	}
	return false
}

func (this *PeekingIterator) next() int {
	this.peek()
	val := this.nextVal
	this.nextVal = 0
	return val
}

func (this *PeekingIterator) peek() int {
	if this.nextVal == 0 {
		this.nextVal = this.iter.next()
	}
	return this.nextVal
}
func main() {

}
