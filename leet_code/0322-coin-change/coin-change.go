package main

import "fmt"

/**
322. 零钱兑换
	https://leetcode-cn.com/problems/coin-change/
题目描述：
	给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
示例 1:
	输入: coins = [1, 2, 5], amount = 11
	输出: 3
	解释: 11 = 5 + 5 + 1
*/

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
		if coins[j] <= amount { //考虑 零钱面值大于 amount 的情况
			dp[coins[j]] = 1
		}
	}
	for i := 1; i <= amount; i++ {
		if dp[i] == 0 { //没有对应的硬币可以直接兑换
			min := -2
			for j := 0; j < len(coins); j++ { //状态转换方程的主体部分
				if i > coins[j] && dp[i-coins[j]] > 0 && (min == -2 || dp[i-coins[j]] < min) {
					min = dp[i-coins[j]]
				}
			}
			dp[i] = min + 1
		}
	}
	return dp[amount]
}
