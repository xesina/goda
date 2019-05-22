package bst

import "fmt"

// Node represents a binary search tree node
type Node struct {
	left  *Node
	data  int
	right *Node
}

// NewNode creates and returns a new node
func NewNode(item int) *Node {
	return &Node{
		data: item,
	}
}

// Insert inserts a new node to binary search tree
func (n *Node) Insert(item int) {
	if n == nil {
		n = &Node{data: item}
		return
	}

	if item < n.data {
		if n.left == nil {
			n.left = &Node{data: item}
			return
		}
		n.left.Insert(item)
		return
	}

	if n.right == nil {
		n.right = &Node{data: item}
		return
	}

	n.right.Insert(item)

}

// InorderTraverse traverses a tree in a inorder fashion
func (n *Node) InorderTraverse() {
	if n == nil {
		return
	}

	n.left.InorderTraverse()

	fmt.Printf("%d ", n.data)

	n.right.InorderTraverse()

}

// PreorderTraverse traverses a tree in a preorder fashion
func (n *Node) PreorderTraverse() {
	if n == nil {
		return
	}

	fmt.Printf("%d ", n.data)

	n.left.PreorderTraverse()

	n.right.PreorderTraverse()

}

// Insert inserts a new node into bst and returns root
func Insert(root *Node, data int) *Node {
	if root == nil {
		root = NewNode(data)
		return root
	}

	if data <= root.data {
		root.left = Insert(root.left, data)
	} else {
		root.right = Insert(root.right, data)
	}

	return root
}

// InorderTraverse traverse the given binary tree in inorder fashion
func InorderTraverse(root *Node) {
	if root == nil {
		return
	}
	InorderTraverse(root.left)
	fmt.Printf("%d ", root.data)
	InorderTraverse(root.right)
}

// GetInorderTraversal returns the inorder traversal of the binary tree
func GetInorderTraversal(root *Node) []int {
	var nodes []int
	if root == nil {
		return nodes
	}

	tmp := GetInorderTraversal(root.left)
	nodes = append(nodes, tmp...)
	nodes = append(nodes, root.data)
	tmp = GetInorderTraversal(root.right)
	nodes = append(nodes, tmp...)

	return nodes
}

// GetLevelOrderTraversal ...
func GetLevelOrderTraversal(root *Node) (nodes []int) {
	q := newQueue()
	q.enqueue(root)
	for !q.isEmpty() {
		node := q.dequeue()
		nodes = append(nodes, node.data)
		if node.left != nil {

			q.enqueue(node.left)
		}
		if node.right != nil {

			q.enqueue(node.right)
		}
	}

	return
}

// Max will find max item in binary tree
func Max(root *Node) int {
	node := root
	var max int

	for node != nil {
		max = node.data
		node = node.right
	}

	return max
}

// RecursiveMax will return max item in binary tree in a recursive fashion
func RecursiveMax(root *Node) int {
	if root == nil {
		return 0
	}

	max := root.data
	if root.right != nil {
		max = RecursiveMax(root.right)
	}

	return max
}

// Min Return min item in a binary tree
func Min(root *Node) int {
	node := root
	var min int

	for node != nil {
		min = node.data
		node = node.left
	}

	return min
}

// RecursiveMin ...
func RecursiveMin(root *Node) int {
	var min int
	if root == nil {
		return min
	}

	min = root.data

	if root.left != nil {
		min = RecursiveMin(root.left)
	}

	return min

}

// Height will return height of a binary tree
func Height(root *Node) int {
	var leftHeight, rightHeight int
	if root == nil {
		return -1
	}
	leftHeight = Height(root.left)
	rightHeight = Height(root.right)

	return max(leftHeight, rightHeight) + 1

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
