package main

import "fmt"

func main() {
	var board = [][]int{
		[]int{0, 1, 0},
		[]int{0, 0, 1},
		[]int{1, 1, 1},
		[]int{0, 0, 0},
	}
	gameOfLife3(board)
	fmt.Println(board)
}

/**
直接做法：另开空间保存
*/
func gameOfLife(board [][]int) {
	if len(board) == 0 {
		return
	}
	ret := make([][]int, len(board))
	for x := 0; x < len(board); x++ {
		ret[x] = make([]int, len(board[0]))
		for y := 0; y < len(board[0]); y++ {
			ret[x][y] = board[x][y]
		}
	}
	dx := []int{-1, -1, 0, 1, 1, 1, 0, -1} //上、右上、右、右下、下、左下、左、左上
	dy := []int{0, 1, 1, 1, 0, -1, -1, -1}
	for x := 0; x < len(ret); x++ {
		for y := 0; y < len(ret[0]); y++ {
			live := 0
			dead := 0
			for i := 0; i < 8; i++ {
				xx, yy := x+dx[i], y+dy[i]
				if xx >= 0 && xx < len(ret) && yy >= 0 && yy < len(ret[0]) {
					if ret[xx][yy] == 1 {
						live++
					} else {
						dead++
					}
				}
			}
			if ret[x][y] == 1 && (live < 2 || live > 3) { //1, 3
				board[x][y] = 0
			}
			if ret[x][y] == 0 && live == 3 { //4
				board[x][y] = 1
			}
		}
	}
}

/**
位运算：最后一位表示最新状态、倒数第二位表示原本状态
*/
func gameOfLife2(board [][]int) {
	if len(board) == 0 {
		return
	}
	dx := []int{-1, -1, 0, 1, 1, 1, 0, -1} //上、右上、右、右下、下、左下、左、左上
	dy := []int{0, 1, 1, 1, 0, -1, -1, -1}
	for x := 0; x < len(board); x++ {
		for y := 0; y < len(board[0]); y++ {
			board[x][y] = board[x][y] << 1
		}
	}
	for x := 0; x < len(board); x++ {
		for y := 0; y < len(board[0]); y++ {
			live := 0
			dead := 0
			for i := 0; i < 8; i++ {
				xx, yy := x+dx[i], y+dy[i]
				if xx >= 0 && xx < len(board) && yy >= 0 && yy < len(board[0]) {
					if (board[xx][yy] >> 1) == 1 {
						live++
					} else {
						dead++
					}
				}
			}
			if (board[x][y]>>1) == 1 && (live < 2 || live > 3) { //1, 3
				continue
			}
			if (board[x][y]>>1) == 0 && live == 3 { //4
				board[x][y] += 1
				continue
			}
			board[x][y] += (board[x][y] >> 1)
		}
	}
	for x := 0; x < len(board); x++ {
		for y := 0; y < len(board[0]); y++ {
			board[x][y] = board[x][y] % 2
		}
	}
}

/**
递归做法
	每一个递归的层级都计算出当前的位置应该更新的状态，然后递归调用下一层，等递归调用返回的时候再更新本层级的状态
*/
func gameOfLife3(board [][]int) {
	if len(board) == 0 {
		return
	}
	recursive(board, 0, 0)
}

func recursive(board [][]int, x, y int) {
	if y == len(board[0]) {
		x++
		y = 0
	}
	if x == len(board) { //递归的出口
		return
	}
	dx := []int{-1, -1, 0, 1, 1, 1, 0, -1} //上、右上、右、右下、下、左下、左、左上
	dy := []int{0, 1, 1, 1, 0, -1, -1, -1}
	live, dead, curr := 0, 0, board[x][y]
	for i := 0; i < 8; i++ {
		xx, yy := x+dx[i], y+dy[i]
		if xx >= 0 && xx < len(board) && yy >= 0 && yy < len(board[0]) {
			if board[xx][yy] == 1 {
				live++
			} else {
				dead++
			}
		}
	}
	if curr == 1 && (live < 2 || live > 3) { //1, 3
		curr = 0
	}
	if curr == 0 && live == 3 { //4
		curr = 1
	}
	recursive(board, x, y+1) //递归下一层
	board[x][y] = curr       //更新本层的状态
}
