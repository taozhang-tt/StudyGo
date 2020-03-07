package main

import "fmt"

/*
1. 两数之和
	https://leetcode-cn.com/problems/two-sum/
题目描述：
	给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
	你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
示例:
	给定 nums = [2, 7, 11, 15], target = 9
	因为 nums[0] + nums[1] = 2 + 7 = 9
	所以返回 [0, 1]
*/

func main() {
	nums := []int{3, 4, 2}
	target := 6
	ret := twoSum2(nums, target)
	fmt.Println(ret)
}

//暴力求解
func twoSum(nums []int, target int) []int {
	ret := make([]int, 0, 2)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				ret = append(ret, i, j)
				return ret
			}
		}
	}
	return ret
}

//使用 map 解法
func twoSum2(nums []int, target int) []int {
	m := make(map[int]int, 0)
	ret := make([]int, 0, 2)
	for index, value := range nums {
		m[value] = index
	}
	for i, v := range nums {
		if index, exist := m[target-v]; exist && i != index {
			ret = append(ret, i, index)
			return ret
		}
	}
	return ret
}
