package main

import "fmt"

func main() {
	//triangle := [][]int{
	//	[]int{2},
	//	[]int{3, 4},
	//	[]int{6, 5, 7},
	//	[]int{4, 1, 8, 3},
	//}
	triangle := [][]int{
		[]int{46},
		[]int{43, 61},
		[]int{10, -16, 3},
		[]int{-26, 41, 36, -72},
		[]int{-28, -76, -22, 26, 51},
		[]int{56, -53, 38, 67, 86, -45},
		[]int{58, 53, 47, -52, -54, -95, 56},
	}
	fmt.Println(minimumTotal(triangle))
	//fmt.Println(minimumTotal2(triangle))
	fmt.Println(minimumTotal3(triangle))
	fmt.Println(minimumTotal4(triangle))

}

/**
动态规划算法（运行慢 8ms 3.3MB）：
	0. 考虑问题：由终点到起点
	1. 定义DP状态：dp[i, j] 表示由最后一层走到 [i, j] 位置，经历的最小路径和
	2. 状态转换方程：dp[i, j] = min(dp[i+1, j], dp[i+1, j+1])
*/
func minimumTotal(triangle [][]int) int {
	height := len(triangle)
	dp := make([][]int, height)
	for i := height - 1; i >= 0; i-- {
		width := len(triangle[i])
		dp[i] = make([]int, width)
		for j := 0; j < width; j++ {
			if i == height-1 {
				dp[i][j] = triangle[i][j]
			} else {
				if dp[i+1][j] < dp[i+1][j+1] {
					dp[i][j] = dp[i+1][j] + triangle[i][j]
				} else {
					dp[i][j] = dp[i+1][j+1] + triangle[i][j]
				}
			}
		}
	}
	return dp[0][0]
}

/**
动态规划加强版 （4ms 3.1MB）
	思路一样，但是不需要保存所有节点的 dp[i,j]，只需要保留下一层的 dp[j] 即可
*/
func minimumTotal2(triangle [][]int) int {
	dp := triangle[len(triangle)-1]
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			if dp[j] < dp[j+1] {
				dp[j] += triangle[i][j]
			} else {
				dp[j] = dp[j+1] + triangle[i][j]
			}
		}
	}
	return dp[0]
}

/**
深度优先搜索算法
	搜索所有路径，并记录每条路径的长度，比较获取最小值
*/
func minimumTotal3(triangle [][]int) int {
	ret := make([]int, 0)
	ret = dfs(triangle, 0, 0, 0, ret)
	min := ret[0]
	for _, item := range ret {
		if item < min {
			min = item
		}
	}
	return min
}

func dfs(triangle [][]int, x, y, distance int, ret []int) []int {
	distance += triangle[x][y]
	if x == len(triangle)-1 {
		return append(ret, distance)
	}
	ret = dfs(triangle, x+1, y, distance, ret)
	ret = dfs(triangle, x+1, y+1, distance, ret)
	return ret
}

/**
深度优先搜索算法：
	只需要返回较短的那个路径即可，这个思路是动态规划的反向思考
 */
func minimumTotal4(triangle [][]int) int {
	return dfs4(triangle, 0, 0, 0)
}

func dfs4(triangle [][]int, x, y, distance int) int {
	distance += triangle[x][y]
	if x == len(triangle)-1 {
		return distance
	}
	l := dfs4(triangle, x+1, y, distance)
	r := dfs4(triangle, x+1, y+1, distance)
	if l < r {
		return l
	}
	return r
}
