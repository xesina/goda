package main

import (
	"testing"
)
func TestNewStack(t *testing.T)  {
	size := 4
	s := NewStack(size)
	
	if s.data == nil {
		t.Errorf("s.data is nil, want not nil")
	}

	if len(s.data) != size {
		t.Errorf("len(s.data) = %d, want %d", len(s.data), size)
	}	
}

func TestSize(t *testing.T)  {
	size := 4
	s := NewStack(size)
	
	if s.Size() != size {
		t.Errorf("s.Size() = %d, want %d", s.Size(), size)
	}	
}

func TestPop(t *testing.T)  {
	size := 4
	item:= 1
	s := NewStack(size)
	s.Push(item)

	p := s.Pop()
	if p != item {
		t.Errorf("s.Pop() = %d, want %d", p, item )
	} 
}

func TestPush(t *testing.T)  {
	size := 4
	item := 1
	s := NewStack(size)

	s.Push(item)
	p := s.Pop()
	if p != item {
		t.Errorf("s.Push(%d), got s.Pop() = %d, want %d", item, p, item)
	} 
}