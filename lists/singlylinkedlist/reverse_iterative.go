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
