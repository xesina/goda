package linkedlist

import (
	"testing"
)

func TestNewList(t *testing.T) {
	l := New()
	if l.Head() != nil {
		t.Fatal("l.head.next is not nil, want nil")
	}
}

func TestAdd(t *testing.T) {
	cases := []struct {
		items []int
		add   int
		want  int
	}{
		{
			items: []int{1, 2, 3},
			add:   4,
			want:  4,
		},
		{
			items: []int{1},
			add:   2,
			want:  2,
		},
		{
			items: []int{},
			add:   1,
			want:  1,
		},
	}

	for _, c := range cases {
		l, last := loadToList(c.items)

		l.Add(c.add)

		if last == nil {
			last = l.Head()
			if last.data != c.want {
				t.Fatalf("head.data = %d, want head.data = %d", last.data, c.want)
			}
			return
		}

		if last.next.data != c.want {
			t.Fatalf("last.next.data = %d, want last.next.data = %d", last.next.data, c.want)
		}

	}
}

func TestHead(t *testing.T) {
	cases := []struct {
		items []int
	}{
		{
			items: []int{1, 2, 3},
		},
		{
			items: []int{1},
		},
		{
			items: []int{},
		},
	}

	for _, c := range cases {
		l, _ := loadToList(c.items)
		// TODO: update the l.output and check test
		l.Print()

	}
}

func TestReverse(t *testing.T) {
	cases := []struct {
		items []int
	}{
		{
			items: []int{1, 2, 3},
		},
		{
			items: []int{1, 2, 3, 4, 5, 6},
		},
		{
			items: []int{1},
		},
		{
			items: []int{},
		},
	}

	for _, c := range cases {
		l, last := loadToList(c.items)
		l.Reverse()
		if l.head != last {
			t.Fatalf("l.head = %d should point to last = %d after reverse", l.head, last)
		}
		// TODO: add complete chain test to make sure reverse works as expected
		l.Print()
	}
}

func loadToList(items []int) (*List, *Node) {
	if len(items) == 0 {
		l := New()
		return l, nil
	}

	l := New()
	l.head = &Node{data: items[0], next: nil}
	last := l.Head()

	for i := 1; i < len(items); i++ {
		last.next = &Node{data: items[i], next: nil}
		last = last.next
	}

	return l, last
}
