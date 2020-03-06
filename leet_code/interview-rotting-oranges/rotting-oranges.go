package main

import (
	"fmt"
)

func main() {
	grid := [][]int{
		[]int{2, 1, 1},
		[]int{1, 1, 0},
		[]int{0, 1, 1},
	}
	fmt.Println(orangesRotting(grid))

	grid = [][]int{
		[]int{2, 1, 1},
		[]int{0, 1, 1},
		[]int{1, 0, 1},
	}
	fmt.Println(orangesRotting(grid))

	grid = [][]int{
		[]int{0, 2},
	}
	fmt.Println(orangesRotting(grid))

	grid = [][]int{
		[]int{1},
		[]int{1},
		[]int{1},
		[]int{1},
	}
	fmt.Println(orangesRotting(grid))

	grid = [][]int{
		[]int{0},
	}
	fmt.Println(orangesRotting(grid))
}

/**
广度优先搜索做法
*/
func orangesRotting(grid [][]int) int {
	dx := []int{-1, 0, 1, 0} //上右下左
	dy := []int{0, 1, 0, -1}
	rows := len(grid)
	cols := len(grid[0])
	queue := make([][]int, 0)
	minute := 0
	//初始化将所有已经腐烂的橘子入队
	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			if grid[x][y] == 2 {
				queue = append(queue, []int{x, y, minute})
			}
		}
	}
	for len(queue) > 0 {
		item := queue[0]  //取出队首元素
		queue = queue[1:] //队首元素出队列
		x := item[0]
		y := item[1]
		minute = item[2]
		for i := 0; i < 4; i++ { //向4个方向感染新鲜橘子
			xx := x + dx[i]
			yy := y + dy[i]
			if xx >= 0 && xx < rows && yy >= 0 && yy < cols && grid[xx][yy] == 1 {
				grid[xx][yy] = 2
				queue = append(queue, []int{xx, yy, minute + 1})
			}
		}
	}

	for x := 0; x < rows; x++ { //检查一遍还有没有未被感染的橘子
		for y := 0; y < cols; y++ {
			if grid[x][y] == 1 {
				return -1
			}
		}
	}
	return minute
}

