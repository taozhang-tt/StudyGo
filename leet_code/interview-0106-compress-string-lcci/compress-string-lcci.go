package main

import (
	"fmt"
	"strconv"
)

/**
面试题 01.06. 字符串压缩
	https://leetcode-cn.com/problems/compress-string-lcci/
题目描述：
	字符串压缩。利用字符重复出现的次数，编写一种方法，实现基本的字符串压缩功能。比如，字符串aabcccccaaa会变为a2b1c5a3。若“压缩”后的字符串没有变短，则返回原先的字符串。你可以假设字符串中只包含大小写英文字母（a至z）
 */

func main() {
	var S = "aabcccccaaa"
	fmt.Println(compressString(S))
}

func compressString(S string) string {
	if len(S) == 0 {
		return S
	}
	var ret string
	var count = 1
	var curr = S[0]
	for i:=1; i<len(S); i++ {
		char := S[i]
		if char != curr {
			ret += string(curr) + strconv.Itoa(count)
			curr = char
			count = 1
			continue
		}
		count++
	}
	ret += string(curr) + strconv.Itoa(count)
	if len(ret) >= len(S) {
		return S
	}
	return ret
}
