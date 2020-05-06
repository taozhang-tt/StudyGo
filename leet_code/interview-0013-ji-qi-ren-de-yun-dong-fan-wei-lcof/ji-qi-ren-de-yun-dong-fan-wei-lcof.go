package main

import "fmt"

/**
面试题13. 机器人的运动范围
	https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/
题目描述：
	地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？
示例 1：
	输入：m = 2, n = 3, k = 1
	输出：3
示例 1：
	输入：m = 3, n = 1, k = 0
	输出：1
提示：
	1 <= n,m <= 100
	0 <= k <= 20
*/

func main() {
	var m, n, k, ret int
	m, n, k = 2, 3, 1
	ret = movingCount2(m, n, k)
	fmt.Println(ret)

	m, n, k = 3, 2, 17
	ret = movingCount2(m, n, k)
	fmt.Println(ret)
}

/**
BFS 做法
*/
func movingCount(m int, n int, k int) int {
	q, visited, ret := []int{0}, map[int]bool{0: false}, 0 //q 模拟队列，visited 记录已经访问过的元素
	for len(q) > 0 {
		x, y := q[0]/n, q[0]%n
		q = q[1:]
		ret++
		lx, ly := x, y+1                                              //右侧元素
		if _, exist := visited[lx*n+ly]; !exist && ly < n && lx < m { //判断是否被访问过且元素是否越界
			visited[lx*n+ly] = true //无论是否满足要求都记录已访问，免得下次还要对这个坐标进行判断
			if (lx%10 + lx/10%10 + lx/100 + ly%10 + ly/10%10 + ly/100) <= k {
				q = append(q, lx*n+ly)
			}
		}

		nx, ny := x+1, y //下侧元素
		if _, exist := visited[nx*n+ny]; !exist && nx < m && ny < n {
			visited[nx*n+ny] = true
			if (nx%10 + nx/10%10 + nx/100 + ny%10 + ny/10%10 + ny/100) <= k {
				q = append(q, nx*n+ny)
			}
		}
	}
	return ret
}

/**
递推做法：
	1. 从左上角到右下角逐个判断，我们发现每次只需要向又和向下移动一步即可
	2. 对于元素 vis[x][y]，
		2.1 如果 x, y 数位上的数字和大于 k，那该元素肯定不满足要求，标记为 0
		2.2 如果 x, y 数位上的数字和小于 k，需要进一步判断该元素是否可达，这时只需要判断 vis[x-1][y]、vis[x][y-1]是否可达，可达标记为 1，不可达标记为 0
*/
func movingCount2(m int, n int, k int) int {
	ret, visited := 1, make([][]int, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]int, n)
	}
	visited[0][0] = 1
	for x := 0; x < m; x++ {
		for y := 0; y < n; y++ {
			if (x == 0 && y == 0) || (x%10+x/10%10+x/100+y%10+y/10%10+y/100) > k {	//元素不满足数位上的数字之和小于等于 k的要求
				continue
			}
			if x-1 >= 0 {	//判断 vis[x-1][y] 是否可达
				visited[x][y] = visited[x][y] | visited[x-1][y]
			}
			if y-1 >= 0 {	//判断 vis[x][y-1] 是否可达
				visited[x][y] = visited[x][y] | visited[x][y-1]
			}
			ret += visited[x][y]
		}
	}
	return ret
}
