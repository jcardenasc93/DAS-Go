package challenges

import ds "github.com/DSA-Go/data_structures"

// Given two binary trees compare them to check if both are equal on shape & values

// func CompareBTrees returns true if both binary trees are equal on shape & values
func CompareBTrees(bA *ds.BinaryTreeNode, bB *ds.BinaryTreeNode) bool {
	if bA == nil && bB == nil {
		return true
	}
	if bA == nil || bB == nil {
		return false
	}

	if bA.Value != bB.Value {
		return false
	}

	return CompareBTrees(bA.Right, bB.Right) && CompareBTrees(bA.Left, bB.Left)
}
