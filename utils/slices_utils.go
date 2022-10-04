package utils

func SlicesAreEqual(a []int, b []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
