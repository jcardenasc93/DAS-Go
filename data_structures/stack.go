package datastructures

type sIntNode struct {
	value int
	prev  *sIntNode
}

/*
	Stacks are the opposite from Queues. Stacks are FILO type of data structures.

and implements push, pop & peek operations
*/
type IntStack struct {
	length int
	head   *sIntNode
}

func (q *IntStack) peek() int {
	if q.head != nil {
		return q.head.value
	}
	return 0
}

func (s *IntStack) pop() (int, bool) {
	var headValue int
	// Check if stack is empty
	s.length -= 1
	if s.length < 0 {
		// If is empty set correct length & clean head
		s.length = 0
		s.head = nil
		return headValue, false
	} else {
		// Store current head value before re-link head node
		headValue = s.head.value
		s.head = s.head.prev
	}
	return headValue, true
}

func (s *IntStack) push(item int) {
	// Creates node to be inserted in stack
	node := sIntNode{
		value: item,
	}
	if s.head == nil {
		// First item in stack
		s.head = &node
	} else {
		node.prev = s.head
		s.head = &node
	}
	s.length += 1

}
