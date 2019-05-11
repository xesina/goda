package main

import (
	"testing"
)

func TestNewStack(t *testing.T) {
	cases := []struct {
		size uint
	}{
		{size: 4},
		{size: 1},
		{size: 2},
		{size: 0},
		{size: 10000},
	}

	for _, c := range cases {
		s := NewStack(c.size)
		if s.data == nil {
			t.Errorf("s.data is nil, want not nil")
		}

		if len(s.data) != int(c.size) {
			t.Errorf("len(s.data) = %d, want %d", len(s.data), int(c.size))
		}
	}

}

func TestSize(t *testing.T) {
	cases := []struct {
		size uint
	}{
		{size: 4},
		{size: 1},
		{size: 2},
		{size: 0},
	}

	for _, c := range cases {
		s := NewStack(c.size)
		if s.Size() != len(s.data) || s.Size() != int(c.size) {
			t.Errorf("s.Size() = %d, want %d", s.Size(), int(c.size))
		}
	}
}

func TestPush(t *testing.T) {
	cases := []struct {
		stack []int // initial stack items
		size  int   // size indicates the stack size
		push  int   // new item that we're going to push
		want  int   // peek item after push
		err   error // push error
	}{
		{
			stack: []int{},
			size:  0,
			push:  1,
			want:  0,
			err:   errStackFull,
		},
		{
			stack: []int{1, 2, 3},
			size:  6,
			push:  4,
			want:  4,
			err:   nil,
		},
	}

	for _, c := range cases {
		s := loadToStack(c.size, c.stack)
		err := s.Push(c.push)
		if err != c.err {
			t.Fatalf("err = \"%s\", want err = \"%s\"", err, c.err)
		}
		p, _ := s.Peek()
		if p != c.want {
			t.Fatalf("got %d, want %d", p, c.want)
		}
	}

}

func TestPop(t *testing.T) {
	cases := []struct {
		stack []int // initial stack items
		want  int   // item that should be returned and removed from stack
		err   error // push error
	}{
		{
			stack: []int{},
			want:  0,
			err:   errStackEmpty,
		},
		{
			stack: []int{1, 2, 3},
			want:  3,
			err:   nil,
		},
		{
			stack: []int{1},
			want:  1,
			err:   nil,
		},
	}

	for _, c := range cases {
		s := loadToStack(len(c.stack), c.stack)
		top := s.top
		p, err := s.Pop()
		if err != c.err {
			t.Fatalf("err = \"%s\", want err = \"%s\"", err, c.err)
		}

		if p != c.want {
			t.Fatalf("got %d, want %d", p, c.want)
		}

		if err != errStackEmpty && s.top != top-1 {
			t.Fatalf("top = %d, want top = %d", s.top, top-1)
		}
	}

}

func TestPeek(t *testing.T) {
	cases := []struct {
		items []int
		want  int
		err   error
	}{
		{
			items: []int{},
			want:  0,
			err:   errStackEmpty,
		},
		{
			items: []int{1, 2, 3},
			want:  3,
			err:   nil,
		},
	}

	for _, c := range cases {
		s := loadToStack(len(c.items), c.items)
		p, err := s.Peek()
		if err != c.err {
			t.Fatal(err)
		}
		if p != c.want {
			t.Fatalf("got %d, want %d", p, c.want)
		}

	}
}

func loadToStack(size int, arr []int) *Stack {
	s := NewStack(uint(size))
	for _, v := range arr {
		s.Push(v)
	}

	return s
}
