package main

import "fmt"

/**
409. 最长回文串
	https://leetcode-cn.com/problems/longest-palindrome/
题目描述
	给定一个包含大写字母和小写字母的字符串，找到通过这些字母构造成的最长的回文串。
	在构造过程中，请注意区分大小写。比如 "Aa" 不能当做一个回文字符串。
	注意:
	假设字符串的长度不会超过 1010。
示例 1:
	输入:
	"abccccdd"
	输出:
	7
	解释:
	我们可以构造的最长的回文串是"dccaccd", 它的长度是 7。
 */
func main()  {
	s := "abccccdd"
	fmt.Println(longestPalindrome(s))
}

/**
遍历一遍，记录每个字符出现的次数，结果 = 所有出现过偶数次的字符的偶数次和 + 1
 */
func longestPalindrome(s string) int {
	sMap := make(map[int32]int)
	count := 0
	for _, char := range s {
		sMap[char] ++
		if sMap[char] == 2 {
			count += 2
			sMap[char] = 0
		}
	}
	if count < len(s) {
		count ++
	}
	return count
}