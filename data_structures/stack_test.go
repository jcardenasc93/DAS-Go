package datastructures

import "testing"

func TestIntStack(t *testing.T) {
	var intStack IntStack

	intStack.push(5)
	intStack.push(7)
	intStack.push(9)

	peekVal := intStack.peek()
	if peekVal != 9 {
		t.Errorf("Wrong peek process. Expected 9 got %v", peekVal)
	}

	if intStack.length != 3 {
		wrongLengthError(t, 3, intStack.length)
	}

	popItem, _ := intStack.pop()
	if popItem != 9 {
		t.Errorf("Wrong pop process. Expected 5 got %v", popItem)
	}
	if intStack.length != 2 {
		wrongLengthError(t, 2, intStack.length)
	}

	popItem, _ = intStack.pop()
	if popItem != 7 {
		t.Errorf("Wrong pop process. Expected 7 got %v", popItem)
	}
	if intStack.length != 1 {
		wrongLengthError(t, 1, intStack.length)
	}

	popItem, _ = intStack.pop()
	if popItem != 5 {
		t.Errorf("Wrong pop process. Expected 9 got %v", popItem)
	}
	if intStack.length != 0 {
		wrongLengthError(t, 0, intStack.length)
	}

	popItem, ok := intStack.pop()
	if ok {
		t.Errorf("Wrong pop process. No items in stack")
	}
	if intStack.length != 0 {
		wrongLengthError(t, 0, intStack.length)
	}

}
