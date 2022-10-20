package datastructures

import (
	"github.com/DSA-Go/utils"
)

/* Graphs are basically trees without any special rule
about hierarchy on nodes position or something like that. Graphs
could be represented by two ways: Adjacent list or Adjacent matrix.

Grapsh could be weighted so this means that paths has a weight in an
specific direction.
*/

type WeightedAdjMatrix struct {
	data [][]int
}

// func BFSOnAdjancentMatrixGraph implements the logic to search a given value
// in a graph represented as adjacent matrix using branch first search approach
// and returns the path took to find the value
func BFSOnAdjancentMatrixGraph(graph WeightedAdjMatrix, startPoint int, value int) []int {
	path := []int{}
	seen := make([]bool, len(graph.data), len(graph.data))    // This slice saves seen nodes
	previous := make([]int, len(graph.data), len(graph.data)) // This slice allows to looking backwards in path traversed
	utils.FillBooleanSlice(&seen, len(graph.data), false)
	utils.FillIntegerSlice(&previous, len(graph.data), -1)

	// Initial setup to search
	seen[startPoint] = true
	queue := []int{startPoint}
	for ok := true; ok; ok = len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current == value {
			break
		}
		adjs := graph.data[current]
		for i := 0; i < len(adjs); i++ {
			if adjs[i] == 0 {
				// Node with no connections
				continue
			}
			if seen[i] == true {
				// Avoid walk twice on the same node
				continue
			}
			seen[i] = true
			previous[i] = current
			queue = append(queue, i)
		}
	}
	if previous[value] == -1 {
		return []int{}
	}

	// Built path with backwards traversed of previous slice
	current := value
	out := []int{}
	for previous[current] != -1 {
		out = append(out, current)
		current = previous[current]
	}
	out = append(out, startPoint)
	// Reverse out to get the path
	for i := len(out) - 1; i >= 0; i-- {
		path = append(path, out[i])
	}
	return path
}
