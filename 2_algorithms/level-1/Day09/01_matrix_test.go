package searchalgo

import (
	"fmt"
	"testing"
)

func print(inp [][]int) {

	for row := 0; row < len(inp); row++ {
		fmt.Println(inp[row])
	}

}

func TestUpdateMatrix(t *testing.T) {

	mat := [][]int{
		{1, 1, 0, 0, 1, 0, 0, 1, 1, 0},
		{1, 0, 0, 1, 0, 1, 1, 1, 1, 1},
		{1, 1, 1, 0, 0, 1, 1, 1, 1, 0},
		{0, 1, 1, 1, 0, 1, 1, 1, 1, 1},
		{0, 0, 1, 1, 1, 1, 1, 1, 1, 0},
		{1, 1, 1, 1, 1, 1, 0, 1, 1, 1},
		{0, 1, 1, 1, 1, 1, 1, 0, 0, 1},
		{1, 1, 1, 1, 1, 0, 0, 1, 1, 1},
		{0, 1, 0, 1, 1, 0, 1, 1, 1, 1},
		{1, 1, 1, 0, 1, 0, 1, 1, 1, 1},
	}

	result := UpdateMatrix(mat)

	fmt.Println("Input")
	print(mat)

	fmt.Println("Output")
	print(result)

}
