package datastructures

/* Binary Search Trees (BTS) are binary trees with simple constraints
- All nodes at left side are less or equal than their parent
- All nodes at right side are greater than their parent */

// DFSOnBST search for a given value in a binary search tree
// implementing deep first search apporach. Returns true if value
// exists in tree
func DFSOnBST(head *BinaryTreeNode, value int) bool {
	if head == nil {
		return false
	}
	if value == head.Value {
		return true
	}
	if value < head.Value {
		return DFSOnBST(head.Left, value)
	}
	return DFSOnBST(head.Right, value)
}
