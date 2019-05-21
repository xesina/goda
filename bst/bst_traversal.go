package bst

import "errors"

type stack struct {
	data []*Node
	top  int
}

func newStack() *stack {
	return &stack{
		top: -1,
	}
}

func (s *stack) push(item *Node) {
	s.data = append(s.data, item)
	s.top++
}

func (s *stack) pop() (*Node, error) {
	if s.top == -1 {
		return nil, errors.New("stack is empty")
	}

	p := s.data[s.top]
	s.data = s.data[0:s.top]
	s.top--
	return p, nil
}

// IterativeInorderTraversal traverses the binary tree in inorder fashion and iteratively!
func IterativeInorderTraversal(root *Node) []int {
	var nodes []int
	m := make(map[*Node]bool)

	s := newStack()
	node := root
	for node != nil {

		// first visit the left subtree if it exists and not visited before
		if node.left != nil && m[node.left] != true {
			s.push(node)
			node = node.left
			continue
		}

		// if we visted the left subtree or it doesn't exists then we add current node and
		// mark it as visited
		if m[node] != true {
			nodes = append(nodes, node.data)
			m[node] = true
		}

		// first visit the right subtree if it exists and not visited before
		if node.right != nil && m[node.right] != true {
			s.push(node)
			node = node.right
			continue
		}

		node, _ = s.pop()

	}

	return nodes

}
