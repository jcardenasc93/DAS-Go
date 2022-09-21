package challenges

import "math"

/* Given two crystal balls that will break if dropped from high enough
distance, determine the exact spot in which it will break in the
most optimized way */

func TwoCrystalBalls(breaks []bool) int {
	jump := int(math.Floor(math.Sqrt(float64(len(breaks)))))
	i := jump
	for ; i < len(breaks); i += (jump) {
		if breaks[i] == true {
			// One of the balls is broken
			break
		}
	}
	i -= jump

	for j := 0; j < jump && i < len(breaks)-1; j++ {
		// Search straightforward on the last jump looking for the first broken ocurrency
		i += 1
		if breaks[i] {
			return i
		}
	}
	return -1
}
