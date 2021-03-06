package main

import "fmt"

/**
14. 最长公共前缀
	https://leetcode-cn.com/problems/longest-common-prefix/
题目描述
	编写一个函数来查找字符串数组中的最长公共前缀。
	如果不存在公共前缀，返回空字符串 ""。
示例 1:
	输入: ["flower","flow","flight"]
	输出: "fl"
示例 2:
	输入: ["dog","racecar","car"]
	输出: ""
	解释: 输入不存在公共前缀。
说明:
	所有输入只包含小写字母 a-z 。
*/
func main() {
	var strs []string
	strs = []string{
		"flower", "flow", "flight",
	}
	fmt.Println(longestCommonPrefix(strs))
	strs = []string{
		"dog","racecar","car",
	}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[0][i] != strs[j][i] {
				return string(strs[0][:i])
			}
		}
	}
	return strs[0]
}
