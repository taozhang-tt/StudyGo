package main

import "fmt"

func main() {
	var arr []int
	arr = []int{1,2,2,1,1,3}
	uniqueOccurrences(arr)
}

func uniqueOccurrences(arr []int) bool {
	nums := make(map[int]int)
	for _, num := range arr {
		nums[num] ++
	}
	fmt.Println(nums)
	return true
}