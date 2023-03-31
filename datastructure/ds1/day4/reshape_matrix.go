package array

/*
In MATLAB, there is a handy function called reshape which can reshape an m x n matrix into a new one with a different size r x c keeping its original data.
You are given an m x n matrix mat and two integers r and c representing the number of rows and the number of columns of the wanted reshaped matrix.
The reshaped matrix should be filled with all the elements of the original matrix in the same row-traversing order as they were.
If the reshape operation with given parameters is possible and legal, output the new reshaped matrix; Otherwise, output the original matrix.

Example 1:
Input: mat = [[1,2],[3,4]], r = 1, c = 4
Output: [[1,2,3,4]]

Example 2:
Input: mat = [[1,2],[3,4]], r = 2, c = 4
Output: [[1,2],[3,4]]
*/

func MatrixReshape(mat [][]int, r int, c int) [][]int {

	matRows := len(mat)
	if matRows == 0 {
		return nil
	}
	matCols := len(mat[0])

	if (matRows * matCols) != (r * c) {
		return mat
	}

	result := make([][]int, r)
	for i := range result {
		result[i] = make([]int, c)
	}

	row, col := 0, 0
	for matRow := 0; matRow < matRows; matRow++ {

		for matCol := 0; matCol < matCols; matCol++ {

			result[row][col] = mat[matRow][matCol]
			col++

			if col == c {
				row++
				col = 0
			}

		}

	}

	return result

}
