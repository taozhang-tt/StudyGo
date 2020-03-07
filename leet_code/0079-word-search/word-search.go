package main

import (
	"fmt"
)

/**
79. 单词搜索
	https://leetcode-cn.com/problems/word-search/
题目描述：
	给定一个二维网格和一个单词，找出该单词是否存在于网格中。
	单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
示例:
	board =
	[
	  ['A','B','C','E'],
	  ['S','F','C','S'],
	  ['A','D','E','E']
	]
	给定 word = "ABCCED", 返回 true.
	给定 word = "SEE", 返回 true.
	给定 word = "ABCB", 返回 false.
*/

func main() {
	//board := [][]byte{
	//	[]byte{'A','B','C','E'},
	//	[]byte{'S','F','C','S'},
	//	[]byte{'A','D','E','E'},
	//}
	board := [][]byte{
		[]byte{'a', 'b'},
	}
	fmt.Println(exist(board, "ba"))
}

func exist(board [][]byte, word string) bool {
	for x := 0; x < len(board); x++ {
		for y := 0; y < len(board[0]); y++ {
			if _, ok := dfs(board, word, "", x, y); ok {
				return ok
			}
		}
	}

	return false
}

//利用搜索
func dfs(board [][]byte, target string, current string, x, y int) ([][]byte, bool) {
	current += string(board[x][y])
	if current == target {
		return board, true
	}
	//剪枝
	if current[len(current)-1] != target[len(current)-1] {
		return board, false
	}
	if len(current) == len(target) {
		return board, false
	}
	//移动方向，x代表行坐标，y代表列坐标，上下左右
	dx := [4]int{-1, 1, 0, 0}
	dy := [4]int{0, 0, -1, 1}
	temp := board[x][y]
	board[x][y] = '@'
	for i := 0; i < 4; i++ {
		xx := x + dx[i]
		yy := y + dy[i]
		if xx >= 0 && xx < len(board) && yy >= 0 && yy < len(board[0]) && board[xx][yy] != '@' {
			if board, ok := dfs(board, target, current, xx, yy); ok {
				return board, ok
			}
		}
	}
	board[x][y] = temp
	return board, false
}
