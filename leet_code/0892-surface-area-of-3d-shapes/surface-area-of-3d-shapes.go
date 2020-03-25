package main

import "fmt"

func main() {
	var grid = [][]int{
		[]int{2},
	}
	fmt.Println(surfaceArea2(grid))
	grid = [][]int{
		[]int{1, 2},
		[]int{3, 4},
	}
	fmt.Println(surfaceArea2(grid))
	grid = [][]int{
		[]int{1, 1, 1},
		[]int{1, 0, 1},
		[]int{1, 1, 1},
	}
	fmt.Println(surfaceArea2(grid))
}

/**
投影: (这个方法不好使，我忽略了中间有个洞的情况)
	找三维形体的三视图，三视图的面积乘以 2 即为答案（忽略中间有洞的情况确实是可行的）
	求 grid 每一行、每一列的最大值，以及一共有多少个非0元素
*/
func surfaceArea(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	dx := make([]int, len(grid))
	dy := make([]int, len(grid[0]))
	total := 0
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			if grid[x][y] == 0 {
				continue
			}
			total++
			if grid[x][y] > dx[x] {
				dx[x] = grid[x][y]
			}
			if grid[x][y] > dy[y] {
				dy[y] = grid[x][y]
			}
		}
	}
	for x := 0; x < len(grid); x++ {
		total += dx[x]
	}
	for y := 0; y < len(grid[0]); y++ {
		total += dy[y]
	}
	return total * 2
}

/**
累加法:
	单独计算每一个形体 v = grid[x][y] 贡献的表面积，然后求所有 v 的和
	对于 v = grid[x][y] 贡献的表面积 s，有：
	（1）如果 grid[x][y] > 0，则 v 贡献上帝面和下底面两个面积，s += 2
	（2）判断 grid[x][y] 相邻的位置是否有其它形体 vv，如果有，计算 v 比 vv 高出的部分 v - vv，只有当 v 大于 vv 时，v 才能贡献表面积
		如果没有，则直接贡献表面积
*/
func surfaceArea2(grid [][]int) int {
	total := 0
	dx := []int{-1, 0, 1, 0}
	dy := []int{0, 1, 0, -1}
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			if grid[x][y] == 0 {
				continue
			}
			total += 2	//贡献上底面和下底面的面积
			for i := 0; i < 4; i++ {	//判断四周是否有其它形体
				xx := x + dx[i]
				yy := y + dy[i]
				if xx < 0 || yy < 0 || xx >= len(grid) || yy >= len(grid[0]) {	//某一面在边界部分，直接贡献表面积
					total += grid[x][y]
					continue
				}
				if grid[x][y] > grid[xx][yy] {	//比周围的形体高时，贡献表面积
					total += grid[x][y] - grid[xx][yy]
				}
			}
		}
	}
	return total
}
