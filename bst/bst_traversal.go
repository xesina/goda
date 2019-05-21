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

func (s *stack) isEmpty() bool {
	return s.top == -1
}

func inorderTraversal(root *Node) []int {
	var nodes []int

	s := newStack()
	node := root
	for node != nil || !s.isEmpty() {
		for node != nil {
			s.push(node)
			node = node.left
		}

		node, _ = s.pop()
		nodes = append(nodes, node.data)
		node = node.right

	}

	return nodes

}
