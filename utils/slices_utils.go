package utils

func SlicesAreEqual(a []int, b []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func FillBooleanSlice(s *[]bool, length int, value bool) {
	for i := 0; i < length; i++ {
		(*s)[i] = value
	}
}

func FillIntegerSlice(s *[]int, length int, value int) {
	for i := 0; i < length; i++ {
		(*s)[i] = value
	}
}
