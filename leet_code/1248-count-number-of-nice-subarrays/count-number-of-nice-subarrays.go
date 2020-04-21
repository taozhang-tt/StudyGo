package main

import "fmt"

/**
1248. 统计「优美子数组」
	https://leetcode-cn.com/problems/count-number-of-nice-subarrays/
题目描述：
	给你一个整数数组 nums 和一个整数 k。
	如果某个 连续 子数组中恰好有 k 个奇数数字，我们就认为这个子数组是「优美子数组」。
	请返回这个数组中「优美子数组」的数目。
示例 1：
	输入：nums = [1,1,2,1,1], k = 3
	输出：2
	解释：包含 3 个奇数的子数组是 [1,1,2,1] 和 [1,2,1,1] 。
示例 2：
	输入：nums = [2,4,6], k = 1
	输出：0
	解释：数列中不包含任何奇数，所以不存在优美子数组。
示例 3：
	输入：nums = [2,2,2,1,2,2,1,2,2,2], k = 2
	输出：16
*/

func main() {
	var (
		nums []int
		k    int
		ret  int
	)
	nums, k = []int{1, 1, 2, 1, 1}, 3
	ret = numberOfSubarrays2(nums, k)
	fmt.Println(ret)

	nums, k = []int{2, 4, 6}, 1
	ret = numberOfSubarrays2(nums, k)
	fmt.Println(ret)

	nums, k = []int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}, 2
	ret = numberOfSubarrays2(nums, k)
	fmt.Println(ret)

}

/**
双指针
	先找出一个包含 k 个奇数的子序列，且序列开头和结尾都是奇数
	判断 left 前面有多少个连续的偶数 lCnt
	判断 right 右边有多少个连续的偶数 rCnt
	包含该子序列的满足要求的子序列个数为 (lCnt+1) * (rCnt+1)
*/
func numberOfSubarrays(nums []int, k int) int {
	cnt, left, right, ret := 0, 0, 0, 0
	for right < len(nums) && left < len(nums) {
		lCnt, rCnt := 0, 0
		for left < len(nums) && nums[left]%2 == 0 {
			lCnt++
			left++
		}
		if left == len(nums) {
			return ret
		}
		if right == 0 {
			right = left
		}
		for right < len(nums) {
			if nums[right]%2 == 1 {
				cnt++
			}
			if cnt == k {
				break
			}
			right++
		}
		if right == len(nums) {
			return ret
		}
		right++
		for right < len(nums) && nums[right]%2 == 0 { //统计该子序列右侧有多少个偶数元素
			rCnt++
			right++
		}
		ret += (lCnt + 1) * (rCnt + 1)
		left++
		cnt--
	}
	return ret
}

/**
双指针（官方题解，好巧妙的双指针啊）
	思路如上
*/
func numberOfSubarrays2(nums []int, k int) int {
	n := len(nums)
	odd, ans, cnt := make([]int, n+2), 0, 0
	odd = append(odd, -1)
	for i := 0; i < n; i++ {
		if nums[i]%2 == 1 {
			cnt++
			odd[cnt] = i
		}
	}
	cnt++
	odd[0], odd[cnt] = -1, n
	for i := 1; i+k <= cnt; i++ {
		ans += (odd[i] - odd[i-1]) * (odd[i+k] - odd[i+k-1])
	}
	return ans
}
