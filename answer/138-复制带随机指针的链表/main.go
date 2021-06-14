package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	cur := head
	for cur != nil {
		clone := &Node{Val: head.Val, Next: cur.Next}
		cur.Next = clone
		cur = clone.Next
	}

	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}

	cur = head
	clonedHead := head.Next
	for cur != nil && cur.Next != nil {
		temp := cur.Next
		cur.Next = cur.Next.Next
		cur = temp
	}
	return clonedHead
}

func main() {

}
