package datastructures

// walkIn func is a helper to implements recursion for traversing binary tree
func walkIn(curr *BinaryTreeNode, path *[]int) []int {
	// Base case when there is no child
	if curr == nil {
		return *path
	}

	// Pre step
	// Recursion
	walkIn(curr.left, path)
	*path = append(*path, curr.value)
	walkIn(curr.right, path)

	// Post step
	return *path
}

// preOrderSearch func traversal binary tree in pre-order way and prints
// each of the found values. In order traversal is depth first type
func inOrderSearch(head *BinaryTreeNode) []int {
	path := []int{}
	return walkIn(head, &path)
}
