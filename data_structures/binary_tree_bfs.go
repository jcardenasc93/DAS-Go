package datastructures

// breadthFirstSearch func search for a number in a binary tree using
// breadth first search approach and returns true if exists
func breadthFirstSearch(head BinaryTreeNode, number int) bool {
	binaryQueue := []BinaryTreeNode{head}
	for len(binaryQueue) > 0 {
		curr := ShiftBinaryQueue(&binaryQueue)
		if curr.value == number {
			return true
		}
		if curr.left != nil {
			binaryQueue = append(binaryQueue, *curr.left)
		}
		if curr.right != nil {
			binaryQueue = append(binaryQueue, *curr.right)
		}
	}
	return false
}
