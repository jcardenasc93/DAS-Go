package datastructures

import (
	"testing"

	"github.com/DSA-Go/utils"
)

var preOrderExpected = []int{
	20,
	10,
	5,
	7,
	15,
	50,
	30,
	29,
	45,
	100,
}

var inOrderExpected = []int{
	5,
	7,
	10,
	15,
	20,
	29,
	30,
	45,
	50,
	100,
}

var postOrderExpected = []int{
	7,
	5,
	15,
	10,
	29,
	45,
	30,
	100,
	50,
	20,
}

var btfsInputTrue = []int{
	45,
	7,
	100,
}

var btfsInputFalse = []int{
	300,
	277,
	0,
}

func TestPreOrderBinaryTree(t *testing.T) {
	result := preOrderSearch(&BTree)
	if utils.SlicesAreEqual(result, preOrderExpected) == false {
		t.Errorf("Pre order traversal fails.\nExpected: %v\n Got: %v", preOrderExpected, result)
	}
}

func TestInOrderBinaryTree(t *testing.T) {
	result := inOrderSearch(&BTree)
	if utils.SlicesAreEqual(result, inOrderExpected) == false {
		t.Errorf("Pre order traversal fails.\nExpected: %v\n Got: %v", inOrderExpected, result)
	}
}

func TestPostOrderBinaryTree(t *testing.T) {
	result := postOrderSearch(&BTree)
	if utils.SlicesAreEqual(result, postOrderExpected) == false {
		t.Errorf("Pre order traversal fails.\nExpected: %v\n Got: %v", postOrderExpected, result)
	}
}

func TestBTFSIn(t *testing.T) {
	for _, number := range btfsInputTrue {
		if breadthFirstSearch(BTree, number) == false {
			t.Errorf("Breadth first search fails. %v is present in binary tree", number)
		}
	}
}

func TestBTFSNotIn(t *testing.T) {
	for _, number := range btfsInputFalse {
		if breadthFirstSearch(BTree, number) == true {
			t.Errorf("Breadth first search fails. %v is present in binary tree", number)
		}
	}
}
