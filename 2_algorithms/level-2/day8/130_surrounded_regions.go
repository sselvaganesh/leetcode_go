package bfs_dfs

/*
130. Surrounded Regions
Medium
7K
1.5K
Companies

Given an m x n matrix board containing 'X' and 'O', capture all regions that are 4-directionally surrounded by 'X'.

A region is captured by flipping all 'O's into 'X's in that surrounded region.

Example 1:
Input: board = [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]
Output: [["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]
Explanation: Notice that an 'O' should not be flipped if:
- It is on the border, or
- It is adjacent to an 'O' that should not be flipped.
The bottom 'O' is on the border, so it is not flipped.
The other three 'O' form a surrounded region, so they are flipped.

Example 2:
Input: board = [["X"]]
Output: [["X"]]


Constraints:
m == board.length
n == board[i].length
1 <= m, n <= 200
board[i][j] is 'X' or 'O'.

[
["O","X","X","O","X"],
["X","O","O","X","O"],
["X","O","X","O","X"],
["O","X","O","O","O"],
["X","X","O","X","O"]
]

[
	["O","O","O","O","X","X"],
	["O","O","O","O","O","O"],
	["O","X","O","X","O","O"],
	["O","X","O","O","X","O"],
	["O","X","O","X","O","O"],
	["O","X","O","O","O","O"]
	
	]

*/

import (
	//"fmt"
)



func solve(board [][]byte)  {

    //dfs_solution(board)
    solution2(board)

}

func solution2(board [][]byte) {

    if board==nil {
        return
    }
    rows, cols := len(board), len(board[0])

    for row:=0; row<rows; row++ {
        for col:=0; col<cols; col++ {
            if (row==0 || col==0 || row==rows-1 || col==cols-1) && board[row][col]=='O' {
                markT(board, rows, cols, row, col)
            }
        }
    }

    findReplace(board, rows, cols, 'O', 'X')
    findReplace(board, rows, cols, 'T', 'O')

}

func markT(board [][]byte, rows, cols, row, col int) {

    if row<0 || col<0 || row>=rows || col>=cols || board[row][col]!='O' {
        return
    }
    board[row][col]='T'
    markT(board, rows, cols, row, col-1)
    markT(board, rows, cols, row, col+1)
    markT(board, rows, cols, row-1, col)
    markT(board, rows, cols, row+1, col)
}

func findReplace(board [][]byte, rows, cols int, find, replace byte) {

    for row:=0; row<rows; row++ {
        for col:=0; col<cols; col++ {
            if board[row][col]==find {
                board[row][col]=replace
            }
        }
    }

}



func dfs_solution(board [][]byte) {

    if board==nil{
        return
    }

    rows, cols := len(board), len(board[0])
    for row:=1; row<rows-1; row++ {
        for col:=1; col<cols-1; col++ {
            if board[row][col]=='O' {                
                if isRegion(board, rows, cols, row, col) {                          
                    board[row][col]='X'
                }
            }
        }
    }
}


func isRegion(board [][]byte, rows, cols, row, col int) bool {

     if col-1==0 && board[row][col-1]=='O' {
    	return false
    } else if row-1==0 && board[row-1][col]=='O' {
    	return false
    } else if row+1==rows-1 && board[row+1][col]=='O' {
    	return false
    } else if col+1==cols-1 && board[row][col+1]=='O' {
    	return false
    }

    left, right, up, down := board[row][col-1]=='X',  board[row][col+1]=='X',  board[row-1][col]=='X', board[row+1][col]=='X'

    if left && right && up && down {        
        return true
    }
    
    board[row][col]='X'

    if !left {
        if left=isRegion(board, rows, cols, row, col-1); !left {
            board[row][col]='O'
            return false
        }
    }
    
    if !right {
        if right=isRegion(board, rows, cols, row, col+1); !right {
            board[row][col]='O'
            return false
        }
    }
    
    if !up {
        if up=isRegion(board, rows, cols, row-1, col); !up {
            board[row][col]='O'
            return false
        }
    }
        
    if !down {
        if down=isRegion(board, rows, cols, row+1, col); !down {
            board[row][col]='O'
            return false
        }
    }

    board[row][col]='O'
    return true
}



