package tree

import "fmt"

type Tree struct {
	V     int
	Left  *Tree
	Right *Tree
}

// Simple tree creation and doing in order, pre order and post order
func main() {
	var root *Tree

	// Manually building tree
	root = &Tree{V: 1}
	root.Left = &Tree{V: 2}
	root.Right = &Tree{V: 3}

	root.Left.Left = &Tree{V: 4}
	root.Left.Right = &Tree{V: 5}

	root.Right.Left = &Tree{V: 6}
	root.Right.Right = &Tree{V: 7}

	fmt.Println("Inorder Traversal:")
	InOrder(root)

	fmt.Println("\nPreorder Traversal:")
	PreOrder(root)

	fmt.Println("\nPostorder Traversal:")
	PostOrder(root)
	fmt.Println("")
}

func InOrder(root *Tree) {
	if root == nil {
		return
	}
	InOrder(root.Left)
	fmt.Print(root.V, " ")
	InOrder(root.Right)
}

func PreOrder(root *Tree) {
	if root == nil {
		return
	}
	fmt.Print(root.V, " ")
	PreOrder(root.Left)
	PreOrder(root.Right)
}

func PostOrder(root *Tree) {
	if root == nil {
		return
	}
	PostOrder(root.Left)
	PostOrder(root.Right)
	fmt.Print(root.V, " ")
}
