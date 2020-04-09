package main

import "fmt"

/**
17. 电话号码的字母组合
	https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/
题目描述：
	给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
	给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
示例:
	输入："23"
	输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
*/

func main() {
	var (
		digits string
		ret    []string
	)
	digits = "23"
	ret = letterCombinations(digits)
	fmt.Println(ret)
}

var m []string = []string{
	"abc",
	"def",
	"ghi",
	"jkl",
	"mno",
	"pqrs",
	"tuv",
	"wxyz",
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	ret := make([]string, 0)
	dfs(digits, "", 0, &ret)
	return ret
}

func dfs(digits, curr string, index int, ret *[]string) {
	if index == len(digits) {
		*ret = append(*ret, curr)
		return
	}
	for _, char := range m[digits[index]-'2'] {
		dfs(digits, curr+string(char), index+1, ret)
	}
}
