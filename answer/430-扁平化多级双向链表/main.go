package main

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}
	head := root
	_ = getFlatten(root)
	return head
}

func getFlatten(root *Node) *Node {
	var pre *Node
	for root != nil {
		if root.Child != nil {
			next := root.Next
			root.Next = root.Child
			root.Child.Prev = root
			root.Child = nil
			root = getFlatten(root.Next)
			root.Next = next
			if next!=nil{
				next.Prev = root
			}
		}
		pre = root
		root = root.Next
	}
	return pre
}

func main() {

}
