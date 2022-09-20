package search

import (
	"testing"
)

var unorderedArr = []uint8{1, 12, 24, 44, 3, 33}
var orderedArr = []int{1, 3, 12, 24, 33, 44}

var linearInput = []struct {
	arr      []uint8
	number   uint8
	expected int
}{
	{unorderedArr, 12, 1},
	{unorderedArr, 9, -1},
}

var binaryInput = []struct {
	arr      []int
	number   int8
	expected int
}{
	{orderedArr, 3, 1},
	{orderedArr, 24, 3},
	{orderedArr, 9, -1},
}

func TestLinearSearch(t *testing.T) {
	for _, testCase := range linearInput {
		index := LinearSearch(testCase.arr, testCase.number)
		if index != testCase.expected {
			t.Errorf("Found index should be %d not %d \n", testCase.expected, index)
		}
	}
}

func TestBinarySearch(t *testing.T) {
	for _, testCase := range binaryInput {
		index := BinarySearch(testCase.arr, int(testCase.number))
		if index != testCase.expected {
			t.Errorf("Found index should be %d not %d \n", testCase.expected, index)
		}
	}
}
