package main

type BST struct {
	Value int
	Left  *BST
	Right *BST
}

func bfsSearch(tree *BST, target int) bool {
	nodes := []*BST{tree}
	var top *BST
	for len(nodes) > 0 {
		top, nodes = nodes[0], nodes[1:]
		if top.Value == target {
			return true
		} else if top == nil {
			continue
		}
		nodes = append(nodes, top.Left, top.Right)
	}
	return false
}
