package main

import "fmt"

/**
5. 最长回文子串
	https://leetcode-cn.com/problems/longest-palindromic-substring/
题目描述
	给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
示例 1：
	输入: "babad"
	输出: "bab"
	注意: "aba" 也是一个有效答案。
示例 2：
	输入: "cbbd"
	输出: "bb"
*/
func main()  {
	var s string
	s = "babad"
	fmt.Println(longestPalindrome(s))
	s = "cbbd"
	fmt.Println(longestPalindrome(s))
	s = "bab"
	fmt.Println(longestPalindrome(s))
}

/**
暴力法、枚举法
	枚举每个位置 i，由该位置向两侧扩散
*/
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	max, L, R := 0, 0, 0
	for i:=0; i<len(s)-1; i++ {
		l, r, cnt := i-1, i+1, 1
		for l >=0 && r < len(s) {
			if s[l] == s[r] {
				cnt += 2
				l--
				r++
			} else {
				break
			}
		}
		if cnt > max {
			max, L, R = cnt, l, r
		}
		l, r, cnt = i, i+1, 0
		for l >=0 && r < len(s) {
			if s[l] == s[r] {
				cnt += 2
				l--
				r++
			} else {
				break
			}
		}
		if cnt > max {
			max, L, R = cnt, l, r
		}
	}
	return string(s[L+1:R])
}