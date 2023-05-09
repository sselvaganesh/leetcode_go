package graph

/*

Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return the number of islands.
An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

Example 1:
Input: grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
Output: 1


Example 2:
Input: grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
Output: 3

*/


func NumIslands(grid [][]byte) int {

    if grid==nil {
        return 0
    }

    rows:=len(grid)
    cols:=len(grid[0])

    newGrid:=make([][]byte, rows)
    for i:= range newGrid {
        newGrid[i]=make([]byte, cols)
    }
    copy(newGrid, grid)

    islands:=0
    for row:=0; row<rows; row++ {

        for col:=0; col<cols; col++ {

            if newGrid[row][col]=='1' {
                islands++
                sink(newGrid, rows, cols, row, col)
            }

        }

    }

    return islands
}

func sink(newGrid [][]byte, rows, cols, row, col int) {

    if row<0 || row>=rows || col<0 || col>=cols || newGrid[row][col]=='0' {
        return
    }

    newGrid[row][col]='0'

    sink(newGrid, rows, cols, row+1, col)
    sink(newGrid, rows, cols, row-1, col)
    sink(newGrid, rows, cols, row, col+1)
    sink(newGrid, rows, cols, row, col-1)

    return
}



