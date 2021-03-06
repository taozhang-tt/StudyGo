package main

import "fmt"

/**
560. 和为K的子数组
	https://leetcode-cn.com/problems/subarray-sum-equals-k/
题目描述
	给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的个数。
示例 1 :
	输入:nums = [1,1,1], k = 2
	输出: 2 , [1,1] 与 [1,1] 为两种不同的情况。
	说明 :
提示：
	数组的长度为 [1, 20,000]。
	数组中元素的范围是 [-1000, 1000] ，且整数 k 的范围是 [-1e7, 1e7]。
*/
func main() {
	var nums []int
	var k int
	nums = []int{1, 1, 1}
	k = 2
	fmt.Println(subarraySum(nums, k))
}

/**
枚举法
	逆向思维使用枚举法，对于以 end 结尾的连续子数组，我们依次枚举 0 ~ end 每一个元素，枚举的元素记录为 start，如果 [start, end] 的和为k
	则找到一个满足的子数组
*/
func subarraySum(nums []int, k int) int {
	cnt := 0
	for end := 0; end < len(nums); end++ {
		sum := 0
		for start := end; start >= 0; start-- {
			sum += nums[start]
			if sum == k {
				cnt++
			}
		}
	}
	return cnt
}

/**
hash + 前缀和
	sum(j) 表示 [0, j] 这些元素的和
	那么，如果 [i, j] 的和为 k，则有：sum(j) - sum(i-1) = k; 从而有 sum(i-1) = sum(j) - k
	我们要求以 j 结尾满足和为 k 的子数组有多少个，只需要求有多少个满足 sum(i-1) = sum(j) - k 且 i < j
	我们不妨使用一个hash map来保存，以和为键，出现次数为对应的值
*/
func subarraySum1(nums []int, k int) int {
	m, cnt, sum := make(map[int]int, 0), 0, 0
	m[0] = 1	//第一个元素本身等于 k
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if num, exist := m[sum-k]; exist {
			cnt += num
		}
		m[sum] += 1
	}
	return cnt
}
