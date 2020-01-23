package main

import "fmt"

func BinarySearch(nums []int, value int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == value {
			return mid
		}
		if nums[mid] < value {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main()  {
	nums := []int{1, 2, 3, 4, 5, 6, 8, 9}
	fmt.Print(BinarySearch(nums, 9))
}