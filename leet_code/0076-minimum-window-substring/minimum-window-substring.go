package main

/**
76. 最小覆盖子串
	https://leetcode-cn.com/problems/minimum-window-substring/
题目描述
	给你一个字符串 S、一个字符串 T，请在字符串 S 里面找出：包含 T 所有字符的最小子串。
示例：
	输入: S = "ADOBECODEBANC", T = "ABC"
	输出: "BANC"
说明：
	如果 S 中不存这样的子串，则返回空字符串 ""。
	如果 S 中存在这样的子串，我们保证它是唯一的答案。
*/
func main()  {
	
}

/**
滑动窗口、双指针
	
*/
func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	tMap := make(map[byte]int)
	for i:=0; i<len(t); i++ {
		tMap[t[i]] += 1
	}
	ret, l, r, sMap := s, 0, 0, make(map[byte]int)
	for l <= r {
		if check(sMap, tMap) {

		}
	}
	
}

func check(sMap, tMap map[byte]int) bool {
	for ch, num := range tMap {
		if sMap[ch] != num {
			return false
		}
	}
	return true
}