package main

import (
	"errors"
)

const initialTop = -1

var (
	errStackEmpty = errors.New("stack: stack is empty")
	errStackFull  = errors.New("stack: stack is full")
)

// Stack represents a stack of int elements
// Notice: This isn't concurrent safe stack
type Stack struct {
	data []int
	top  int
}

// NewStack creates and returns a new stack
func NewStack(size uint) *Stack {
	return &Stack{
		data: make([]int, size),
		top:  initialTop,
	}
}

// Pop removes and returns the top element on stack
func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errStackEmpty
	}

	item := s.data[s.top]
	s.top--

	return item, nil
}

// Peek returns the value at the top of the stack, but does not remove it.
func (s *Stack) Peek() (int, error) {
	if !s.IsEmpty() {
		return s.data[s.top], nil
	}

	return 0, errStackEmpty
}

// Push add new element to top of stack
func (s *Stack) Push(item int) error {
	if s.IsFull() {
		return errStackFull
	}

	s.top++
	s.data[s.top] = item

	return nil
}

// Size returns the size of stack
func (s *Stack) Size() int {
	return len(s.data)
}

// IsEmpty returns true if stack is empty
func (s *Stack) IsEmpty() bool {
	return s.top == initialTop
}

// IsFull returns true if stack is full
func (s *Stack) IsFull() bool {
	return s.top == s.Size()-1
}

func main() {

}
