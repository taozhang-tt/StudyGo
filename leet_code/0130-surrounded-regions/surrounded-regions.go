package main

import "fmt"

/**
130. 被围绕的区域
	https://leetcode-cn.com/problems/surrounded-regions/
题目描述：
	给定一个二维的矩阵，包含 'X' 和 'O'（字母 O）。
	找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。
示例:
	X X X X
	X O O X
	X X O X
	X O X X
	运行你的函数后，矩阵变为：
	X X X X
	X X X X
	X X X X
	X O X X
解释:
	被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 任何不在边界上，或不与边界上的 'O' 相连的 'O' 最终都会被填充为 'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的
*/

func main() {
	var board [][]byte
	board = [][]byte{
		[]byte{'X', 'X', 'X', 'X'},
		[]byte{'X', 'O', 'O', 'X'},
		[]byte{'X', 'X', 'O', 'X'},
		[]byte{'X', 'O', 'X', 'X'},
	}
	solve(board)
	fmt.Println(board)

	board = [][]byte{
		[]byte{'O','X','X','O','X'},
		[]byte{'X','O','O','X','O'},
		[]byte{'X','O','X','O','X'},
		[]byte{'O','X','O','O','O'},
		[]byte{'X','X','O','X','O'},
	}
	solve(board)
	fmt.Println(board)

	board = [][]byte{
		[]byte{'X', 'X', 'X', 'X', 'X'},
		[]byte{'X', 'X', 'X', 'X', 'X'},
		[]byte{'X', 'X', 'X', 'O', 'X'},
		[]byte{'X', 'X', 'O', 'O', 'O'},
		[]byte{'X', 'X', 'O', 'X', 'O'},
	}
	solve(board)
	fmt.Println(board)
}


/**
先用深度优先搜索所有和边界上的 'O' 连通的点，并标记为 ‘#’
再遍历一遍，将所有的 'O' 替换为 'X'，将所有的 '#' 替换为 'O'
*/
func solve(board [][]byte) {
	for x := 0; x < len(board); x++ {
		for y := 0; y < len(board[0]); y++ {
			if (x == 0 || x == len(board)-1 || y == 0 || y == len(board[0])-1) && board[x][y] == 'O' {
				dfs(board, x, y)
			}
		}
	}

	for x := 0; x < len(board); x++ {
		for y := 0; y < len(board[0]); y++ {
			if board[x][y] == 'O' {
				board[x][y] = 'X'
			}
			if board[x][y] == '#' {
				board[x][y] = 'O'
			}
		}
	}
}

var (
	dx = []int{-1, 0, 1, 0}
	dy = []int{0, 1, 0, -1}
)

func dfs(board [][]byte, x, y int) {
	board[x][y] = '#'
	for i := 0; i < 4; i++ {
		xx := x + dx[i]
		yy := y + dy[i]
		if xx >= 0 && xx < len(board) && yy >= 0 && yy < len(board[0]) && board[xx][yy] == 'O' {
			dfs(board, xx, yy)
		}
	}
}
