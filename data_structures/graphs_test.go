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

//	  (1) --- (4) ---- (5)
//	/                   |
//
// (0)  --------------- |
//
//	\                   |
//	  (2) --- (3) ---- (6)
var weigthGraphList = WeightedAdjList{
	data: [][]GraphEdge{
		{
			GraphEdge{
				to:     1,
				weight: 3,
			},
			GraphEdge{
				to:     2,
				weight: 1,
			},
		}, //0
		{
			GraphEdge{
				to:     0,
				weight: 3,
			},
			GraphEdge{
				to:     4,
				weight: 1,
			},
		}, //1
		{
			GraphEdge{
				to:     0,
				weight: 1,
			},
			GraphEdge{
				to:     3,
				weight: 7,
			},
		}, //2
		{
			GraphEdge{
				to:     2,
				weight: 7,
			},
			GraphEdge{
				to:     6,
				weight: 4,
			},
		}, //3
		{
			GraphEdge{
				to:     1,
				weight: 1,
			},
			GraphEdge{
				to:     5,
				weight: 2,
			},
		}, //4
		{
			GraphEdge{
				to:     6,
				weight: 1,
			},
			GraphEdge{
				to:     4,
				weight: 2,
			},
		}, //5
		{
			GraphEdge{
				to:     3,
				weight: 1,
			},
			GraphEdge{
				to:     5,
				weight: 1,
			},
		}, //6
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

func TestDFSOnWeightGraph(t *testing.T) {
	expected := []int{
		0,
		1,
		4,
		5,
		6,
	}
	searchPath := DFSOnAdjancentMatrixGraph(weigthGraphList, 0, 6)
	if utils.SlicesAreEqual(expected, searchPath) == false {
		t.Errorf("Error on retrieved path.\nExpected:\n%v\nGot:\n%v", expected, searchPath)
	}

	// Not found case
	expected = []int{}
	searchPath = DFSOnAdjancentMatrixGraph(weigthGraphList, 6, 0)
	if utils.SlicesAreEqual(expected, searchPath) == false {
		t.Errorf("Error on retrieved path.\nExpected:\n%v\nGot:\n%v", expected, searchPath)
	}
}
