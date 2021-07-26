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
		node := &Node{
			Val:  cur.Val,
			Next: cur.Next,
		}
		cur.Next = node
		cur = node.Next
	}

	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}

	cur = head
	clonedHead := cur.Next
	for cur.Next != nil {
		temp := cur.Next
		cur.Next = cur.Next.Next
		cur = temp
	}

	return clonedHead
}

func main() {

}
