package main

import "fmt"

/**
994. 腐烂的橘子
	https://leetcode-cn.com/problems/rotting-oranges/
题目描述：
	在给定的网格中，每个单元格可以有以下三个值之一：
	值 0 代表空单元格；
	值 1 代表新鲜橘子；
	值 2 代表腐烂的橘子。
	每分钟，任何与腐烂的橘子（在 4 个正方向上）相邻的新鲜橘子都会腐烂。
	返回直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回 -1。
*/

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
