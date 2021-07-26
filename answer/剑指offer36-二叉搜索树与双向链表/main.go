package main

// 未验证对错

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func treeToDoublyList(root *Node) *Node {
	if root == nil {
		return nil
	}
	root = helper(root)

	minRoot := minNode(root)
	maxRoot := maxNode(root)
	minRoot.Left = maxRoot
	maxRoot.Right = minRoot

	return root
}

func helper(root *Node) *Node {
	if root == nil {
		return nil
	}
	helper(root.Left)
	helper(root.Right)
	leftMax := maxNode(root.Left)
	rightMin := minNode(root.Right)
	// 连接root
	root.Left = leftMax
	leftMax.Right = root
	root.Right = rightMin
	rightMin.Left = root

	return root
}

func minNode(root *Node) *Node {
	if root == nil {
		return nil
	}
	min := root
	for root != nil {
		if root.Left != nil {
			min = root.Left
		}
		root = root.Left
	}
	return min
}

func maxNode(root *Node) *Node {
	if root == nil {
		return nil
	}
	max := root
	for root.Right != nil {
		if root.Right != nil {
			max = root.Right
		}
		root = root.Right
	}
	return max
}

func main() {

}
