package datastructures

// walkPost func is a helper to implements recursion for traversing binary tree
func walkPost(curr *BinaryTreeNode, path *[]int) []int {
	// Base case when there is no child
	if curr == nil {
		return *path
	}

	// Pre step
	// Recursion
	walkPost(curr.left, path)
	walkPost(curr.right, path)

	// Post step
	*path = append(*path, curr.value)
	return *path
}

// preOrderSearch func traversal binary tree in pre-order way and prints
// each of the found values. Post order traversal is depth first type
func postOrderSearch(head *BinaryTreeNode) []int {
	path := []int{}
	return walkPost(head, &path)
}
