package main

import "fmt"

func main() {
	// var nums = []int{5, 2, 3, 1}
	// fmt.Println(sortArray(nums))
	// nums = []int{5,1,1,2,0,0}
	// fmt.Println(sortArray(nums))
	var nums = []int{-1, 2, -8, -10}
	fmt.Println(sortArray(nums))

}

func sortArray(nums []int) []int {
	selectSort(nums)
	return nums
}

//快排
func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	l, r, mid := left, right, nums[left]
	for l < r {
		for l < r && nums[r] > mid {
			r--
		}
		if l < r {
			nums[l] = nums[r]
			l++
		}
		for l < r && nums[l] < mid {
			l++
		}
		if l < r {
			nums[r] = nums[l]
			r--
		}
	}
	nums[l] = mid
	quickSort(nums, left, l-1)
	quickSort(nums, l+1, right)
}

//冒泡（超时喽）
func bubbleSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

//选择排序
func selectSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
}

//插入排序
func insertSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		temp := nums[i]
		var j = i - 1
		for j >= 0 && nums[j] > temp {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = temp
	}
}
