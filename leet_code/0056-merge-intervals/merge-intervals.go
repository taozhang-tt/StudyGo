package main

import "fmt"

/**
56. 合并区间
	https://leetcode-cn.com/problems/merge-intervals/
题目描述：
	给出一个区间的集合，请合并所有重叠的区间。
示例 1:
	输入: [[1,3],[2,6],[8,10],[15,18]]
	输出: [[1,6],[8,10],[15,18]]
	解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2:
	输入: [[1,4],[4,5]]
	输出: [[1,5]]
	解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
*/

func main() {
	var intervals, ret [][]int
	intervals = [][]int{
		[]int{1, 3},
		[]int{8, 10},
		[]int{2, 6},
		[]int{15, 18},
	}
	ret = merge2(intervals)
	fmt.Println(ret)

	intervals = [][]int{
		[]int{1, 4},
		[]int{4, 5},
	}
	ret = merge2(intervals)
	fmt.Println(ret)
}

/**
自己写的
和官方题解一样的思路
*/
func merge(intervals [][]int) [][]int {
	ret := make([][]int, 0)
	if len(intervals) <= 0 {
		return ret
	}
	quickSort(intervals, 0, len(intervals)-1)
	left, right := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		l, r := intervals[i][0], intervals[i][1]
		if l > right {
			ret = append(ret, []int{left, right})
			left, right = intervals[i][0], intervals[i][1]
		} else if r > right {
			right = r
		}
	}
	ret = append(ret, []int{left, right})
	return ret
}

/**
官方题解写法
*/
func merge2(intervals [][]int) [][]int {
	ret := make([][]int, 0)
	if len(intervals) <= 0 {
		return ret
	}
	quickSort(intervals, 0, len(intervals)-1)
	for _, item := range intervals {
		l, r := item[0], item[1]
		if len(ret) == 0 || ret[len(ret)-1][1] < l {
			ret = append(ret, item)
		} else if ret[len(ret)-1][1] < r {
			ret[len(ret)-1][1] = r
		}
	}
	return ret
}

func quickSort(intervals [][]int, left, right int) {
	if left >= right {
		return
	}
	mid, l, r := intervals[left], left, right
	for l < r {
		for l < r && intervals[r][0] > mid[0] {
			r--
		}
		if l < r {
			intervals[l] = intervals[r]
			l++
		}
		for l < r && intervals[l][0] < mid[0] {
			l++
		}
		if l < r {
			intervals[r] = intervals[l]
			r--
		}
	}
	intervals[l] = mid
	quickSort(intervals, left, l-1)
	quickSort(intervals, l+1, right)
}
