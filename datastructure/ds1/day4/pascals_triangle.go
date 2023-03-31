package array

/*
Given an integer numRows, return the first numRows of Pascal's triangle.
In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:

Example 1:
Input: numRows = 5
Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]

Example 2:
Input: numRows = 1
Output: [[1]]
*/

func Generate(numRows int) [][]int {

	if numRows == 0 {
		return nil
	}

	var result [][]int
	for i := 1; i <= numRows; i++ {

		eachRow := make([]int, i)
		eachRow[0] = 1
		eachRow[i-1] = 1

		if i <= 2 {
			result = append(result, eachRow)
			continue
		}

		prevRow := i - 2
		for idx := 1; idx < i-1; idx++ {
			eachRow[idx] = result[prevRow][idx-1] + result[prevRow][idx]
		}

		result = append(result, eachRow)
	}

	return result

}
