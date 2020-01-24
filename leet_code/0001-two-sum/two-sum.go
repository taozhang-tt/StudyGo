package main

import "fmt"

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

func main() {
	nums := []int{3, 4, 2}
	target := 6
	ret := twoSum2(nums, target)
	fmt.Println(ret)
}