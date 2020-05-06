package main

import "fmt"

func main() {
	var nums []int
	nums = []int{7, 5, 6, 4}
	fmt.Println(reversePairs(nums))
}

/**
暴力解法(超时)
	两层循环
*/
func reversePairs(nums []int) int {
	ret := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i+1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				ret++
			}
		}
	}
	return ret
}
