package sorting

/*
	Bubble sort has a running time of O(n^2). Iterates over each number in

the array swaping positions from left to right (arr[n], arr[n+1] = arr[n+1], arr[n])
if the next number is lower than the current one. Until get all the array sorted
*/
func BubbleSort(arr []int) []int {
	n := len(arr)
	for j := n; j > 1; j-- {
		for i := 0; i < (n - 1); i++ {
			if i <= n-2 {
				if arr[i] > arr[i+1] {
					arr[i], arr[i+1] = arr[i+1], arr[i]
				}
			}
		}
	}
	return arr
}
