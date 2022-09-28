package datastructures

import (
	"errors"
)

type LinkListNode struct {
	value int
	prev  *LinkListNode
	next  *LinkListNode
}

type DoubleLinkList struct {
	lenght int
	head   *LinkListNode
	tail   *LinkListNode
}

func (l *DoubleLinkList) prepend(value int) {
	node := LinkListNode{value: value}
	if l.lenght == 0 {
		l.head = &node
		l.tail = &node
		l.lenght++
		return
	}
	node.next = l.head
	l.head.prev = &node
	l.head = &node

	if l.lenght == 1 {
		l.tail = &node
	}

	l.lenght++
}

func (l *DoubleLinkList) append(value int) {
	node := LinkListNode{value: value}
	if l.lenght == 0 {
		l.head = &node
		l.tail = &node
		l.lenght++
		return
	}
	node.prev = l.tail
	l.tail.next = &node
	l.tail = &node

	if l.lenght == 1 {
		l.head = node.prev
	}
	l.lenght++
}

func (l *DoubleLinkList) insertAt(idx int, value int) {
	if idx == l.lenght {
		l.append(value)
		return
	} else if idx == 0 {
		l.prepend(value)
		return
	}
	curr, _ := l.getAt(idx)
	node := LinkListNode{value: value}
	node.prev = curr.prev
	node.next = curr
	curr.prev = &node

	if node.prev != nil {
		node.prev.next = &node
	}

	l.lenght++
}

func (l *DoubleLinkList) remove(value int) (int, error) {
	if l.lenght == 0 {
		return 0, errors.New("List is empty")
	}
	// Search value
	curr := l.head
	for i := 0; i < l.lenght; i++ {
		if curr.value == value {
			break
		}
		curr = curr.next
	}

	if curr == nil {
		// Value not found
		return 0, errors.New("Value not found")
	}
	return l.removeNode(curr), nil
}

func (l *DoubleLinkList) removeAt(idx int) (int, error) {
	node, err := l.getAt(idx)
	if err != nil {
		return 0, err
	}
	return l.removeNode(node), err
}

func (l *DoubleLinkList) removeNode(node *LinkListNode) int {
	l.lenght--
	value := node.value

	if l.lenght == 0 {
		l.head = nil
		l.tail = nil
		return value
	}

	// Found in head position
	if node == l.head {
		node.next.prev = nil
		l.head = node.next
		node = nil
		return value
	}
	// Found in last position
	if node == l.tail {
		node.prev.next = nil
		l.tail = node.prev
		node = nil
		return value
	}

	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	return value
}

func (l *DoubleLinkList) getAt(idx int) (*LinkListNode, error) {
	if idx > l.lenght {
		return nil, errors.New("Index out of range")
	}
	curr := l.head
	for i := 0; curr != nil && i < idx; i++ {
		curr = curr.next
	}
	return curr, nil
}

func (l *DoubleLinkList) get(idx int) (int, error) {
	node, err := l.getAt(idx)
	return node.value, err
}
