package main

import (
	"fmt"
)

/**
1162. 地图分析
	https://leetcode-cn.com/problems/as-far-from-land-as-possible/
题目描述：
	你现在手里有一份大小为 N x N 的『地图』（网格） grid，上面的每个『区域』（单元格）都用 0 和 1 标记好了。其中 0 代表海洋，1 代表陆地，你知道距离陆地区域最远的海洋区域是是哪一个吗？请返回该海洋区域到离它最近的陆地区域的距离。
	我们这里说的距离是『曼哈顿距离』（ Manhattan Distance）：(x0, y0) 和 (x1, y1) 这两个区域之间的距离是 |x0 - x1| + |y0 - y1| 。
	如果我们的地图上只有陆地或者海洋，请返回 -1。

示例：
	输入：[[1,0,1],[0,0,0],[1,0,1]]
	输出：2
	解释：
	海洋区域 (1, 1) 和所有陆地区域之间的距离都达到最大，最大距离为 2。

	输入：[[1,0,0],[0,0,0],[0,0,0]]
	输出：4
	解释：
	海洋区域 (2, 2) 和所有陆地区域之间的距离都达到最大，最大距离为 4。
*/

func main() {
	var grid = [][]int{
		[]int{1, 0, 1},
		[]int{0, 0, 0},
		[]int{1, 0, 1},
	}
	ret := maxDistance2(grid)
	fmt.Println(ret)

	grid = [][]int{
		[]int{1, 0, 0},
		[]int{0, 0, 0},
		[]int{0, 0, 0},
	}
	ret = maxDistance2(grid)
	fmt.Println(ret)

	grid = [][]int{
		[]int{0, 0, 0, 0},
		[]int{0, 0, 0, 0},
		[]int{0, 0, 0, 0},
		[]int{0, 0, 0, 0},
	}
	ret = maxDistance2(grid)
	fmt.Println(ret)

	grid = [][]int{
		[]int{1, 1, 1, 1},
		[]int{1, 1, 1, 1},
		[]int{1, 1, 1, 1},
		[]int{1, 1, 1, 1},
	}
	ret = maxDistance2(grid)
	fmt.Println(ret)
}

/**
暴力法、宽度优先搜索：(超时)
	遍历每一个海洋区域，求它与陆地的最近距离 d，取所有海洋对应的 d 的最大值，即为结果
*/
func maxDistance(grid [][]int) int {
	max := -1
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			if grid[x][y] == 0 {
				dis := bfsDistance(x, y, grid) //求每个海洋到陆地的最近距离
				if dis > max {
					max = dis
				}
			}
		}
	}
	return max
}

//广度优先搜索海洋 grid[x][y] 到陆地的最近距离
func bfsDistance(x, y int, grid [][]int) int {
	dx := []int{-1, 0, 1, 0} //定义方向数组
	dy := []int{0, 1, 0, -1}
	queue := [][]int{ //队列初始化放入海洋的位置
		[]int{x, y, 0},
	}
	visited := make([][]int, len(grid)) //初始化一个二维切片用于记录该位置已被访问过，不要重复放入队列
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]int, len(grid))
		for j := 0; j < len(grid[0]); j++ {
			visited[i][j] = 0
		}
	}
	for len(queue) > 0 { //每次取队首元素
		x, y, dis := queue[0][0], queue[0][1], queue[0][2]
		visited[x][y] = 1 //标记该位置已被访问过
		queue = queue[1:] //弹出队首元素
		for i := 0; i < 4; i++ {
			xx, yy := x+dx[i], y+dy[i]
			if xx >= 0 && xx < len(grid) && yy >= 0 && yy < len(grid[0]) && visited[xx][yy] == 0 {
				if grid[xx][yy] == 1 { //判断该位置是否为陆地，为陆地则直接返回
					return dis + 1
				} else {
					queue = append(queue, []int{xx, yy, dis + 1}) //不为陆地放入队列，并记录当前位置距离海洋 grid[x][y] 的距离
				}
			}
		}
	}
	return -1
}

