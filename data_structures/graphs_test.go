package datastructures

import (
	"github.com/DSA-Go/utils"
	"testing"
)

//	 >(1)<--->(4) ---->(5)
//	/          |       /|
//
// (0)   ------|------- |
//
//	\   v      v        v
//	 >(2) --> (3) <----(6)

var weightGraph = WeightedAdjMatrix{
	data: [][]int{
		{0, 3, 1, 0, 0, 0, 0}, // 0
		{0, 0, 0, 0, 1, 0, 0},
		{0, 0, 7, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 5, 0, 2, 0},
		{0, 0, 18, 0, 0, 0, 1},
		{0, 0, 0, 1, 0, 0, 1},
	},
}

func TestBFSOnWeightGraph(t *testing.T) {
	expected := []int{
		0,
		1,
		4,
		5,
		6,
	}
	searchPath := BFSOnAdjancentMatrixGraph(weightGraph, 0, 6)
	if utils.SlicesAreEqual(expected, searchPath) == false {
		t.Errorf("Error on retrieved path.\nExpected:\n%v\nGot:\n%v", expected, searchPath)
	}

	// Not found case
	expected = []int{}
	searchPath = BFSOnAdjancentMatrixGraph(weightGraph, 6, 0)
	if utils.SlicesAreEqual(expected, searchPath) == false {
		t.Errorf("Error on retrieved path.\nExpected:\n%v\nGot:\n%v", expected, searchPath)
	}
}
