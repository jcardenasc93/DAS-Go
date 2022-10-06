package challenges

import (
	"testing"

	datastructures "github.com/DSA-Go/data_structures"
)

func TestBTreesComparator(t *testing.T) {
	comp := CompareBTrees(&datastructures.BTree, &datastructures.BTree)
	if comp != true {
		t.Errorf("Error comparing binary trees. They are equal")
	}

	comp = CompareBTrees(&datastructures.BTree, &datastructures.BTree2)
	if comp != false {
		t.Errorf("Error comparing binary trees. They are not equal")
	}
}
