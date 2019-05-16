package doublylinkedlist

import "testing"

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
			if l.head.data != c.want {
				t.Fatalf("l.head.data = %d, want l.head.data = %d", l.head.data, c.want)
			}

			return
		}

		if last.next.data != c.want {
			t.Fatalf("last.next.data = %d, want last.next.data = %d", last.next.data, c.want)
		}

	}
}

func TestTraverse(t *testing.T) {
	cases := []struct {
		items []int
	}{
		{
			items: []int{1, 2, 3, 4, 5},
		},
		{
			items: []int{},
		},
	}

	for _, c := range cases {
		l, _ := loadToList(c.items)
		l.Print()
	}

}

func TestAddToFront(t *testing.T) {
	cases := []struct {
		items []int // initial list items
		add   int   // item to be added to front of the list
		want  int   // head item after add
	}{
		{
			items: []int{2, 3, 4},
			add:   1,
			want:  1,
		},
		{
			items: []int{1, 2},
			add:   0,
			want:  0,
		},
		{
			items: []int{},
			add:   1,
			want:  1,
		},
	}

	for _, c := range cases {
		l, _ := loadToList(c.items)
		l.AddToFront(c.add)

		if l.head.data != c.want {
			t.Fatalf("l.head.data == %d, want l.head.data = %d", l.head.data, c.want)
		}
	}
}

func loadToList(items []int) (*List, *Node) {
	l := New()
	if len(items) == 0 {
		return l, nil
	}

	l.head = &Node{data: items[0]}
	last := l.head
	for i := 1; i < len(items); i++ {
		last.next = &Node{
			data:     items[i],
			previous: last,
		}
		last = last.next
	}

	return l, last
}
