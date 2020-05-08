package main

import "fmt"

func main()  {
	var nums []int
	nums = []int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println(maxSubArray(nums))
}

/**
动态规划
	dp[i]: 到第 i 个元素为止，最大和（可能包含第 i 个，也可能不包含）
状态转换
	dp[i] = max(dp[i-1]+nums[i], nums[i])
*/
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	ret := nums[0]
	for i:=1; i<len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		ret = max(ret, dp[i])
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
优化动态规划
	不需要重开数组保存dp状态
*/
func maxSubArray2(nums []int) int {
	ret := nums[0]
	for i:=1; i<len(nums); i++ {
		if nums[i] + nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > ret {
			ret = nums[i]
		}
	}
	return ret
}