package main

type BST struct {
	Value int
	Left  *BST
	Right *BST
}

func dfsSearch(tree *BST, target int) bool {
	if tree.Value == target {
		return true
	} else if tree.Left == nil && tree.Right == nil {
		return false
	} else {
		return dfsSearch(tree.Left, target) || dfsSearch(tree.Right, target)
	}
}