package main

import "fmt"

/**
680. 验证回文字符串 Ⅱ
	https://leetcode-cn.com/problems/valid-palindrome-ii/
题目描述
	给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。
示例 1:
	输入: "aba"
	输出: True
示例 2:
	输入: "abca"
	输出: True
	解释: 你可以删除c字符。
注意:
	字符串只包含从 a-z 的小写字母。字符串的最大长度是50000。
*/

func main() {
	var s string
	s = "aba"
	fmt.Println(validPalindrome(s))
	s = "abca"
	fmt.Println(validPalindrome(s))
}

/**
双指针
*/
func validPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] == s[r] {	//字符相等进行下一个判断
			l++
			r--
			continue
		}
		return isPalindrome(s, l+1, r) || isPalindrome(s, l, r-1)	//字符不等时分两种情况，去除当前左侧字符 or 去除当前右侧字符
	}
	return true
}

func isPalindrome(s string, l, r int) bool {	//判断字符串 s 中，l～r 部分是否回文
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
