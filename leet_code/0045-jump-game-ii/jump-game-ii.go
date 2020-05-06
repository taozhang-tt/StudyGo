package main

import "fmt"

/**
45. 跳跃游戏 II
	https://leetcode-cn.com/problems/jump-game-ii/
题目描述：
	给定一个非负整数数组，你最初位于数组的第一个位置。
	数组中的每个元素代表你在该位置可以跳跃的最大长度。
	你的目标是使用最少的跳跃次数到达数组的最后一个位置。
示例:
	输入: [2,3,1,1,4]
	输出: 2
解释: 跳到最后一个位置的最小跳跃数是 2。
	 从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
*/

func main() {
	var nums []int

	nums = []int{2}
	fmt.Println(jump3(nums))

	nums = []int{2, 1}
	fmt.Println(jump3(nums))

	nums = []int{2, 3, 0, 1, 4}
	fmt.Println(jump3(nums))

	nums = []int{2, 1, 1, 1, 4}
	fmt.Println(jump3(nums))
}

/**
深度优先搜索（超时）
	对于任意元素 i，它到达最后一个元素的最小距离 = 所有它能直接到到达的位置 i + 1、i + 2 ...，这些位置到达最后一个元素的距离 + 1
*/
func jump(nums []int) int {
	ret := dfs(nums, 0, 0)
	return ret
}

func dfs(nums []int, curr, dis int) int {
	if curr >= len(nums)-1 {
		return dis
	}
	min := len(nums) - 1
	for i := 1; i <= nums[curr]; i++ {
		d := dfs(nums, curr+i, dis+1)
		if d < min {
			min = d
		}
	}
	return min
}

/**
动态规划做法
	状态定义：dp[i]，表示从位置 i 到达最后一个元素所需的最少步数
	递推方程：dp[i] = min(dp[i+1], dp[i+2]...)，所有他能够直接到达的位置的最小 dp 值，或者是他可以直接到达最后一个元素，则 dp[i] = 1
*/

func jump2(nums []int) int {
	dp := make([]int, len(nums))
	for i := len(nums) - 2; i >= 0; i-- {
		dp[i] = len(nums) - 1
		for j := 1; j <= nums[i]; j++ {
			if i+j >= len(nums)-1 {
				dp[i] = 1
				break
			} else {
				if dp[i+j]+1 < dp[i] {
					dp[i] = dp[i+j] + 1
				}
			}
		}
	}
	return dp[0]
}

/**
层次遍历（不确定这个说法是否正确）
	1. 从位置 0 开始是第 0 层
	2. 第一跳能到达的位置是第 1 层，能到达的最远位置记录为 end
	3. 如果 end >= 最后一个位置，说明在该层就能到达最后一个元素，返回层数即为结果
	4. 否则
 */
func jump3(nums []int) int {
	ret, start, end := 0, 0, 0
	for end < len(nums)-1 {
		max := 0
		for i := start; i <= end; i++ {
			if i+nums[i] > max {
				max = i + nums[i]
			}
		}
		ret++
		start, end = end+1, max
	}
	return ret
}