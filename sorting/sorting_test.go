package sorting

import (
	"testing"

	"github.com/DSA-Go/utils"
)

var result = []int{9, 3, 7, 4, 69, 420, 42}
var sorted = []int{3, 4, 7, 9, 42, 69, 420}
var arr1 = []int{122, 33, 7, 4, 169, 20, 42, 5, 2}
var sorted1 = []int{2, 4, 5, 7, 20, 33, 42, 122, 169}

var inputs = []struct {
	arr    []int
	sorted []int
}{
	{result, sorted},
	{arr1, sorted1},
}

func TestBubbleSort(t *testing.T) {
	for _, input := range inputs {
		result = BubbleSort(input.arr)
		if utils.SlicesAreEqual(result, input.sorted) == false {
			t.Errorf("Array is not sorted!\n%v is not sorted. Should be like\n%v", result, input.sorted)

		}
	}
}

func TestQuickSort(t *testing.T) {
	for _, input := range inputs {
		result = QuickSort(&input.arr)
		if utils.SlicesAreEqual(result, input.sorted) == false {
			t.Errorf("Array is not sorted!\n%v is not sorted. Should be like\n%v", result, input.sorted)

		}
	}
}
