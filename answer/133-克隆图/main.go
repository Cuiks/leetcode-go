package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	visited := make(map[*Node]*Node, 0)
	return clone(node, visited)
}

func clone(node *Node, visited map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}

	v, ok := visited[node]
	if ok {
		return v
	}

	clonedNode := &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, len(node.Neighbors)),
	}
	visited[node] = clonedNode

	for i := 0; i < len(node.Neighbors); i++ {
		clonedNode.Neighbors[i] = clone(node.Neighbors[i], visited)
	}
	return clonedNode
}

func main() {

}
