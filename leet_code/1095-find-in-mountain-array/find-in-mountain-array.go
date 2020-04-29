package main

import "fmt"

/**
1095. 山脉数组中查找目标值
	https://leetcode-cn.com/problems/find-in-mountain-array/
题目描述:
	（这是一个 交互式问题 ）
	给你一个 山脉数组 mountainArr，请你返回能够使得 mountainArr.get(index) 等于 target 最小 的下标 index 值。
	如果不存在这样的下标 index，就请返回 -1。
	何为山脉数组？如果数组 A 是一个山脉数组的话，那它满足如下条件：
	首先，A.length >= 3
	其次，在 0 < i < A.length - 1 条件下，存在 i 使得：
	A[0] < A[1] < ... A[i-1] < A[i]
	A[i] > A[i+1] > ... > A[A.length - 1]
	你将 不能直接访问该山脉数组，必须通过 MountainArray 接口来获取数据：
	MountainArray.get(k) - 会返回数组中索引为k 的元素（下标从 0 开始）
	MountainArray.length() - 会返回该数组的长度
注意：
	对 MountainArray.get 发起超过 100 次调用的提交将被视为错误答案。此外，任何试图规避判题系统的解决方案都将会导致比赛资格被取消。
	为了帮助大家更好地理解交互式问题，我们准备了一个样例 “答案”：https://leetcode-cn.com/playground/RKhe3ave，请注意这 不是一个正确答案。
示例 1：
	输入：array = [1,2,3,4,5,3,1], target = 3
	输出：2
	解释：3 在数组中出现了两次，下标分别为 2 和 5，我们返回最小的下标 2。
示例 2：
	输入：array = [0,1,2,4,2,1], target = 3
	输出：-1
	解释：3 在数组中没有出现，返回 -1。
 

提示：
	3 <= mountain_arr.length() <= 10000
	0 <= target <= 10^9
	0 <= mountain_arr.get(index) <= 10^9
*/

func main() {
	var mountainArr *MountainArray
	var ret int
	mountainArr = &MountainArray{
		Nums: []int{1, 2, 3, 4, 5, 3, 1},
	}
	ret = findInMountainArray(2, mountainArr)
	fmt.Println(ret)
}

// This is the MountainArray's API interface.
// You should not implement it, or speculate about its implementation
type MountainArray struct {
	Nums []int
}

func (this *MountainArray) get(index int) int {
	return this.Nums[index]
}
func (this *MountainArray) length() int {
	return len(this.Nums)
}

/**
二分查找
	1. 先通过二分查找法获取最高点位置
		对于区间 [l, r]，取其中点 mid，如果 mid的值 < mid+1的值，说明已进入递增区间，则最大值应该落在区间 [mid+1, r]，否则最大值落在 [l, mid]
	2. 以最高点将区间分为左右两部分，分别使用二分法查找指定值

*/
func findInMountainArray(target int, mountainArr *MountainArray) int {
	l, r := 0, mountainArr.length()-1
	for l < r {	//通过一次二分查找获取最高点
		mid := (l+r)/2
		if mountainArr.get(mid) < mountainArr.get(mid+1) {	//处在递增区间
			l = mid+1
		} else {	//处于递减区间
			r = mid
		}
	}
	peek := l

	l, r = 0, peek	//对左半部分递增的区间进行二分查找
	for l <= r {
		mid := (l+r)/2
		value := mountainArr.get(mid)
		if value == target {
			return mid
		} 
		if value < target {
			l = mid+1
		} else {
			r = mid-1
		}
	}

	l, r = peek, mountainArr.length()-1	//对左右部分递减的区间进行二分查找
	for l <= r {
		mid := (l+r)/2
		value := mountainArr.get(mid)
		if value == target {
			return mid
		}
		if value < target {
			r = mid-1
		} else {
			l = mid+1
		}
	}
	return -1
}
