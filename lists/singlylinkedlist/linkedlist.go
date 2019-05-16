package linkedlist

import (
	"fmt"
	"io"
	"os"
)

// Node represents a node in a linedlist
type Node struct {
	data int
	next *Node
}

// List represents a Singly LinkedList
type List struct {
	head   *Node
	output io.WriteCloser
}

// New creates and returns a singly linkedlist
func New() *List {
	return &List{
		head:   nil,
		output: os.Stdout,
	}
}

// Add adds a new node to the end of linkedlist
func (l *List) Add(data int) {
	if l.head == nil {
		l.head = &Node{
			data: data,
			next: nil,
		}
		return
	}

	newNode := &Node{
		data: data,
		next: nil,
	}

	node := l.head
	for node.next != nil {
		node = node.next
	}

	node.next = newNode
}

// Head returns the head node of the list
func (l *List) Head() *Node {
	return l.head
}

// Print will print the data item of all linked list nodes
func (l *List) Print() {
	node := l.head
	if node == nil {
		fmt.Fprint(l.output, "linked list is empty.")
	}

	for node != nil {
		fmt.Fprintf(l.output, "%d ", node.data)
		node = node.next
	}

	fmt.Println()

	return
}

// PrintRecursively will print the data item of all linked list nodes in a recursive fashion
func PrintRecursively(head *Node) {
	if head == nil {
		fmt.Println()
		return
	}

	fmt.Printf("%d ", head.data)
	PrintRecursively(head.next)
}

// PrintReverse will print linked list nodes in reverse order
func PrintReverse(head *Node) {
	if head == nil {
		return
	}

	PrintReverse(head.next)
	fmt.Printf("%d ", head.data)
}
