package main

import "fmt"

/**
面试题57 - II. 和为s的连续正数序列
	https://leetcode-cn.com/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof/
题目描述：
	输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。
	序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。
示例 1：
	输入：target = 9
	输出：[[2,3,4],[4,5]]
*/

func main() {
	fmt.Println(findContinuousSequence2(2))
	fmt.Println(findContinuousSequence2(3))
	fmt.Println(findContinuousSequence2(9))
	fmt.Println(findContinuousSequence2(15))
}

/**
暴力解法
	暴力枚举，由于题目要求至少有两个正整数，所以可知枚举的上限为 (target+1)/2
*/
func findContinuousSequence(target int) [][]int {
	limit := (target + 1) / 2
	ret := make([][]int, 0)
	for i := 1; i < limit; i++ {
		sum := i
		item := []int{i}
		j := i + 1
		for sum < target && j <= limit {
			sum += j
			item = append(item, j)
			if sum == target {
				ret = append(ret, item)
				break
			}
			j++
		}
	}
	return ret
}

/**
双指针做法 || 滑动窗口
	指针 l 指向集合的头，指针 r 指向指针的尾部，[l,r] 这个集合所有元素的和 Sn = (l+r)*(r-l+1)/2
	初始时 l = 1, r = 2
	当 Sn < target 时，r 向右移动一位
	当 Sn > target 时，l 向右移动一位
	当 Sn ==target 时，成功找到一个结果，下一步 l 向右移动一步（因为以l开头满足条件的只可能有一个结果）
	r 的上限为 (target + 1)/2
*/
func findContinuousSequence2(target int) [][]int {
	limit := (target + 1) / 2
	ret := make([][]int, 0)
	l := 1
	r := 2
	for l < r && r <= limit {
		sn := (l + r) * (r - l + 1) / 2
		if sn == target {
			item := make([]int, 0, r-l+1)
			for i := l; i <= r; i++ {
				item = append(item, i)
			}
			ret = append(ret, item)
			l++
		} else if sn < target {
			r++
		} else {
			l++
		}
	}
	return ret
}
