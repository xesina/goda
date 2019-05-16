package doublylinkedlist

import "fmt"

// Node represents a node of doubly linked list
type Node struct {
	data     int
	previous *Node
	next     *Node
}

// List represents a doubly linked list
type List struct {
	head *Node
}

// New create a new empty douby linked list
func New() *List {
	return &List{
		head: nil,
	}
}

// Add inserts a new node at the end of doubly linked list
func (l *List) Add(item int) {
	node := l.head
	if node == nil {
		node = &Node{
			data: item,
		}

		l.head = node

		return
	}

	for node.next != nil {
		node = node.next
	}

	newNode := &Node{
		data:     item,
		previous: node,
	}

	node.next = newNode

	return
}

// AddToFront adds a node to front of the doubly linked list
func (l *List) AddToFront(item int) {
	newNode := &Node{
		data:     item,
		next:     l.head,
		previous: nil,
	}

	if l.head != nil {
		l.head.previous = newNode
	}

	l.head = newNode
}

// Print prints the doubly linked list nodes data
func (l *List) Print() {
	if l.head == nil {
		fmt.Println("list is empty.")
		return
	}

	node := l.head

	for node != nil {
		fmt.Printf("%d ", node.data)

		node = node.next
	}

	fmt.Println()
}
