package main

import "fmt"
import ""

//利用双端队列，一个可用但是效率不高的做法
func maxSlidingWindow(nums []int, k int) []int {
	var ret = make([]int, 0)
	if len(nums) == 0 {
		return ret
	}
	var deque = nums[0:k]
	var max = nums[0]
	var maxIndex = 0
	max, maxIndex = findMaxAndIndex(deque)
	deque = deque[maxIndex:]
	ret = append(ret, max)
	for i := k; i < len(nums); i++ {
		if len(deque) == k {
			deque = deque[1:]
			deque = append(deque, nums[i])
			max, maxIndex = findMaxAndIndex(deque)
		} else {
			if nums[i] < max {
				deque = append(deque, nums[i])
			} else {
				max = nums[i]
				deque = []int{max}
			}
		}
		ret = append(ret, max)
	}
	return ret
}

func findMaxAndIndex(nums []int) (max int, maxIndex int) {
	max = nums[0]
	for index, value := range nums {
		if value > max {
			max = value
			maxIndex = index
		}
	}
	return
}

func main() {
	nums := []int{1, -1}
	k := 1
	ret := maxSlidingWindow3(nums, k)
	fmt.Println(ret)
}

func maxSlidingWindow2(nums []int, k int) []int {
	if len(nums) == 0 || len(nums) < k {
		return make([]int, 0)
	}
	window := make([]int, 0, k) // store the index of nums
	result := make([]int, 0, len(nums)-k+1)
	for i, v := range nums { // if the left-most index is out of window, remove it
		if i >= k && window[0] <= i-k {
			window = window[1:len(window)]
		}
		for len(window) > 0 && nums[window[len(window)-1]] < v { // maintain window
			window = window[0 : len(window)-1]
		}
		window = append(window, i) // store the index of nums
		if i >= k-1 {
			result = append(result, nums[window[0]]) // the left-most is the index of max value in nums
		}
	}
	return result
}

//利用双端队列，借鉴了极客时间
func maxSlidingWindow3(nums []int, k int) []int {
	if len(nums) == 0 || len(nums) < k {
		return make([]int, 0)
	}
	var (
		ret    = make([]int, 0, len(nums)-k+1) //用户存储要返回的结果
		window = make([]int, 0, k)             //用于存储窗口内的元素的下标，这些下标对应的元素在 window 里是有序的，越靠近左侧，元素越大
	)
	for index, value := range nums {
		if index >= k && window[0] <= index-k { //窗口滑动时，判断窗口的最左侧保存的下标是否已经滑动到窗口外
			window = window[1:]
		}
		for len(window) > 0 && nums[window[len(window)-1]] < value {
			window = window[0 : len(window)-1]
		}
		window = append(window, index)
		if index >= k-1 {
			ret = append(ret, nums[window[0]])
		}
	}
	return ret
}
