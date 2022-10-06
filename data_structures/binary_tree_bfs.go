package datastructures

// breadthFirstSearch func search for a number in a binary tree using
// breadth first search approach and returns true if exists
func breadthFirstSearch(head BinaryTreeNode, number int) bool {
	binaryQueue := []BinaryTreeNode{head}
	for len(binaryQueue) > 0 {
		curr := ShiftBinaryQueue(&binaryQueue)
		if curr.Value == number {
			return true
		}
		if curr.Left != nil {
			binaryQueue = append(binaryQueue, *curr.Left)
		}
		if curr.Right != nil {
			binaryQueue = append(binaryQueue, *curr.Right)
		}
	}
	return false
}
