package main

import (
	"fmt"
	"strconv"
)

/**
37. 解数独
	https://leetcode-cn.com/problems/sudoku-solver/
题目描述：
	编写一个程序，通过已填充的空格来解决数独问题。
	一个数独的解法需遵循如下规则：
	1. 数字 1-9 在每一行只能出现一次。
	2. 数字 1-9 在每一列只能出现一次。
	3. 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
	空白格用 '.' 表示。
*/

func main() {
	param := [][]byte{
		[]byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		[]byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		[]byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		[]byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		[]byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		[]byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		[]byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		[]byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		[]byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	solveSudoku(param)
}

func solveSudoku(board [][]byte) {
	if board, ok := solve(board); ok {
		fmt.Println(board)
	}
}

func solve(board [][]byte) ([][]byte, bool) {
	if !isValid(board) {
		return board, false
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			cell := string(board[i][j])
			if cell == "." {
				for _, item := range [9]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'} {
					board[i][j] = item
					if board, ok := solve(board); ok {
						return board, ok
					} else {
						board[i][j] = '.'
					}
				}
				return board, false
			}
		}
	}
	return board, true
}

//判断数独当前的状态是否有效
func isValid(board [][]byte) bool {
	var boxes, rows, cols [9][10]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			cell := string(board[i][j])
			if cell != "." {
				cellInt, err := strconv.Atoi(cell)
				if err != nil {
					fmt.Println(cell)
					return false
				}
				if rows[i][cellInt] == 1 {
					return false
				}
				rows[i][cellInt] = 1
				if cols[j][cellInt] == 1 {
					return false
				}
				cols[j][cellInt] = 1
				if boxes[i/3*3+j/3][cellInt] == 1 {
					return false
				}
				boxes[i/3*3+j/3][cellInt] = 1
			}
		}
	}
	return true
}
