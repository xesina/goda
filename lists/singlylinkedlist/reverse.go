package linkedlist

// Reverse reverses a linked list
func (l *List) Reverse() {
	node := l.head

	var prev, next *Node

	for node != nil {

		next = node.next
		node.next = prev

		// keep previuos visited node
		prev = node
		// move forward
		node = next
	}

	// update head to hew head at end
	l.head = prev

}

// ResursiveReverse will reverse the linked list nodes in a recursive fashion
func (l *List) ResursiveReverse(n *Node) {
	if n.next == nil {
		l.head = n
		return
	}

	l.ResursiveReverse(n.next)

	next := n.next
	next.next = n
	n.next = nil
}
