package searchalgo

/*
An image is represented by an m x n integer grid image where image[i][j] represents the pixel value of the image.

You are also given three integers sr, sc, and color. You should perform a flood fill on the image starting from the pixel image[sr][sc].

To perform a flood fill, consider the starting pixel, plus any pixels connected 4-directionally to the starting pixel of the same color as the starting pixel, plus any pixels connected 4-directionally to those pixels (also with the same color), and so on. Replace the color of all of the aforementioned pixels with color.

Return the modified image after performing the flood fill.

*/

func FloodFill(image [][]int, sr int, sc int, color int) [][]int {

	if image == nil {
		return nil
	}

	rows := len(image)
	cols := len(image[0])

	result := make([][]int, rows)
	for i := range result {
		result[i] = make([]int, cols)
	}
	copy(result, image)

	fillColor(result, rows, cols, sr, sc, image[sr][sc], color)

	return result
}

func fillColor(result [][]int, rows, cols, rowIdx, colIdx, oldColor, newColor int) {

	if rowIdx < 0 || rowIdx >= rows || colIdx < 0 || colIdx >= cols || result[rowIdx][colIdx] != oldColor || result[rowIdx][colIdx] == newColor {
		return
	}

	result[rowIdx][colIdx] = newColor

	fillColor(result, rows, cols, rowIdx+1, colIdx, oldColor, newColor)
	fillColor(result, rows, cols, rowIdx-1, colIdx, oldColor, newColor)
	fillColor(result, rows, cols, rowIdx, colIdx+1, oldColor, newColor)
	fillColor(result, rows, cols, rowIdx, colIdx-1, oldColor, newColor)

}

func fillColor_recursion(result [][]int, rows, cols, rowIdx, colIdx, oldColor, newColor int) {

	if rowIdx < 0 || rowIdx >= rows || colIdx < 0 || colIdx >= cols || result[rowIdx][colIdx] != oldColor || result[rowIdx][colIdx] == newColor {
		return
	}

	result[rowIdx][colIdx] = newColor

	fillColor_recursion(result, rows, cols, rowIdx-1, colIdx, oldColor, newColor)
	fillColor_recursion(result, rows, cols, rowIdx, colIdx-1, oldColor, newColor)
	fillColor_recursion(result, rows, cols, rowIdx, colIdx+1, oldColor, newColor)
	fillColor_recursion(result, rows, cols, rowIdx+1, colIdx, oldColor, newColor)

}
