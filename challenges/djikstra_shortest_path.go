package challenges

import (
	"math"

	datastructures "github.com/DSA-Go/data_structures"
	"github.com/DSA-Go/utils"
)

/* Djikstra's shortest path algorithm
is based on the principle to calc the shortest (lowest weight) to
reach a node from a source for every single node and traverse the
previous array to get the full shortest path */

var infinity = math.MaxInt

// func dijkstraList returns the list of nodes that be part of
// the shortest path from a given source to reach a given target
func dijkstraList(source int, target int, graph datastructures.WeightedAdjList) []int {
	seen := make([]bool, len(graph.Data), len(graph.Data)) // This slice saves seen nodes
	utils.FillBooleanSlice(&seen, len(seen), false)
	distances := make([]int, len(graph.Data), len(graph.Data)) // Keep track for distances between nodes
	utils.FillIntegerSlice(&distances, len(distances), infinity)
	prev := make([]int, len(graph.Data), len(graph.Data)) // Keep track for distances between nodes
	utils.FillIntegerSlice(&prev, len(prev), -1)

	distances[source] = 0 // Distance from source to source is zero

	for hasUnvisited(&seen, &distances) {
		lowest := getLowestUnvisited(&seen, &distances)
		seen[lowest] = true
		lowestEdges := graph.Data[lowest]
		for _, edge := range lowestEdges {
			if seen[edge.To] == true {
				continue
			}
			dist := distances[lowest] + edge.Weight // Calcs current distance
			if dist < distances[edge.To] {
				distances[edge.To] = dist // Update distance for the node if is shorter than current one
				prev[edge.To] = lowest    // Update the comming from node when a shorter distance is found
			}
		}
	}
	// At this point we got all shortest distances between nodes
	out := []int{}
	curr := target
	if prev[curr] == -1 {
		return out
	}
	for prev[curr] != -1 {
		out = append(out, curr)
		curr = prev[curr]
	}
	out = append(out, source) // Adds source
	// Reverse out in order to get path
	path := make([]int, len(out), len(out))
	for i := 0; i < len(path); i++ {
		path[i] = out[len(out)-i-1]
	}
	return path
}

// func hasUnvisited returns true if there at least one unseen index which distances is
// lowest than inf
func hasUnvisited(seen *[]bool, distances *[]int) bool {
	for i, s := range *seen {
		if s == false && (*distances)[i] <= infinity {
			return true
		}
	}
	return false
}

// func getLowestUnvisited returns the index of the lowest unvisited
func getLowestUnvisited(seen *[]bool, distances *[]int) int {
	idx := -1
	lowestDistance := infinity
	for i, s := range *seen {
		if s == true {
			continue
		}
		if (*distances)[i] < lowestDistance {
			lowestDistance = (*&*distances)[i]
			idx = i
		}
	}
	return idx
}
