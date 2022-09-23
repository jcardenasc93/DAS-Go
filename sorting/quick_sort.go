package sorting

/*
	Quick sort is based on divide & conquer premise

Pick a random pivot and divide array from it organizing all items less from the
chosen pivot to the left side and greater ones on right side

Repeat this process until sub-array has lenght = 1
*/

// partition func performs the weak sorting based on pivot. Must return pivot index
func partition(arr *[]int, low int, high int) int {
	pivot := (*arr)[high]
	index := low - 1
	for i := 0; i < high; i++ {
		if (*arr)[i] <= pivot {
			index++
			(*arr)[i], (*arr)[index] = (*arr)[index], (*arr)[i]

		}
	}
	index++
	(*arr)[index], (*arr)[high] = pivot, (*arr)[index]
	return index
}

// qs func is recursive allowing to perform sorting on each partition
func qs(arr *[]int, low int, high int) {
	if low >= high {
		// array is sorted
		return
	}
	pivotIdx := partition(arr, low, high)
	qs(arr, low, pivotIdx-1)  // -> right side
	qs(arr, pivotIdx+1, high) // -> left side
}

// QuickSort is quicksort algo entry point
func QuickSort(arr *[]int) []int {
	qs(arr, 0, 1)
	return *arr
}
