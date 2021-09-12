package main

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func NodeDepths(root *BinaryTree) int {
	return calcDepth(root, 0)
}

func calcDepth(root *BinaryTree, sum int) int {
	if root == nil {
		return 0
	} else {
		return sum + calcDepth(root.Left, sum + 1) + calcDepth(root.Right, sum + 1)
	}
}
