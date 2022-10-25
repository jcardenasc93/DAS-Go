package challenges

import (
	"testing"

	datastructures "github.com/DSA-Go/data_structures"
	"github.com/DSA-Go/utils"
)

//      (1) --- (4) ---- (5)
//    /  |       |       /|
// (0)   | ------|------- |
//    \  |/      |        |
//      (2) --- (3) ---- (6)

var weightGraph = datastructures.WeightedAdjList{
	Data: [][]datastructures.GraphEdge{
		{
			datastructures.GraphEdge{
				To:     1,
				Weight: 3,
			},
			datastructures.GraphEdge{
				To:     2,
				Weight: 1,
			},
		}, // 0
		{
			datastructures.GraphEdge{
				To:     0,
				Weight: 3,
			},
			datastructures.GraphEdge{
				To:     2,
				Weight: 4,
			},
			datastructures.GraphEdge{
				To:     4,
				Weight: 1,
			},
		}, // 1
		{
			datastructures.GraphEdge{
				To:     1,
				Weight: 4,
			},
			datastructures.GraphEdge{
				To:     3,
				Weight: 7,
			},
			datastructures.GraphEdge{
				To:     0,
				Weight: 1,
			},
		}, // 2
		{
			datastructures.GraphEdge{
				To:     2,
				Weight: 7,
			},
			datastructures.GraphEdge{
				To:     4,
				Weight: 5,
			},
			datastructures.GraphEdge{
				To:     6,
				Weight: 1,
			},
		}, // 3
		{
			datastructures.GraphEdge{
				To:     1,
				Weight: 1,
			},
			datastructures.GraphEdge{
				To:     3,
				Weight: 5,
			},
			datastructures.GraphEdge{
				To:     5,
				Weight: 2,
			},
		}, // 4
		{
			datastructures.GraphEdge{
				To:     6,
				Weight: 1,
			},
			datastructures.GraphEdge{
				To:     4,
				Weight: 2,
			},
			datastructures.GraphEdge{
				To:     2,
				Weight: 18,
			},
		}, // 5
		{
			datastructures.GraphEdge{
				To:     3,
				Weight: 1,
			},
			datastructures.GraphEdge{
				To:     5,
				Weight: 1,
			},
		}, // 6
	},
}

func TestDjikstra(t *testing.T) {
	expected := []int{0, 1, 4, 5, 6}
	shortestPath := dijkstraList(0, 6, weightGraph)
	if utils.SlicesAreEqual(expected, shortestPath) == false {
		t.Errorf("Shortest path wrong.\nExpected: %v\n Got: %v", expected, shortestPath)
	}
}