/**
动态规划：（超时）
	状态定义：dp[x][y]，代表海洋 grid[x][y] 到陆地的最近距离
	状态转换：
		（1）从左上到右下：
			如果 grid[x][y] 是陆地，则 dp[x][y] = 0
			如果 grid[x][y] 是海洋，则 dp[x][y] = min(dp[x-1][y], dp[x][y-1]) + 1
		(2) 从右下到左上：
			如果 grid[x][y] 是陆地，则 dp[x][y] = 0
			如果 grid[x][y] 是海洋，则 dp[x][y] = min(dp[x+1][y], dp[x][y+1]) + 1
*/
func maxDistance2(grid [][]int) int {
	dp := make([][]int, len(grid))
	for x := 0; x < len(grid); x++ { //初始化dp状态
		dp[x] = make([]int, len(grid[0]))
		for y := 0; y < len(grid[0]); y++ {
			if grid[x][y] == 0 { //grid[x][y]为陆地时，dp[x][y] 为 0，否则初始为一个较大的值
				dp[x][y] = len(grid) * len(grid[0])
			}
		}
	}
	for x := 0; x < len(grid); x++ { //从左上到右下进行一次动态规划
		for y := 0; y < len(grid[0]); y++ {
			if x-1 >= 0 && dp[x][y] > dp[x-1][y]+1 {
				dp[x][y] = dp[x-1][y] + 1
			}
			if y-1 >= 0 && dp[x][y] > dp[x][y-1]+1 {
				dp[x][y] = dp[x][y-1] + 1
			}
		}
	}
	for x := len(grid) - 1; x >= 0; x-- { //从右下到左上进行一次动态规划
		for y := len(grid[0]) - 1; y >= 0; y-- {
			if x+1 < len(grid) && dp[x][y] > dp[x+1][y]+1 {
				dp[x][y] = dp[x+1][y] + 1
			}
			if y+1 < len(grid[0]) && dp[x][y] > dp[x][y+1]+1 {
				dp[x][y] = dp[x][y+1] + 1
			}
		}
	}
	ret := -1
	for x := 0; x < len(grid); x++ {	//遍历一遍dp，寻找最大值
		for y := 0; y < len(grid[0]); y++ {
			if grid[x][y] == 0 && dp[x][y] > ret {
				ret = dp[x][y]
			}
		}
	}
	if ret == len(grid)*len(grid[0]) {
		return -1
	}
	return ret
}

/**
多源广度优先搜索（类似 994. 腐烂的橘子）
	（1）队列初始化为所有陆地的位置，并标记这些位置距离陆地的地址为 0
	（2）循环取队首元素，判断该位置周围（上、右、下、左）是否有海洋，如果有，将该海洋标记为陆地，放入队列尾部，标记该海洋到陆地的距离为队首元素到陆地的距离 +1
	（3）直到队列为空
*/
func maxDistance4(grid [][]int) int {
	if len(grid) == 0 {
		return -1
	}
	queue := make([][]int, 0) //队列用于记录当前有哪些还未被遍历的陆地
	maxDistance := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {	//将所有陆地的位置存入队列，作为广度优先搜索的源头，0 代表该位置距离陆地的最近距离
				queue = append(queue, []int{i, j, 0})
			}
		}
	}
	if len(queue) == 0 || len(queue) == (len(grid))*(len(grid[0])) {	//如果均为陆地或者均为海洋则直接返回
		return -1
	}
	dx := []int{-1, 0, 1, 0}
	dy := []int{0, 1, 0, -1}
	for len(queue) > 0 {
		land := queue[0]
		queue = queue[1:]
		x, y, distance := land[0], land[1], land[2]
		if distance > maxDistance {
			maxDistance = distance
		}
		for i := 0; i < 4; i++ {
			xx, yy := x+dx[i], y+dy[i]
			if xx >= 0 && xx < len(grid) && yy >= 0 && yy < len(grid[0]) && grid[xx][yy] == 0 {
				grid[xx][yy] = 1
				queue = append(queue, []int{xx, yy, distance + 1})
			}
		}
	}
	return maxDistance
}
