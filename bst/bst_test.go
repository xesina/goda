package bst

import (
	"fmt"
	"testing"
)

func TestNodeInsert(t *testing.T) {
	root := testBST()

	root.Insert(8)

	root2 := NewNode(2)
	root2.left = NewNode(1)
	root2.right = NewNode(3)

	root.InorderTraverse()
	fmt.Println()
	root.PreorderTraverse()
}

func TestInsert(t *testing.T) {
	root := testBST()
	root = Insert(root, 11)

	if root.left.left.right.data != 11 {
		t.Fatalf("got %d, want %d", root.left.left.right.data, 11)
	}

	root = Insert(root, 33)
	if root.right.right.right.data != 33 {
		t.Fatalf("got %d, want %d", root.right.right.right, 33)
	}

	root.InorderTraverse()
}

func TestInorderTravers(t *testing.T) {
	root := testBST()
	InorderTraverse(root)
}

func TestGetInorderTraversal(t *testing.T) {
	testCase := []int{10, 12, 13, 20, 23, 24}
	root := testBST()
	nodes := GetInorderTraversal(root)
	for k, v := range testCase {
		if nodes[k] != v {
			t.Fatalf("got %d, want %d", nodes[k], v)
		}
	}
}

func TestIterativeInorderTraversal(t *testing.T) {
	testCase := []int{10, 12, 13, 20, 23, 24}
	root := testBST()
	nodes := GetInorderTraversal(root)
	for k, v := range testCase {
		if nodes[k] != v {
			t.Fatalf("got %d, want %d", nodes[k], v)
		}
	}
}

func TestLevelOrderTraversal(t *testing.T) {
	testCase := []int{13, 12, 23, 10, 20, 24}
	root := testBST()
	nodes := GetLevelOrderTraversal(root)
	for k, v := range testCase {
		if nodes[k] != v {
			t.Fatalf("got %d, want %d", nodes[k], v)
		}
	}
}

func TestMax(t *testing.T) {
	want := 24
	root := testBST()
	max := Max(root)
	if max != want {
		t.Fatalf("got max = %d, want max = %d", max, want)
	}
}

func TestRecursiveMax(t *testing.T) {
	root := testBST()
	root.Insert(67)
	want := 67
	max := RecursiveMax(root)
	if max != want {
		t.Fatalf("got max = %d, want max = %d", max, want)
	}
}

func TestMin(t *testing.T) {
	root := testBST()
	want := 2
	root.Insert(want)
	min := Min(root)

	if min != want {
		t.Fatalf("got min = %d, want min = %d", min, want)
	}
}

func TestRecursiveMin(t *testing.T) {
	root := testBST()
	want := 2
	root.Insert(want)
	min := RecursiveMin(root)
	if min != want {
		t.Fatalf("got min = %d want min = %d,", min, want)
	}
}

func TestHeight(t *testing.T) {
	root := testBST()
	want := 2
	h := Height(root)

	if h != want {
		t.Fatalf("got height = %d, want height = %d", h, want)
	}
}

//           13
//	        /  \
//        12    23
//       /     / \
//     10     20  24
//
//
func testBST() *Node {
	root := NewNode(13)

	root.left = NewNode(12)
	root.right = NewNode(23)

	root.left.left = NewNode(10)
	root.right.left = NewNode(20)
	root.right.right = NewNode(24)

	return root
}
