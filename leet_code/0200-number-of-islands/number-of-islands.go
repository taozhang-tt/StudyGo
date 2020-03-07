package main

import "fmt"

/**
200. 岛屿数量
	https://leetcode-cn.com/problems/number-of-islands/
题目描述：
	给定一个由 '1'（陆地）和 '0'（水）组成的的二维网格，计算岛屿的数量。一个岛被水包围，并且它是通过水平方向或垂直方向上相邻的陆地连接而成的。你可以假设网格的四个边均被水包围。
示例 1:
	输入:
	11110
	11010
	11000
	00000
	输出: 1
*/

func main() {
	grid := [][]byte{
		[]byte{'1', '1', '1', '1', '0'},
		[]byte{'1', '1', '0', '1', '0'},
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'0', '0', '0', '0', '0'},
	}
	fmt.Println(numIslands2(grid))

	grid = [][]byte{
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'0', '0', '1', '0', '0'},
		[]byte{'0', '0', '0', '1', '1'},
	}
	fmt.Println(numIslands2(grid))

	grid = [][]byte{
		[]byte{'1', '1', '1', '1', '0'},
		[]byte{'1', '1', '0', '1', '0'},
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'0', '0', '0', '0', '0'},
	}
	fmt.Println(numIslands2(grid))

	grid = [][]byte{
		[]byte{'1', '0', '0', '1', '1', '1', '0', '1', '1', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0'},
		[]byte{'1', '0', '0', '1', '1', '0', '0', '1', '0', '0', '0', '1', '0', '1', '0', '1', '0', '0', '1', '0'},
		[]byte{'0', '0', '0', '1', '1', '1', '1', '0', '1', '0', '1', '1', '0', '0', '0', '0', '1', '0', '1', '0'},
		[]byte{'0', '0', '0', '1', '1', '0', '0', '1', '0', '0', '0', '1', '1', '1', '0', '0', '1', '0', '0', '1'},
		[]byte{'0', '0', '0', '0', '0', '0', '0', '1', '1', '1', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0'},
		[]byte{'1', '0', '0', '0', '0', '1', '0', '1', '0', '1', '1', '0', '0', '0', '0', '0', '0', '1', '0', '1'},
		[]byte{'0', '0', '0', '1', '0', '0', '0', '1', '0', '1', '0', '1', '0', '1', '0', '1', '0', '1', '0', '1'},
		[]byte{'0', '0', '0', '1', '0', '1', '0', '0', '1', '1', '0', '1', '0', '1', '1', '0', '1', '1', '1', '0'},
		[]byte{'0', '0', '0', '0', '1', '0', '0', '1', '1', '0', '0', '0', '0', '1', '0', '0', '0', '1', '0', '1'},
		[]byte{'0', '0', '1', '0', '0', '1', '0', '0', '0', '0', '0', '1', '0', '0', '1', '0', '0', '0', '1', '0'},
		[]byte{'1', '0', '0', '1', '0', '0', '0', '0', '0', '0', '0', '1', '0', '0', '1', '0', '1', '0', '1', '0'},
		[]byte{'0', '1', '0', '0', '0', '1', '0', '1', '0', '1', '1', '0', '1', '1', '1', '0', '1', '1', '0', '0'},
		[]byte{'1', '1', '0', '1', '0', '0', '0', '0', '1', '0', '0', '0', '0', '0', '0', '1', '0', '0', '0', '1'},
		[]byte{'0', '1', '0', '0', '1', '1', '1', '0', '0', '0', '1', '1', '1', '1', '1', '0', '1', '0', '0', '0'},
		[]byte{'0', '0', '1', '1', '1', '0', '0', '0', '1', '1', '0', '0', '0', '1', '0', '1', '0', '0', '0', '0'},
		[]byte{'1', '0', '0', '1', '0', '1', '0', '0', '0', '0', '1', '0', '0', '0', '1', '0', '1', '0', '1', '1'},
		[]byte{'1', '0', '1', '0', '0', '0', '0', '0', '0', '1', '0', '0', '0', '1', '0', '1', '0', '0', '0', '0'},
		[]byte{'0', '1', '1', '0', '0', '0', '1', '1', '1', '0', '1', '0', '1', '0', '1', '1', '1', '1', '0', '0'},
		[]byte{'0', '1', '0', '0', '0', '0', '1', '1', '0', '0', '1', '0', '1', '0', '0', '1', '0', '0', '1', '1'},
		[]byte{'0', '0', '0', '0', '0', '0', '1', '1', '1', '1', '0', '1', '0', '0', '0', '1', '1', '0', '0', '0'},
	}
	fmt.Println(numIslands2(grid))
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

//使用并查集
func numIslands2(grid [][]byte) int {
	dx := [4]int{-1, 1, 0, 0}
	dy := [4]int{0, 0, -1, 1}
	height := len(grid)
	if height == 0 {
		return 0
	}
	width := len(grid[0])
	unionFind := Construct(grid)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				for k := 0; k < 4; k++ {
					x := i + dx[k]
					y := j + dy[k]
					if x >= 0 && x < height && y >= 0 && y < width && grid[x][y] == '1' {
						unionFind.Union(i*width+j, x*width+y)
					}
				}
			}
		}
	}
	return unionFind.count
}

type UnionFind struct {
	roots []int
	count int
}

func Construct(grid [][]byte) *UnionFind {
	height := len(grid)
	width := len(grid[0])
	unionFind := &UnionFind{
		roots: make([]int, height*width, height*width),
		count: 0,
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				unionFind.roots[i*width+j] = i*width + j
				unionFind.count++
			}
		}
	}
	return unionFind
}

func (obj *UnionFind) Find(index int) int {
	if obj.roots[index] != index {
		obj.roots[index] = obj.Find(obj.roots[index])
	}
	return obj.roots[index]
}

func (obj *UnionFind) Union(p, q int) {
	pRoot := obj.Find(p)
	qRoot := obj.Find(q)
	if pRoot != qRoot {
		obj.roots[pRoot] = qRoot
		obj.count--
	}
}
