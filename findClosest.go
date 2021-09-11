package main

type BST struct {
	Value int
	Left  *BST
	Right *BST
}

func (tree *BST) FindClosestValue(target int) int {
	return tree.findClosestValue(target, tree.Value)
}

func (tree *BST) findTarget(target int) bool {
	if tree.Value == target {
		return true
	} else if target < tree.Value && tree.Left != nil {
		return tree.Left.findTarget(target)
	} else if target > tree.Value && tree.Right != nil {
		return tree.Right.findTarget(target)
	}
	return false
}

func (tree *BST) findClosestValue(target, closest int) int {
	if absdiff(target, closest) > absdiff(target, tree.Value) {
		closest = tree.Value
	}
	if target < tree.Value && tree.Left != nil {
		return tree.Left.findClosestValue(target, closest)
	} else if target > tree.Value && tree.Right != nil {
		return tree.Right.findClosestValue(target, closest)
	}
	return closest
}

func absdiff(a, b int) int {
	switch a < b {
	case true:
		return b - a
	default:
		return a - b
	}
}
