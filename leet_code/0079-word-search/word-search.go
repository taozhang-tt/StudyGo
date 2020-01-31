package main

import (
	"fmt"
)

func main() {
	//board := [][]byte{
	//	[]byte{'A','B','C','E'},
	//	[]byte{'S','F','C','S'},
	//	[]byte{'A','D','E','E'},
	//}
	board := [][]byte {
		[]byte{'a', 'b'},
	}
	fmt.Println(exist(board, "ba"))
}

func exist(board [][]byte, word string) bool {
	for x:=0; x<len(board); x++ {
		for y:=0; y<len(board[0]); y++ {
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
	for i:=0; i<4; i++ {
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