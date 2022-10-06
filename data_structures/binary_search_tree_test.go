package datastructures

import "testing"

func TestBinarySearchTreeFind(t *testing.T) {
	isPresent := DFSOnBST(&BTree, 45)
	if isPresent != true {
		t.Errorf("Binary search tree find fails. %v exists in the given binary tree", 45)
	}
	isPresent = DFSOnBST(&BTree, 7)
	if isPresent != true {
		t.Errorf("Binary search tree find fails. %v exists in the given binary tree", 7)
	}
	isPresent = DFSOnBST(&BTree, 69)
	if isPresent != false {
		t.Errorf("Binary search tree find fails. %v doesn't exist in the given binary tree", 69)
	}
}
