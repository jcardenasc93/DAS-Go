package datastructures

import "testing"

var list = DoubleLinkList{}

func getFail(t *testing.T, value int, expected int) {
	t.Errorf("Get fails. Expected %v but got %v", expected, value)
}

func removeFail(t *testing.T, value int, expected int) {
	t.Errorf("Remove fails. Expected %v but got %v", expected, value)
}

func TestDoubleLinkList(t *testing.T) {
	list.append(5)
	list.append(7)
	list.append(9)

	getValue, err := list.get(2)
	if getValue != 9 || err != nil {
		getFail(t, getValue, 9)
	}
	valueRemove, err := list.removeAt(1)
	if valueRemove != 7 || err != nil {
		removeFail(t, valueRemove, 7)
	}

	if list.lenght != 2 {
		t.Errorf("Length is incorrect. Expected %v but got %v", list.lenght, 2)
	}

	list.append(11)
	valueRemove, err = list.removeAt(1)
	if valueRemove != 9 || err != nil {
		removeFail(t, valueRemove, 9)
	}
	valueRemove, err = list.removeAt(9)
	if err == nil {
		t.Errorf("Expected index error")
	}
	valueRemove, err = list.removeAt(0)
	if valueRemove != 5 || err != nil {
		removeFail(t, valueRemove, 5)
	}
	valueRemove, err = list.removeAt(0)
	if valueRemove != 11 || err != nil {
		removeFail(t, valueRemove, 11)
	}

	if list.lenght != 0 {
		t.Errorf("Length is incorrect. Expected %v but got %v", list.lenght, 0)
	}

	list.prepend(5)
	list.prepend(7)
	list.prepend(9)

	getValue, err = list.get(2)
	if getValue != 5 || err != nil {
		getFail(t, getValue, 5)
	}
	getValue, err = list.get(0)
	if getValue != 9 || err != nil {
		getFail(t, getValue, 9)
	}
	valueRemove, err = list.remove(9)
	if valueRemove != 9 || err != nil {
		removeFail(t, valueRemove, 9)
	}
	if list.lenght != 2 {
		t.Errorf("Length is incorrect. Expected %v but got %v", list.lenght, 2)
	}
	getValue, err = list.get(0)
	if getValue != 7 || err != nil {
		getFail(t, getValue, 7)
	}
}
