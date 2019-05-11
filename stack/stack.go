package main 

// Stack represents a stack of int elements
type Stack struct {
	data []int
}

// NewStack creates and returns a new stack
func NewStack(size int) *Stack {
	return &Stack{
		data: make([]int, size),
	}
}

// Pop removes and returns the top element on stack 
func (s *Stack) Pop() int  {
	return 0
}

// Push add new element to top of stack
func (s *Stack) Push(item int)  {
	
}

// Size returns the size of stack
func (s *Stack) Size() int {
	return 0
}

func main()  {
	
}