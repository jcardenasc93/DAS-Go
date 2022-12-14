package datastructures

// walk func is a helper to implements recursion for traversing binary tree
func walk(curr *BinaryTreeNode, path *[]int) []int {
	// Base case when there is no child
	if curr == nil {
		return *path
	}

	// Pre step
	*path = append(*path, curr.Value)
	// Recursion
	walk(curr.Left, path)
	walk(curr.Right, path)

	// Post step
	return *path
}

// preOrderSearch func traversal binary tree in pre-order way and prints
// each of the found Values
func preOrderSearch(head *BinaryTreeNode) []int {
	path := []int{}
	return walk(head, &path)
}
