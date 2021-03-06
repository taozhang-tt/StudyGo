package main

import (
	"fmt"
	"strings"
)

//3. 无重复字符的最长子串
//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

func main() {
	var str = "abcabcbb"
	num := lengthOfLongestSubstring(str)
	fmt.Println(num)
	num2 := lengthOfLongestSubstring2(str)
	fmt.Println(num2)
}

//窗口滑动法，左闭右开
func lengthOfLongestSubstring(s string) int {
	var (
		left   = 0	//窗口左边界
		right  = 0	//窗口右边界
		max    = 0	//窗口最大宽度
		sSlice = strings.Split(s, "")
		sMap   = make(map[string]int)
		length = len(sSlice)
	)
	for ; left < length && right < length; {
		charRight := sSlice[right]	//取右边界字符
		if _, exist := sMap[charRight]; exist {	//如果右边界字符已经存在于窗口中
			charLeft := sSlice[left]	//向右移动左边界，右边界保持不动
			delete(sMap, charLeft)
			left++
		} else {	//如果窗口中并未包含右边界对应的字符
			sMap[charRight] = right	//将该字符存入窗口
			right++	//窗口边界右移
			if distance := right - left; distance > max {	//判断当前窗口，记录较大值
				max = distance
			}
		}
	}
	return max
}

//窗口滑动法优化，左闭右开
func lengthOfLongestSubstring2(s string) int {
	var (
		left   = 0
		right  = 0
		max    = 0
		sSlice = strings.Split(s, "")
		sMap   = make(map[string]int) //用于检查字符是否已经存在及存在的位置
		length = len(sSlice)
	)
	//创建一个滑动窗口，left 为左边界，right 为右边界（约定为左闭右开）, 右边界开始向右滑动
	for ; right < length; {
		charRight := sSlice[right]
		//如果已经存在，移动左边界到已存在字符的下一个字符
		if index, exist := sMap[charRight]; exist {
			//防止左边界左移
			if index >= left {
				left = index + 1
			}
		}
		sMap[charRight] = right
		right++
		if distance := right - left; distance > max {
			max = distance
		}
	}
	return max
}
