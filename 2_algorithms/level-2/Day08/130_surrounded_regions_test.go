package bfs_dfs

import (
	"fmt"
	"testing"
)

func print(input [][]byte) {
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			fmt.Printf("%c\t", input[i][j])
		}
		fmt.Println()
	}
}

func TestSolve(t *testing.T) {

	/*
		input:=[][]byte {
		{'O','O','O','O','X','X'},
		{'O','O','O','O','O','O'},
		{'O','X','O','X','O','O'},
		{'O','X','O','O','X','O'},
		{'O','X','O','X','O','O'},
		{'O','X','O','O','O','O'},
		}

	*/

	input := [][]byte{{'O', 'X', 'X', 'O', 'X'}, {'X', 'O', 'O', 'X', 'O'}, {'X', 'O', 'X', 'O', 'X'}, {'O', 'X', 'O', 'O', 'O'}, {'X', 'X', 'O', 'X', 'O'}}

	print(input)
	fmt.Println("----------------")

	solve(input)

	fmt.Printf("----------------\nResult:\n")
	print(input)

}
