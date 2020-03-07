package main

import (
	"fmt"
	"strconv"
)

/**
36. 有效的数独
	https://leetcode-cn.com/problems/valid-sudoku/
题目描述：
	判断一个 9x9 的数独是否有效。只需要根据以下规则，验证已经填入的数字是否有效即可。
	1. 数字 1-9 在每一行只能出现一次。
	2. 数字 1-9 在每一列只能出现一次。
	3. 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
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

	fmt.Println(isValidSudoku2(param))

}

/**
数独的定义：
  数字 1-9 在每一行只能出现一次。
  数字 1-9 在每一列只能出现一次。
  数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
*/
func isValidSudoku(board [][]byte) bool {
	//判断行是否有重复元素
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == byte('.') {
				continue
			}
			for i := col + 1; i < 9; i++ {
				if board[row][col] == board[row][i] {
					return false
				}
			}
		}
	}
	//判断列是否有重复元素
	for col := 0; col < 9; col++ {
		for row := 0; row < 9; row++ {
			for i := row + 1; i < 9; i++ {
				if board[row][col] == byte('.') {
					continue
				}
				if board[row][col] == board[i][col] {
					return false
				}
			}
		}
	}
	//判断 3*3 的box是否满足
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			temp := [10]int{}
			for ii := i * 3; ii < i*3+3; ii++ {
				for jj := j * 3; jj < j*3+3; jj++ {
					item := string(board[ii][jj])
					if item != "." {
						itemInt, _ := strconv.Atoi(item)
						if temp[itemInt] == 1 {
							return false
						}
						temp[itemInt] = 1
					}
				}
			}
		}
	}
	return true
}

//加速，使用空间换时间
/**
使用 rows 这个二维数组缓存每一行已经出现的数字，每次遍历一个元素时，判断该元素是否已经在该行出现过；二维数组的第一维代表行号，第二维代表该行出现的元素
使用 cols 这个二维数组缓存每一列已经出现的数字，每次遍历一个元素时，判断该元素是否已经在该列出现过；二维数组的第一维代表列号，第二维代表该列出现的元素

3*3 的 box 判断：
  拆分：对于 9*9 的方格，可以拆分为 9 个 3*3 的 box，从左到右从上到下依次编号为 0 ~ 8；
  缓存：那么我们可以用一个二维数组（[9][9])缓存所有box中出现的元素；为了方便直接记录，我们用一个 [9][10] 的数组缓存，这样数字9也可以被直接丢进去
  映射：对于 9*9 方格中的任一元素 （i, j)，怎么确定它属于哪个box呢？对于元素 (i, j)，其所属的 box 编号为：i/3*3 + j/3；（这里的 / 代表整除)
*/
func isValidSudoku2(board [][]byte) bool {
	var boxes, rows, cols [9][10]int
	for i := 0; i < 9; i++ { //扫描行
		for j := 0; j < 9; j++ { //扫描列
			cell := string(board[i][j])
			if cell != "." {
				cellInt, _ := strconv.Atoi(cell)
				if rows[i][cellInt] == 1 { //
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
