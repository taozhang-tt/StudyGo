package main

import "fmt"

/**
15. 三数之和
	https://leetcode-cn.com/problems/3sum/
题目描述
	给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
	注意：答案中不可以包含重复的三元组。
示例：
	给定数组 nums = [-1, 0, 1, 2, -1, -4]，
	满足要求的三元组集合为：
	[
	[-1, 0, 1],
	[-1, -1, 2]
	]
*/

func main() {
	var nums = []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}

/**
排序加双指针
	1. 如果元素个数小于3，直接返回无结果
	2. 排序
	3. 遍历排序后到数组元素索引记作 i，如果 nums[i] > 0 ，因为数组已排序，后面不可能出现小于 0 的元素，跳过，遍历下个元素
	4. 如果当前元素 nums[i] 和前面的一个元素 nums[i-1] 相等，可以直接跳过；因为以该元素开头的解必然在遍历上一个元素时已被找出
	5. 设置两个指针 l, r；分别指向下一个元素和最后一个元素，求这三个元素的和sum与0做比较
		如果 sum > 0，则左移 r 指针使 sum 变小
		如果 sum < 0，则右移 l 指针使 sum 变大
		如果 sum == 0，进一步判断以 i, l, r 组成的元组是否已经包含在结果中，未包含则放入，然后左移 r 指针，右移 l 指针
*/
func threeSum(nums []int) [][]int {
	ret := make([][]int, 0)
	if len(nums) < 3 {
		return ret
	}
	quickSort(nums, 0, len(nums)-1)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1 //双指针
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum < 0 {
				l++
			}
			if sum > 0 {
				r--
			}
			if sum == 0 {
				if r == len(nums)-1 || nums[r] != nums[r+1] {
					ret = append(ret, []int{nums[i], nums[l], nums[r]})
				}
				l++
				r--
			}
		}
	}
	return ret
}

func quickSort(nums []int, L, R int) {
	if L >= R {
		return
	}
	var l, r, mid = L, R, nums[L]
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
	quickSort(nums, L, l-1)
	quickSort(nums, l+1, R)
}
