package main

import "fmt"

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
