package challenges

/* LRU: Least Recently Used cache is an implementation of cache
that allows to keep track of the least used element from an user
add new elements to the cache and reorganize its uses once the
user asks for an element.
Ex: Let's assume the next cache
v0 <-> v1 <-> v2 <-> v3 <-> v4
At this point user asked for v2. So as the element exists
the cache will return v2 and reorganize uses on cache as follows:
v2 <-> v0 <-> v1 <-> v3 <-> v4
*/

// LRU sturct allows to handle LRU cache operations:
// This struct is build up using hashing map to get data
// directly using a key (O(1)) and a double linked list nodes
// as the values that keys points out

type Node[T any] struct {
	value T
	next  *Node[T]
	prev  *Node[T]
}

func createNode[V any](value V) Node[V] {
	return Node[V]{
		value: value,
	}
}

type LRU[K comparable, V any] struct {
	lenght        int
	capacity      int
	head          *Node[V]
	tail          *Node[V]
	lookup        map[K]*Node[V]
	reverseLookup map[*Node[V]]K
}

func (lru *LRU[K, V]) update(key K, value V) {
	// Check if key exists
	node, exists := lru.lookup[key]
	if exists == false {
		// If key doesn't exist then the node should be created
		newNode := createNode(value)
		// Update lenght and check capacity
		lru.lenght += 1
		lru.prepend(&newNode)
		lru.trimCache()
		lru.lookup[key] = &newNode
		lru.reverseLookup[&newNode] = key
	} else {
		// Update cache; moving the node to the front of the linked list
		lru.detach(node)
		lru.prepend(node)
		node.value = value
	}
}

func (lru *LRU[K, V]) deleteTail() {
	key := lru.reverseLookup[lru.tail]
	delete(lru.lookup, key)
	delete(lru.reverseLookup, lru.tail)
	lru.detach(lru.tail)
	lru.lenght -= 1
}

func (lru *LRU[K, V]) trimCache() {
	if lru.lenght > lru.capacity {
		lru.deleteTail()
	}
}

func (lru *LRU[K, V]) get(key K) (V, bool) {
	// Check the cache for key existance
	node, exists := lru.lookup[key]
	if exists == false {
		var empty V
		return empty, false
	}
	// If key exists then cache needs to be updated
	// first detach the node from the linked list
	lru.detach(node)
	// then put the node as the most recently used item (prepend operation in linked list)
	lru.prepend(node)
	return node.value, true
}

func (lru *LRU[K, V]) detach(node *Node[V]) {
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}

	// Edge cases
	if node == lru.head {
		lru.head = lru.head.next
	}
	if node == lru.tail {
		lru.tail = lru.tail.prev
	}

	node.next = nil
	node.prev = nil
}

func (lru *LRU[K, V]) prepend(node *Node[V]) {
	if lru.head == nil {
		lru.head, lru.tail = node, node
		return
	}
	node.next = lru.head
	lru.head.prev = node
	lru.head = node
}
