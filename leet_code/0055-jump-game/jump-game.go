package main

import "fmt"

/**
55. 跳跃游戏
	https://leetcode-cn.com/problems/jump-game/
题目描述
	给定一个非负整数数组，你最初位于数组的第一个位置。
	数组中的每个元素代表你在该位置可以跳跃的最大长度。
	判断你是否能够到达最后一个位置。
示例 1:
	输入: [2,3,1,1,4]
	输出: true
	解释: 我们可以先跳 1 步，从位置 0 到达 位置 1, 然后再从位置 1 跳 3 步到达最后一个位置。
示例 2:
	输入: [3,2,1,0,4]
	输出: false
	解释: 无论怎样，你总会到达索引为 3 的位置。但该位置的最大跳跃长度是 0 ， 所以你永远不可能到达最后一个位置。
*/

func main() {
	var nums []int
	nums = []int{2, 3, 1, 1, 4}
	fmt.Println(canJump(nums))

	nums = []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(nums))

	nums = []int{2, 3, 1, 1, 4}
	fmt.Println(canJump2(nums))

	nums = []int{3, 2, 1, 0, 4}
	fmt.Println(canJump2(nums))
}

/**
动态规划（我也不知道算不算动态规划）
	对于最后一个元素，他肯定是可达自己的

*/
func canJump(nums []int) bool {
	dp := make([]bool, len(nums))
	dp[len(nums)-1] = true
	for i := len(nums) - 2; i >= 0; i-- {
		dp[i] = false
		for j := 1; j <= nums[i]; j++ {
			if i+j > len(nums)-1 {
				break
			}
			if dp[i+j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[0]
}

/**
贪心算法
	1. 从位置 0 起跳，记录能跳到的最远距离 max，则 0 ～ max 都是可以到达的位置
	2. 遍历 1 ～ max，计算每个位置起跳能到达的最远距离 dis，如果 dis > max ，更新 max = dis
	3. 如果 max >= 数组长度，说明可以到达最后一个元素
*/
func canJump2(nums []int) bool {
	dis := 0
	for index, value := range nums {
		if index <= dis && index + value > dis{
			dis = index + value
		}
		if dis >= len(nums) - 1 {
			return true
		}
	}
	return false
}
