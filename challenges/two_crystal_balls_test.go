package challenges

import (
	"math/rand"
	"testing"
)

func initializeValues(withBreak bool) ([]bool, int) {
	const arrSize = 5
	var idx int = 0 + rand.Intn(arrSize-0)

	arr := make([]bool, arrSize)
	for i := range arr {
		arr[i] = false
	}
	if withBreak {
		for i := idx; i < len(arr); i++ {
			arr[i] = true
		}
	} else {
		idx = -1
	}
	return arr, idx
}

type crystalBallsInput struct {
	arr      []bool
	expected int
}

func createInputs() []crystalBallsInput {
	var inputs []crystalBallsInput
	arr, idx := initializeValues(true)
	inputs = append(inputs, crystalBallsInput{arr: arr, expected: idx})
	arr, idx = initializeValues(false)
	inputs = append(inputs, crystalBallsInput{arr: arr, expected: idx})
	return inputs
}

func TestTwoCrystalBalls(t *testing.T) {
	inputs := createInputs()
	for _, input := range inputs {
		result := TwoCrystalBalls(input.arr)
		if result != input.expected {
			t.Errorf("Fails: %d is not %d", result, input.expected)
		}
	}
}
