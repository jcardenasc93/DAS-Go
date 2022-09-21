package sorting

import "testing"

var arr = []int{9, 3, 7, 4, 69, 420, 42}
var sorted = []int{3, 4, 7, 9, 42, 69, 420}
var arr1 = []int{122, 33, 7, 4, 169, 20, 42, 5, 2}
var sorted1 = []int{2, 4, 5, 7, 20, 33, 42, 122, 169}

var inputs = []struct {
	arr    []int
	sorted []int
}{
	{arr, sorted},
	{arr1, sorted1},
}

func slicesAreEqual(a []int, b []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestBubbleSort(t *testing.T) {
	for _, input := range inputs {
		arr = BubbleSort(input.arr)
		if slicesAreEqual(arr, input.sorted) == false {
			t.Errorf("Array is not sorted!\n%v is not sorted. Should be like\n%v", arr, sorted)

		}
	}

}
