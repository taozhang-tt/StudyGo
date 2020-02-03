package main

import "fmt"

func main() {
	//nums := []int{2,3,-2,4}
	nums := []int{-2, 0, -1}
	fmt.Println(maxProduct(nums))
	fmt.Println(maxProduct2(nums))

}

/**
动态规划解法
	1. 状态定义： dp[i]，表示包含第i个元素时，乘积的最大值；解释：dp[i] 是 0 - i 这些元素某一个连续子序列的乘积，可能是 [i]*[i-1]*[i-2]，可以明确的是一定包含 [i]
	2. 状态转换方程：
		如果只考虑所有的元素都大于 0：dp[i] = dp[i-1]*[i]；
		如果考虑包含小于0的元素：（1）当 [i] > 0时：dp[i] = dp[i-1]*[i]；（2）当 [i] < 0时：dp[i] = dp[i-1]的最小值*[i]
	所以状态定义 为 maxDp[i] 和 minDp[i]
*/
func maxProduct(nums []int) int {
	maxDp, minDp, max := make([]int, len(nums)), make([]int, len(nums)), 0
	maxDp[0], minDp[0], max = nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		m := nums[i] * maxDp[i-1]
		n := nums[i] * minDp[i-1]
		if m > n {
			maxDp[i] = m
			minDp[i] = n
		} else {
			maxDp[i] = n
			minDp[i] = m
		}
		if nums[i] > maxDp[i] {
			maxDp[i] = nums[i]
		}
		if nums[i] < minDp[i] {
			minDp[i] = nums[i]
		}
		if maxDp[i] > max {
			max = maxDp[i]
		}
	}
	return max
}

/**
动态规划优化
	dp 的状态只需要保存前一个位置的即可
 */
func maxProduct2(nums []int) int {
	maxDp, minDp, max := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		maxDp = nums[i] * maxDp
		minDp = nums[i] * minDp
		if maxDp < minDp {
			maxDp, minDp = minDp, maxDp
		}
		if nums[i] > maxDp {
			maxDp = nums[i]
		}
		if nums[i] < minDp {
			minDp = nums[i]
		}
		if maxDp > max {
			max = maxDp
		}
	}
	return max
}
