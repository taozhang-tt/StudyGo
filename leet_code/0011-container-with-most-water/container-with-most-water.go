package main

import "fmt"

/**
11. 盛最多水的容器
	https://leetcode-cn.com/problems/container-with-most-water/
题目描述
	给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
	说明：你不能倾斜容器，且 n 的值至少为 2。
示例：
	输入：[1,8,6,2,5,4,8,3,7]
	输出：49
*/

func main() {
	var height []int
	var ret int
	height = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	ret = maxArea2(height)
	fmt.Println(ret)

	height = []int{0, 0}
	ret = maxArea2(height)
	fmt.Println(ret)
}

/**
双指针（我的思路）
	固定一个左边界，使用一个指针从最后一个边界向前扫，找到当前左边界的最大容积
*/
func maxArea(height []int) int {
	ret := 0
	for i := 0; i < len(height)-1; i++ {
		j := len(height) - 1
		for i < j {
			h := height[j]
			if height[i] < height[j] {
				h = height[i]
			}
			if h*(j-i) > ret {
				ret = h * (j - i)
			}
			if height[i] < height[j] {
				break
			}
			j--
		}
	}
	return ret
}

/**
双指针（官方思路）
*/
func maxArea2(height []int) int {
	ret, l, r := 0, 0, len(height)-1
	for l < r {
		h := height[j]	//取两个边界的较小值
		if height[i] < height[j] {
			h = height[i]
		}
		if h*(j-i) > ret {	//判断当前边界组成的容器容积与当前最大容积的关系
			ret = h * (j - i)
		}
		if height[l] < height[r] {	//如果左边界较小，左边界右移
			l++
		} else {	//如果右边界较小，右边界左移
			r--
		}
	}
	return ret
}
