package main

import "fmt"

func main() {
	nums := []int{10,9,2,5,3,7,101,18}
	fmt.Println(lengthOfLIS(nums))
}

/**
动态规划解题思路：
	状态定义：dp[i]，包含元素 i 在内的最长上升序列，i 取值 0 到 n-1
	状态转换方程：dp[i] = max(dp[j]) + 1；j 取值 0 到 i-1，同时要满足 a[j]<a[i]
 */
func lengthOfLIS(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	ret := 1
	dp := make([]int, len(nums))
	dp[0] = 1
	for i:=1; i<len(nums); i++ {
		max := 0
		for j:=0; j<i; j++ {
			if nums[j] < nums[i] && dp[j] > max {
				max = dp[j]
			}
		}
		dp[i] = max + 1
		if dp[i] > ret {
			ret = dp[i]
		}
	}
	return ret
}
