package datastructures

// walkIn func is a helper to implements recursion for traversing binary tree
func walkIn(curr *BinaryTreeNode, path *[]int) []int {
	// Base case when there is no child
	if curr == nil {
		return *path
	}

	// Pre step
	// Recursion
	walkIn(curr.Left, path)
	*path = append(*path, curr.Value)
	walkIn(curr.Right, path)

	// Post step
	return *path
}

// preOrderSearch func traversal binary tree in pre-order way and prints
// each of the found Values. In order traversal is depth first type
func inOrderSearch(head *BinaryTreeNode) []int {
	path := []int{}
	return walkIn(head, &path)
}
