package datastructures

type qIntNode struct {
	value int
	next  *qIntNode
}

// Queues are data structures based on simple linked lists of FIFO type
type IntQueue struct {
	length int
	head   *qIntNode
	tail   *qIntNode
}

func (q *IntQueue) peek() int {
	return q.head.value
}

func (q *IntQueue) deque() (int, bool) {
	// First check if head exists and decrease 1 length
	if q.head == nil {
		var noValue int
		return noValue, false
	}
	q.length -= 1
	// Store current head value before re-link head node
	var head_value int = q.head.value
	q.head = q.head.next

	return head_value, true
}

func (q *IntQueue) enqueue(item int) {
	// Creates node to be inserted in queue
	node := qIntNode{
		value: item,
	}
	// No tail means empty queue, so set head & tail with same value
	if q.tail == nil {
		q.tail, q.head = &node, &node
		q.head.next = q.tail
	} else {
		// If there is a tail then set next with new node
		q.tail.next = &node
		// Then update tail with new node
		q.tail = &node
	}
	q.length += 1

}
