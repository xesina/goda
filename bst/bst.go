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
