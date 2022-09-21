package datastructures

import "testing"

func wrongLengthError(t *testing.T, expected int, got int) {
	t.Errorf("Int queue length is wrong. Expected %v got %v", expected, got)

}

func TestIntQueue(t *testing.T) {
	var intQueue IntQueue

	intQueue.enqueue(5)
	intQueue.enqueue(7)
	intQueue.enqueue(9)

	if intQueue.length != 3 {
		wrongLengthError(t, 3, intQueue.length)
	}

	dequeItem, _ := intQueue.deque()
	if dequeItem != 5 {
		t.Errorf("Wrong deque process. Expected 5 got %v", dequeItem)
	}
	if intQueue.length != 2 {
		wrongLengthError(t, 2, intQueue.length)
	}

	dequeItem, _ = intQueue.deque()
	if dequeItem != 7 {
		t.Errorf("Wrong deque process. Expected 7 got %v", dequeItem)
	}
	if intQueue.length != 1 {
		wrongLengthError(t, 1, intQueue.length)
	}

	dequeItem, _ = intQueue.deque()
	if dequeItem != 9 {
		t.Errorf("Wrong deque process. Expected 9 got %v", dequeItem)
	}
	if intQueue.length != 0 {
		wrongLengthError(t, 0, intQueue.length)
	}

}
