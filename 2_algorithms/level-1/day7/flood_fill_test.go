package searchalgo

import (
	"testing"
)

func TestFloodFill(t *testing.T) {

	image := [][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}
	sr := 1
	sc := 1
	color := 2

	result := FloodFill(image, sr, sc, color)
	t.Logf("Result: %v", result)

}
