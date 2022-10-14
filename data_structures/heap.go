package datastructures

/* Heap is a binary tree where every child and grand child is smaller (MaxHeap)
or larger(MinHeap) than the current node.
This means:
- Every time a node is added. Tree must be adjusted
- Every time a node is deleted. Tree must be adjusted
Tree is always filled from left to right
Ex: Heap could be a priority queue
*/

/*
	MinHeap struct stores data of the binary tree but instead of using a tree is simpler to

store data in an array and keep tracking index and length to know parent/child relationships
MinHeap rule:
- head is the smallest of all tree nodes
*/
type MinHeap struct {
	Length int
	data   []int
}

func (h *MinHeap) Add(value int) {
	if h.Length == 0 {
		h.data = []int{}
	}
	h.data = append(h.data, value)
	h.heapifyUp(h.Length)
	h.Length += 1
}

func (h *MinHeap) Delete() int {
	if h.Length == 0 {
		return -1
	}
	h.Length -= 1
	value := h.data[0]
	if h.Length == 0 {
		h.data = []int{}
		return value
	}
	h.data[0] = h.data[h.Length]
	h.heapifyDown(0)
	return value
}

/*
func heapifyUp adjust tree from given idx in order to place node at given idx the right position
move node up until get to head (if is the smallest value) or its parent is lower than node
related to its current parent
*/
func (h *MinHeap) heapifyUp(idx int) {
	if idx == 0 {
		return
	}
	parent := h.getParent(idx)
	if h.data[parent] > h.data[idx] {
		// Swap values
		h.data[idx], h.data[parent] = h.data[parent], h.data[idx]
		h.heapifyUp(parent)
	}
}

/*
func heapifyDown adjust tree from given idx in order to place node at given idx the right position
move node down until get to last idx (if is the largest value) or its parent is grater than node
related to its current parent
*/
func (h *MinHeap) heapifyDown(idx int) {
	leftIdx := h.getLeftChild(idx)
	rightIdx := h.getRigthChild(idx)
	// Looks if get the latest idx on data or there are no more children
	if idx >= h.Length || leftIdx >= h.Length {
		return
	}

	// Look for smaller child to heapifyDown
	leftV := h.data[leftIdx]
	rightV := h.data[rightIdx]
	value := h.data[idx]
	if leftV > rightV && rightV < value {
		h.data[rightIdx], h.data[idx] = value, rightV
		h.heapifyDown(rightIdx)
	}
	if leftV < rightV && leftV < value {
		h.data[leftIdx], h.data[idx] = value, leftV
		h.heapifyDown(leftIdx)
	}
}

// func getParent returns index of the parent of the node placed in the given idx
func (h *MinHeap) getParent(idx int) int {
	return (idx - 1) / 2
}

// func getParent returns index of the left child of the node placed in the given idx
func (h *MinHeap) getLeftChild(idx int) int {
	return idx*2 + 1
}

// func getParent returns index of the right child of the node placed in the given idx
func (h *MinHeap) getRigthChild(idx int) int {
	return idx*2 + 2
}
