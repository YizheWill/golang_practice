package main

// Do not edit the class below except
// for the depthFirstSearch method.
// Feel free to add new properties
// and methods to the class.
type Node struct {
	Name     string
	Children []*Node
}

func (n *Node) DepthFirstSearch(array []string) []string {
	// Write your code here.
	buildArray(n, &array)
	return array
}

func buildArray(n *Node, array *[]string) {
	if n == nil {
		return
	}
	*array = append(*array, n.Name)
	for _, node := range n.Children {
		buildArray(node, array)
	}
}