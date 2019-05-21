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
