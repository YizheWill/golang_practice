package main

// This is the struct of the input root. Do not edit it.
type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

func BranchSums(root *BinaryTree) []int {
	result := []int{}
	sumTheBranches(root, 0, &result)
	return result
}

func sumTheBranches(node *BinaryTree, currSum int, sum *[]int) {
	if node == nil {
		return
	}
	currSum += node.Value
	if node.Left == nil && node.Right == nil {
		*sum = append(*sum, currSum)
	}
	sumTheBranches(node.Left, currSum, sum)
	sumTheBranches(node.Right, currSum, sum)
}