package main

import (
	"errors"
)

const initialTop = -1

var (
	errQueueEmpty = errors.New("queue: queue is empty")
	errQueueFull  = errors.New("queue: queue is full")
)

// Queue represents a queue data structure
type Queue struct {
	data []int
	head int
}

// NewQueue creates and returns a queue of int elements
func NewQueue(size uint) *Queue {
	return &Queue{
		data: make([]int, size),
	}
}

// Enqueue adds an item to the queue
func (q *Queue) Enqueue(item int) error {
	return nil
}

// Dequeue returns and removes an item from the queue
func (q *Queue) Dequeue() (int, error) {
	return 0, nil
}

// IsEmpty returns true if queue is empty
func (q *Queue) IsEmpty() bool {
	return true
}

// IsFull returns true if queue is full
func (q *Queue) IsFull() bool {
	return true
}

// Tail returns the item at end of the queue
func (q *Queue) Tail() (int, error) {
	return 0, nil
}

func main() {
}
