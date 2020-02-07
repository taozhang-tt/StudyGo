package main

import "fmt"

func main() {
	coins := []int{474, 83, 404, 3}
	fmt.Println(coinChange(coins, 249))
}

/**
动态规划解题思路：
	状态定义：dp[i]，兑换 i 元所需要的最小硬币数
	状态转换方程：dp[i] = min(dp[i-coins[j]) + 1，i 表示要兑换的金额，j 表示第 j 种面额的金币
*/
func coinChange(coins []int, amount int) int {
	if amount <= 0 {
		return 0
	}
	dp := make([]int, amount+1)
	//初始化
	for j := 0; j < len(coins); j++ {
		if coins[j] <= amount {	//考虑 零钱面值大于 amount 的情况
			dp[coins[j]] = 1
		}
	}
	for i := 1; i <= amount; i++ {
		if dp[i] == 0 { //没有对应的硬币可以直接兑换
			min := -2
			for j := 0; j < len(coins); j++ {	//状态转换方程的主体部分
				if i > coins[j] && dp[i-coins[j]] > 0 && (min == -2 || dp[i-coins[j]] < min) {
					min = dp[i-coins[j]]
				}
			}
			dp[i] = min + 1
		}
	}
	return dp[amount]
}
