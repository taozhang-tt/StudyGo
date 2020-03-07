package main

import "fmt"

/**
32. 最长有效括号
	https://leetcode-cn.com/problems/longest-valid-parentheses/
题目描述：
	给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。
示例 1:
	输入: "(()"
	输出: 2
	解释: 最长有效括号子串为 "()"
示例 2:
	输入: ")()())"
	输出: 4
	解释: 最长有效括号子串为 "()()"
*/

func main() {
	s := "()(()"
	fmt.Println(longestValidParentheses2(s))
	s = "(()"
	fmt.Println(longestValidParentheses2(s))
	s = ")()())"
	fmt.Println(longestValidParentheses2(s))
}

/**
动态规划解法：
	状态定义：dp[i]，包含 a[i]在内，最长有效括号长度，当 a[j] 与 a[i] 匹配时（i>j）为了标记 a[j] 已被匹配，会赋值 dp[j] = dp[i]
	状态转换：
		（1）a[i] == '('
			dp[i] = 0
*/
func longestValidParentheses(s string) int {
	dp := make([]int, len(s))
	max := 0
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			for j := i - 1; j >= 0; j-- {
				if s[j] == '(' && dp[j] == 0 {
					dp[j] = dp[i-1] + 2
					if j >= 1 {
						dp[j] += dp[j-1]
					}
					dp[i] = dp[j]
					break
				}
			}
			if dp[i] > max {
				max = dp[i]
			}
		}
	}
	return max
}

func longestValidParentheses2(s string) int {
	if len(s) <= 0 {
		return 0
	}
	dp := make([]int, len(s))
	sign := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		dp[i] = dp[i-1]
		if s[i] == ')' {
			dp[i] = dp[i-1]
			for j := i - 1; j >= 0; j-- {
				if s[j] == '(' && sign[j] == 0 {
					dp[i] += 2
					sign[j] = 1
					break
				}
			}
		}
	}
	return dp[len(s)-1]
}
