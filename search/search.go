package search

// LinearSearch search for an specific value in a given array. Returns index of the value in the array or -1 if not found
func LinearSearch(arr []uint8, number uint8) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == number {
			return i
		}
	}
	return -1
}

// BinarySearch binary search implementation. Always assumes that input arr is a sorted array
func BinarySearch(arr []int, number int) int {
	response := -1
	low := 0
	high := len(arr)
	for low < high {
		middle := (low + (high - low)) / 2
		value := arr[middle]
		if value == number {
			response = middle
			break
		} else if value > number {
			high = middle - 1
		} else {
			low = middle + 1
		}
	}
	return response
}
