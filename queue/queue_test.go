package main

import (
	"testing"
)

func TestNewQueue(t *testing.T) {
	cases := []struct {
		size uint
	}{
		{size: 0},
		{size: 4},
		{size: 1},
		{size: 100},
	}

	for _, c := range cases {
		q := NewQueue(c.size)

		if len(q.data) != int(c.size) {
			t.Errorf("len(q.data) = %d, want len(q.data) = %d", len(q.data), c.size)
		}
	}
}

func TestEnqueue(t *testing.T) {
	cases := []struct {
		size    uint  // size indicates the queue size
		items   []int // initial queue items
		enqueue int   // new item that we're going to enqueue
		tail    int   // tail item after enqueue
		err     error // enqueue error
	}{
		{
			size:    5,
			items:   []int{1, 2, 3},
			enqueue: 4,
			tail:    4,
			err:     nil,
		},
		{
			size:    0,
			items:   []int{},
			enqueue: 4,
			tail:    0,
			err:     errQueueFull,
		},
		{
			size:    3,
			items:   []int{1, 2, 3},
			enqueue: 4,
			tail:    3,
			err:     errQueueFull,
		},
	}

	for _, c := range cases {
		q := loadToQueue(c.size, c.items)
		err := q.Enqueue(c.enqueue)

		if err != c.err {
			t.Fatalf(`err = "%s", want err = "%s"`, err, c.err)
		}

		tail, _ := q.Tail()
		if tail != c.tail {
			t.Fatalf("tail = %d, want tail = %d", tail, c.tail)
		}

	}
}

func TestTail(t *testing.T) {
	cases := []struct {
		size  uint  // size indicates the queue size
		items []int // initial queue items
		want  int   // tailed item
		err   error // tail error
	}{
		{
			size:  5,
			items: []int{1, 2, 3},
			want:  3,
			err:   nil,
		},
		{
			size:  0,
			items: []int{},
			want:  0,
			err:   errQueueEmpty,
		},
		{
			size:  1,
			items: []int{1},
			want:  1,
			err:   nil,
		},
	}

	for _, c := range cases {
		q := loadToQueue(c.size, c.items)
		tail, err := q.Tail()

		if err != c.err {
			t.Fatalf(`err = "%s", want err = "%s"`, err, c.err)
		}

		if tail != c.want {
			t.Fatalf("tail = %d, want tail = %d", tail, c.want)
		}
	}
}

func TestDequeue(t *testing.T) {
	cases := []struct {
		size  uint  // size indicates the queue size
		items []int // initial queue items
		want  int   // dequeued item
		err   error // dequeue error
	}{
		{
			size:  5,
			items: []int{1, 2, 3},
			want:  1,
			err:   nil,
		},
		{
			size:  0,
			items: []int{},
			want:  0,
			err:   errQueueEmpty,
		},
		{
			size:  1,
			items: []int{1},
			want:  1,
			err:   nil,
		},
		{
			size:  5,
			items: []int{5, 4, 3, 2, 1},
			want:  5,
			err:   nil,
		},
	}

	for _, c := range cases {
		q := loadToQueue(c.size, c.items)
		d, err := q.Dequeue()

		if err != c.err {
			t.Fatalf(`err = "%s", want err = "%s"`, err, c.err)
		}

		if d != c.want {
			t.Fatalf("dequeued = %d, want dequeued = %d", d, c.want)
		}
	}
}

func TestIsEmpty(t *testing.T) {
	cases := []struct {
		size  uint  // size indicates the queue size
		items []int // initial queue items
		want  bool  // empty or not result
	}{
		{
			size:  5,
			items: []int{1, 2, 3},
			want:  false,
		},
		{
			size:  0,
			items: []int{},
			want:  true,
		},
		{
			size:  1,
			items: []int{1},
			want:  false,
		},
	}

	for _, c := range cases {
		q := loadToQueue(c.size, c.items)
		empty := q.IsEmpty()

		if empty != c.want {
			t.Fatalf("empty = %v, want empty = %v", empty, c.want)
		}
	}
}

func TestIsFull(t *testing.T) {
	cases := []struct {
		size  uint  // size indicates the queue size
		items []int // initial queue items
		want  bool  // full or not resulr
	}{
		{
			size:  5,
			items: []int{1, 2, 3},
			want:  false,
		},
		{
			size:  0,
			items: []int{},
			want:  true,
		},
		{
			size:  1,
			items: []int{1},
			want:  true,
		},
	}

	for _, c := range cases {
		q := loadToQueue(c.size, c.items)
		full := q.IsFull()

		if full != c.want {
			t.Fatalf("full = %v, want full = %v", full, c.want)
		}
	}
}

func TestSequenceOrder(t *testing.T) {
	cases := []struct {
		size    uint  // size indicates the queue size
		enqueue []int // enqueued items
		dequeue []int // dequeued items
	}{
		{
			size:    5,
			enqueue: []int{1, 2, 3},
			dequeue: []int{1, 2, 3},
		},
	}

	for _, c := range cases {
		q := loadToQueue(c.size, c.enqueue)
		for _, v := range c.enqueue {
			tmp, _ := q.Dequeue()
			if tmp != v {
				t.Fatalf("got %d, want %d", tmp, v)
			}
		}
	}
}

// loadToQueue is a testing helper func to create and load a Queue with items
func loadToQueue(size uint, arr []int) *Queue {
	s := NewQueue(size)
	for _, v := range arr {
		s.Enqueue(v)
	}

	return s
}
