package main

import "fmt"

/**
面试题 17.16. 按摩师
	https://leetcode-cn.com/problems/the-masseuse-lcci/
题目描述
	一个有名的按摩师会收到源源不断的预约请求，每个预约都可以选择接或不接。
	在每次预约服务之间要有休息时间，因此她不能接受相邻的预约。
	给定一个预约请求序列，替按摩师找到最优的预约集合（总预约时间最长），返回总的分钟数。
示例 1：
	输入： [1,2,3,1]
	输出： 4
	解释： 选择 1 号预约和 3 号预约，总时长 = 1 + 3 = 4。
*/

func main() {
	var nums = []int{1, 2, 3, 1}
	fmt.Println(massage(nums))
	nums = []int{2, 7, 9, 3, 1}
	fmt.Println(massage(nums))
	nums = []int{2, 1, 4, 5, 3, 1, 1, 3}
	fmt.Println(massage(nums))
}

/**
动态规划解法
	状态定义：dp[i][0]：第i天不接受预约的情况下，最长总预约时间, dp[i][1]：第i天接受预约的情况下，最长总预约时间
	状态递推：dp[i][0] = max(dp[i-1][0], dp[i-1][1]), dp[i][1] = max(dp[i-1][1], dp[i-1][0]+a[i])
*/
func massage(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([][]int, len(nums))
	for i, v := range nums {
		dp[i] = make([]int, 2)
		if i == 0 {
			dp[i][0] = 0
			dp[i][1] = v
			continue
		}
		dp[i][0] = dp[i-1][0]
		if dp[i-1][1] > dp[i][0] {
			dp[i][0] = dp[i-1][1]
		}
		dp[i][1] = dp[i-1][1]
		if dp[i-1][0]+v > dp[i][1] {
			dp[i][1] = dp[i-1][0] + v
		}
	}
	if dp[len(nums)-1][0] > dp[len(nums)-1][1] {
		return dp[len(nums)-1][0]
	}
	return dp[len(nums)-1][1]
}
