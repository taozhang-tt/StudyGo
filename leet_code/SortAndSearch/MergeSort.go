package main

import "fmt"

func main() {
	var nums []int
	nums = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	mergeSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
/**

*/
func mergeSort(nums []int, L, R int) {
	if L < R {
		mid := (L + R) / 2
		mergeSort(nums, L, mid)
		mergeSort(nums, mid+1, R)
		merge(nums, L, mid, R)
	}
}

func merge(nums []int, L, M, R int) {
	temp, l, r, t := make([]int, R-L+1), L, M+1, 0
	for l <= M && r <= R {
		if nums[l] < nums[r] {
			temp[t] = nums[l]
			l++
		} else {
			temp[t] = nums[r]
			r++
		}
		t++
	}
	for l <= M {
		temp[t] = nums[l]
		t++
		l++
	}
	for r <= R {
		temp[t] = nums[r]
		r++
		t++
	}
	t = 0
	for L <= R {
		nums[L] = temp[t]
		t++
		L++
	}
}
