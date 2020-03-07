package main

import "fmt"

/**
169. 多数元素
	https://leetcode-cn.com/problems/majority-element/
题目描述：
	给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
	你可以假设数组是非空的，并且给定的数组总是存在多数元素。
示例 1:
	输入: [3,2,3]
	输出: 3
*/

func main() {
	nums := []int{6, 5, 5}
	fmt.Println(majorityElement5(nums))
}

//暴力解法，超时
func majorityElement(nums []int) int {
	maxCount := 0
	maxCountEle := nums[0]
	for i := 0; i < len(nums); i++ {
		count := 0
		for j := 0; j < len(nums); j++ {
			if nums[j] == nums[i] {
				count++
			}
		}
		if count > (len(nums)-1)/2 {
			return nums[i]
		}
		if count > maxCount {
			maxCount = count
			maxCountEle = nums[i]
		}
	}
	return maxCountEle
}

//通过map计数
func majorityElement2(nums []int) int {
	m := make(map[int]int, len(nums))
	for _, item := range nums {
		m[item] ++
		if m[item] > (len(nums)-1)/2 {
			return item
		}
	}
	return -1
}

//先排序，再统计
func majorityElement3(nums []int) int {
	quickSort(nums, 0, len(nums)-1)
	count := 1
	ele := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			ele = nums[i]
			count = 1
		} else {
			count++
			if count > (len(nums)-1)/2 {
				return ele
			}
		}
	}
	return ele
}

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	mid := nums[left]
	l := left
	r := right
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

/**
leetcode go项目中的一个解法
题目条件有 count > n/2，那即是说，该元素的数量多于其它所有元素总数量，当我们遍历所有元素时：
如果是目标元素，则 count + 1
如果不是目标元素，则 count - 1
使用ret 记录 count 不为 0 时的元素
遍历结束时，count 必大于 0 且此时 ret 记录的即为目标元素
*/
func majorityElement4(nums []int) int {
	ret, count := nums[0], 0
	for _, item := range nums {
		if count == 0 {
			ret, count = item, 1
		} else {
			if ret == item {
				count++
			} else {
				count--
			}
		}
	}
	return ret
}

/**
分治法
*/
func majorityElement5(nums []int) int {
	return find(nums, 0, len(nums)-1)
}
func find(nums []int, left, right int) int {
	if left == right { //只有一个元素，则为目标众数
		return nums[left]
	}
	mid := (left + right) / 2 //将问题拆解成两个子问题
	lRet := find(nums, left, mid)
	rRet := find(nums, mid+1, right)
	if lRet == rRet { //如果在两个子问题中查找到的众数一致，则该数即为众数
		return lRet
	}
	lCount, rCount := 0, 0 //如果两个子问题中查找到的众数不一致，则分别统计两个众数出现的次数，返回较多的那个
	for i := left; i <= mid; i++ {
		if nums[i] == lRet {
			lCount++
		}
	}
	for j := mid + 1; j <= right; j++ {
		if nums[j] == rRet {
			rCount++
		}
	}
	if lCount > rCount {
		return lRet
	}
	return rRet
}
