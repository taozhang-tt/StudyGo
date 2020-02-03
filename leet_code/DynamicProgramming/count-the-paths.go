package main

import "fmt"

func main() {
	grid := [][]byte{
		[]byte{0, 0, 0, 0, 0, 0, 0, 0},
		[]byte{0, 0, 1, 0, 0, 0, 1, 0},
		[]byte{0, 0, 0, 0, 1, 0, 0, 0},
		[]byte{1, 0, 1, 0, 0, 1, 0, 0},
		[]byte{0, 0, 1, 0, 0, 0, 0, 0},
		[]byte{0, 0, 0, 1, 1, 0, 1, 0},
		[]byte{0, 1, 0, 0, 0, 1, 0, 0},
		[]byte{0, 0, 0, 0, 0, 0, 0, 0},
	}
	fmt.Println(countThePaths(grid))
}

/*
深度优先搜索做法：
	对于坐标 (x, y) 到 （m, n) 的路径数 = (x+1, y) 到 (m, n)的路径数 + (x, y+1) 到 (m, n)的路径数
*/
func countThePaths(grid [][]byte) int {
	return F(grid, 0, 0)
}

func F(grid [][]byte, x, y int) int {
	m := len(grid)
	n := len(grid[0])
	if x == m-1 && y == n-1 { //到达终点
		return 1
	}
	if x == m || y == n || grid[x][y] == 1 { //不合法返回0
		return 0
	}
	return F(grid, x+1, y) + F(grid, x, y+1)
}
