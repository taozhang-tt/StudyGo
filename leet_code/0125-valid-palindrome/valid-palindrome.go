package main

import "fmt"

/**
125. 验证回文串
	https://leetcode-cn.com/problems/valid-palindrome/
题目描述
	给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
说明：本题中，我们将空字符串定义为有效的回文串。

示例 1:
	输入: "A man, a plan, a canal: Panama"
	输出: true
示例 2:
	输入: "race a car"
	输出: false
*/
func main() {
	var s string
	s = "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))

	s = "race a ca"
	fmt.Println(isPalindrome(s))
}

/**
双指针
*/
func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if !isWordOrNum(s[l]) {
			l++
			continue
		}
		if !isWordOrNum((s[r])) {
			r--
			continue
		}
		if lowerCase(s[l]) != lowerCase(s[r]) {
			return false
		}
		l++
		r--
	}
	return true
}

//字符转小写
func lowerCase(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		b += 'a' - 'A'
	}
	return b
}

func isWordOrNum(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9')
}