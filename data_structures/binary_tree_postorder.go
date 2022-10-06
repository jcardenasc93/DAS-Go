package datastructures

// walkPost func is a helper to implements recursion for traversing binary tree
func walkPost(curr *BinaryTreeNode, path *[]int) []int {
	// Base case when there is no child
	if curr == nil {
		return *path
	}

	// Pre step
	// Recursion
	walkPost(curr.Left, path)
	walkPost(curr.Right, path)

	// Post step
	*path = append(*path, curr.Value)
	return *path
}

// preOrderSearch func traversal binary tree in pre-order way and prints
// each of the found Values. Post order traversal is depth first type
func postOrderSearch(head *BinaryTreeNode) []int {
	path := []int{}
	return walkPost(head, &path)
}
