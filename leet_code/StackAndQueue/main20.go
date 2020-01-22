package main

import (
	"fmt"
)

func main() {
	s := "[]"
	fmt.Println(isValid2(s))
}

//使用map保存括号的对应关系，逻辑清晰，但是有额外的空间开销
func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	var m = map[rune]rune{
		']': '[',
		'}': '{',
		')': '(',
	}
	stack := make([]rune, 0)
	for _, item := range s {
		if v, ok := m[item]; ok {
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] != v {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, item)
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}

//不实用map记录对应关系，直接逻辑中判断
func isValid2(s string) bool {
	if len(s) == 0 {
		return true
	}
	stack := make([]rune, 0)
	for _, item := range s {
		if item == '{' || item == '(' || item == '[' {
			stack = append(stack, item)
		} else {
			if len(stack) == 0 || (item == ']' && stack[len(stack)-1] != '[') || (item == '}' && stack[len(stack)-1] != '{') || (item == ')' && stack[len(stack)-1] != '('){
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}