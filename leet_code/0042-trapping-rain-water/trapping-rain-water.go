package main

import "fmt"

/**
42. 接雨水
	https://leetcode-cn.com/problems/trapping-rain-water/
题目描述：
	给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
示例:
	输入: [0,1,0,2,1,0,1,3,2,1,2,1]
	输出: 6
*/

func main() {
	var height = []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height))
	fmt.Println(trap2(height))
	fmt.Println(trap3(height))
	fmt.Println(trap4(height))
}

/**
暴力做法：一层一层去扫扫描
	每一层高度记为level，扫描每一层有多少个小于 level 的元素
*/
func trap(height []int) int {
	ret, level := 0, -1
	for {
		level++
		left, right := 0, 0
		for i := 0; i < len(height); i++ { //从左侧开始向左扫，找到这一层第一个大于 level 的位置
			if height[i] > level {
				left = i
				break
			}
		}
		for i := len(height) - 1; i > 0; i-- { //从右侧开始向左扫，找到这一层最后一个大于 level 的位置
			if height[i] > level {
				right = i
				break
			}
		}
		if left == right { //left 等于 right 说明这一层已经没有可以盛雨水的地方，结束层级扫描
			break
		}
		for i := left; i <= right; i++ { //判断这一层有多少个地方可以盛雨水
			if height[i] < level+1 {
				ret++
			}
		}
	}
	return ret
}

/**
暴力做法
	对于任一位置，从当前位置开始向左扫描找一个最大值 maxL，从当前位置向右扫描找一个最大值 maxR，取 maxL、maxR 的最小值减去当前位置的值，即为当前位置能盛的雨水
*/
func trap2(height []int) int {
	ret, size := 0, len(height)
	for i := 1; i < len(height)-1; i++ {
		maxL, maxR := 0, 0
		for j := i; j >= 0; j-- {
			if height[j] > maxL {
				maxL = height[j]
			}
		}
		for j := i; j < size; j++ {
			if height[j] > maxR {
				maxR = height[j]
			}
		}
		minMax := maxL
		if minMax > maxR {
			minMax = maxR
		}
		ret += minMax - height[i]
	}
	return ret
}

/**
动态编程
	先存储每个位置对应的左侧最大值和右侧最大值，分别用两个数组记录 maxL, maxR，然后遍历每个位置 i, min(maxL[i], maxR[i]) - height[i] 的值即为当前位置可以接的雨水
	对于 maxL[i]，其值应该为 max(maxL[i-1], height[i])
	对于 maxR[i]，其值应该为 max(maxR[i+1], height[i])
*/
func trap3(height []int) int {
	size, ret := len(height), 0
	if size == 0 {
		return ret
	}
	maxL, maxR := make([]int, size), make([]int, size)
	maxL[0] = height[0]
	for i := 1; i < size; i++ {
		if maxL[i-1] > height[i] {
			maxL[i] = maxL[i-1]
		} else {
			maxL[i] = height[i]
		}
	}
	maxR[size-1] = height[size-1]
	for i := size - 2; i >= 0; i-- {
		if maxR[i+1] > height[i] {
			maxR[i] = maxR[i+1]
		} else {
			maxR[i] = height[i]
		}
	}
	for i := 1; i < size-1; i++ {
		minMax := maxL[i]
		if maxR[i] < minMax {
			minMax = maxR[i]
		}
		ret += minMax - height[i]
	}
	return ret
}

/**
栈的应用
	我们在遍历数组时维护一个栈。
	如果当前的条形块小于或等于栈顶的条形块，我们将条形块的索引入栈，意思是当前的条形块被栈中的前一个条形块界定。
	如果我们发现一个条形块长于栈顶，我们可以确定栈顶的条形块被当前条形块和栈的前一个条形块界定，因此我们可以弹出栈顶元素并且累加答案到 ans 。
*/
func trap4(height []int) int {
	curr, stack, size, ret := 0, make([]int, 0), len(height), 0
	if size == 0 {
		return ret
	}
	for curr < size {	//遍历数组元素
		for len(stack) > 0 && height[curr] > height[stack[len(stack)-1]] {	//当前元素大于栈顶元素
			top := stack[len(stack)-1]	//栈顶元素
			stack = stack[0 : len(stack)-1]	//栈顶元素出栈
			if len(stack) == 0 {
				break
			}
			dis := curr - stack[len(stack)-1] - 1	//考虑出现连续顶元素相等的情况
			bounded_height := height[curr]	//计算高度差
			if height[stack[len(stack)-1]] < bounded_height {
				bounded_height = height[stack[len(stack)-1]]
			}
			ret += dis * (bounded_height - height[top])	//将能够盛的雨水累加到结果
		}
		stack = append(stack, curr)
		curr++
	}
	return ret
}
