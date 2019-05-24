package heap

import "testing"

func TestLeftChildIndex(t *testing.T) {
	cases := []struct {
		name  string
		items []int
		index int
		want  int
	}{
		{
			name:  "LeftChildIndex(0)==1",
			items: []int{7, 8, 9},
			index: 0,
			want:  1,
		},
		{
			name:  "LeftChildIndex(0)==0",
			items: []int{7},
			index: 0,
			want:  1,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			m := NewMinHeap()
			m.data = c.items
			i := m.leftChildIndex(c.index)
			if i != c.want {
				t.Fatalf("got index=%d, want index=%d", i, c.want)
			}
		})
	}
}

func TestLeftChild(t *testing.T) {
	cases := []struct {
		name  string
		items []int
		index int
		want  int
	}{
		{
			name:  "leftChild(0)=8",
			items: []int{7, 8, 9},
			index: 0,
			want:  8,
		},
		{
			name:  "leftChild(0)=0",
			items: []int{9},
			index: 0,
			want:  0,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			m := NewMinHeap()
			m.data = c.items
			l := m.leftChild(c.index)
			if l != c.want {
				t.Fatalf("got left = %d, want left = %d", l, c.want)
			}
		})
	}
}

func TestRightChild(t *testing.T) {
	cases := []struct {
		name  string
		items []int
		index int
		want  int
	}{
		{
			name:  "rightChild(0)==9",
			items: []int{7, 8, 9},
			index: 0,
			want:  9,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			m := NewMinHeap()
			m.data = c.items
			r := m.rightChild(c.index)
			if r != c.want {
				t.Fatalf("got right = %d, want right = %d", r, c.want)
			}
		})
	}
}

func TestPush(t *testing.T) {
	cases := []struct {
		items []int
		push  int
		want  []int
	}{
		{
			items: []int{10, 11, 12, 13, 14},
			push:  9,
			want:  []int{9, 11, 10, 13, 14, 12},
		},
		{
			items: []int{10},
			push:  9,
			want:  []int{9, 10},
		},
		{
			items: []int{},
			push:  9,
			want:  []int{9},
		},
	}

	for _, c := range cases {

		m := NewMinHeap()
		m.data = c.items
		m.Push(c.push)
		for k := range c.want {
			if m.data[k] != c.want[k] {
				t.Fatalf("got %+v, want %+v", m.data, c.want)
			}
		}
	}

}

func TestPop(t *testing.T) {
	cases := []struct {
		name  string
		items []int
		poped int
		want  []int
	}{
		{
			name:  "test pop1",
			items: []int{10, 11, 12, 13, 14},
			poped: 10,
			want:  []int{11, 13, 12, 14},
		},
		{
			name:  "test pop2",
			items: []int{22, 33, 44},
			poped: 22,
			want:  []int{33, 44},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			m := NewMinHeap()
			m.data = c.items
			p := m.Pop()
			if p != c.poped {
				t.Fatalf("got poped = %d, want poped = %d", p, c.poped)
			}

			for k := range c.want {
				if m.data[k] != c.want[k] {
					t.Fatalf("got %+v, want %+v", m.data, c.want)
				}
			}
		})
	}

}
