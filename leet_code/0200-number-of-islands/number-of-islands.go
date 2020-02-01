package main

import "fmt"

func main() {
	grid := [][]byte{
		[]byte{'1', '1', '1', '1', '0'},
		[]byte{'1', '1', '0', '1', '0'},
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'0', '0', '0', '0', '0'},
	}
	fmt.Println(numIslands(grid))

	grid = [][]byte{
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'0', '0', '1', '0', '0'},
		[]byte{'0', '0', '0', '1', '1'},
	}
	fmt.Println(numIslands(grid))

	grid = [][]byte{
		[]byte{'1', '1', '1', '1', '0'},
		[]byte{'1', '1', '0', '1', '0'},
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'0', '0', '0', '0', '0'},
	}
	fmt.Println(numIslands(grid))
}

/**
染色法：
	遍历每一个元素，当发现该元素为 1 时，将 该元素'染色'为 0，并递归判断该元素的上下左右元素是否为 1，为 1 则递归 '染色' 过程
*/
func numIslands(grid [][]byte) int {
	ret := 0
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			if grid[x][y] == '1' {
				ret++
				dfsFloodFill(grid, x, y)
			}
		}
	}
	return ret
}

//深度优先搜索并染色
func dfsFloodFill(grid [][]byte, x, y int) {
	dx := [4]int{-1, 1, 0, 0} //移动方向：上下左右
	dy := [4]int{0, 0, -1, 1}
	grid[x][y] = '0' //将该点染色
	for i := 0; i < 4; i++ {
		xx := x + dx[i]
		yy := y + dy[i]
		if xx >= 0 && xx < len(grid) && yy >= 0 && yy < len(grid[0]) && grid[xx][yy] == '1' {
			dfsFloodFill(grid, xx, yy)
		}
	}
}


