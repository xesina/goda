package main

import (
	"errors"
)

const initialTail = -1

var (
	errQueueEmpty = errors.New("queue: queue is empty")
	errQueueFull  = errors.New("queue: queue is full")
)

// Queue represents a queue data structure
type Queue struct {
	data []int
	tail int // index of last item in queue
}

// NewQueue creates and returns a queue of int elements
func NewQueue(size uint) *Queue {
	return &Queue{
		data: make([]int, size),
		tail: initialTail,
	}
}

// Enqueue adds an item to the queue
func (q *Queue) Enqueue(item int) error {
	if q.IsFull() {
		return errQueueFull
	}
	q.tail++
	q.data[q.tail] = item
	return nil
}

// Dequeue returns and removes an item from the queue
func (q *Queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, errQueueEmpty
	}

	d := q.data[0]      // keep 0th item in a temp var
	q.data = q.data[1:] // delete 0th aka first item from queue
	q.tail--

	return d, nil

}

// IsEmpty returns true if queue is empty
func (q *Queue) IsEmpty() bool {
	return q.tail == initialTail
}

// IsFull returns true if queue is full
func (q *Queue) IsFull() bool {
	return q.Size() == len(q.data)
}

// Size returns count of elements in queue
func (q *Queue) Size() int {
	return q.tail + 1
}

// Tail returns the item at end of the queue
func (q *Queue) Tail() (int, error) {
	if q.IsEmpty() {
		return 0, errQueueEmpty
	}
	return q.data[q.tail], nil
}

func main() {
}
