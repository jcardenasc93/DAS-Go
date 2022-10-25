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
	Data: [][]GraphEdge{
		{
			GraphEdge{
				To:     1,
				Weight: 3,
			},
			GraphEdge{
				To:     2,
				Weight: 1,
			},
		}, //0
		{
			GraphEdge{
				To:     0,
				Weight: 3,
			},
			GraphEdge{
				To:     4,
				Weight: 1,
			},
		}, //1
		{
			GraphEdge{
				To:     0,
				Weight: 1,
			},
			GraphEdge{
				To:     3,
				Weight: 7,
			},
		}, //2
		{
			GraphEdge{
				To:     2,
				Weight: 7,
			},
			GraphEdge{
				To:     6,
				Weight: 4,
			},
		}, //3
		{
			GraphEdge{
				To:     1,
				Weight: 1,
			},
			GraphEdge{
				To:     5,
				Weight: 2,
			},
		}, //4
		{
			GraphEdge{
				To:     6,
				Weight: 1,
			},
			GraphEdge{
				To:     4,
				Weight: 2,
			},
		}, //5
		{
			GraphEdge{
				To:     3,
				Weight: 1,
			},
			GraphEdge{
				To:     5,
				Weight: 1,
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
