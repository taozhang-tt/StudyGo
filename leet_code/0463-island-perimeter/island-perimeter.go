package main

import "fmt"

/**
463. 岛屿的周长
	https://leetcode-cn.com/problems/island-perimeter/
题目描述
	给定一个包含 0 和 1 的二维网格地图，其中 1 表示陆地 0 表示水域。
	网格中的格子水平和垂直方向相连（对角线方向不相连）。整个网格被水完全包围，但其中恰好有一个岛屿（或者说，一个或多个表示陆地的格子相连组成的岛屿）。
	岛屿中没有“湖”（“湖” 指水域在岛屿内部且不和岛屿周围的水相连）。格子是边长为 1 的正方形。网格为长方形，且宽度和高度均不超过 100 。计算这个岛屿的周长。
示例 :
	输入:
	[[0,1,0,0],
	[1,1,1,0],
	[0,1,0,0],
	[1,1,0,0]]
	输出: 16
*/

func main() {
	var grid = [][]int{
		[]int{0, 1, 0, 0},
		[]int{1, 1, 1, 0},
		[]int{0, 1, 0, 0},
		[]int{1, 1, 0, 0},
	}
	fmt.Println(islandPerimeter(grid))
}

func islandPerimeter(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				return dfs(grid, i, j)
			}
		}
	}
	return 0
}

func dfs(grid [][]int, x, y int) int { //x 代表行号，y 代表列号
	dx := []int{-1, 0, 1, 0} //上、右、下、左
	dy := []int{0, 1, 0, -1}
	cnt := 0
	grid[x][y] = 2 //被访问过的标记为2
	for i := 0; i < 4; i++ {
		xx := x + dx[i]
		yy := y + dy[i]
		if xx >= 0 && xx < len(grid) && yy >= 0 && yy < len(grid[0]) { //边界内分三种情况：已经被访问过的、未被访问过且为1的，未被访问过且为0的
			if grid[xx][yy] == 1 { //未被访问过且为1的
				cnt += dfs(grid, xx, yy)
			} else if grid[xx][yy] == 0 { //未被访问过且为0的
				cnt++
			}
		} else { //到达边界加 1
			cnt++
		}
	}
	return cnt
}
