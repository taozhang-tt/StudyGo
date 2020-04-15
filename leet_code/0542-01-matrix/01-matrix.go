package main

import "fmt"

/**
542. 01 矩阵
	https://leetcode-cn.com/problems/01-matrix/
题目描述
	给定一个由 0 和 1 组成的矩阵，找出每个元素到最近的 0 的距离。
	两个相邻元素间的距离为 1 。
示例 1:
	输入:
		0 0 0
		0 1 0
		0 0 0
	输出:
		0 0 0
		0 1 0
		0 0 0
*/

func main() {
	var matrix [][]int
	matrix = [][]int{
		[]int{0, 0, 0},
		[]int{0, 1, 0},
		[]int{0, 0, 0},
	}
	fmt.Println(updateMatrix(matrix))

	matrix = [][]int{
		[]int{0, 0, 0},
		[]int{0, 1, 0},
		[]int{1, 1, 1},
	}
	fmt.Println(updateMatrix(matrix))
}

/**
多源广度优先搜索
*/
func updateMatrix(matrix [][]int) [][]int {
	dx := []int{-1, 0, 1, 0}
	dy := []int{0, 1, 0, -1}
	queue := make([][]int, 0)
	visited := make([][]bool, len(matrix))
	for i := 0; i < len(matrix); i++ {
		visited[i] = make([]bool, len(matrix[0]))
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				queue = append(queue, []int{i, j, 0})
				visited[i][j] = true
			} else {
				visited[i][j] = false
			}
		}
	}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		dis := curr[2]
		for i := 0; i < 4; i++ {
			x := curr[0] + dx[i]
			y := curr[1] + dy[i]
			if x >= 0 && x < len(matrix) && y >= 0 && y < len(matrix[0]) && visited[x][y] != true {
				queue = append(queue, []int{x, y, dis + 1})
				visited[x][y] = true
				matrix[x][y] = dis + 1
			}
		}
	}
	return matrix
}
