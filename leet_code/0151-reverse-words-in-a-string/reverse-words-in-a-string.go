package main

import "fmt"

func main() {
	var s, ret string
	s = "the sky is blue"
	ret = reverseWords2(s)
	fmt.Println(ret)
	s = "  hello world!  "
	ret = reverseWords2(s)
	fmt.Println(ret)
	s = "a good   example"
	ret = reverseWords2(s)
	fmt.Println(ret)
	s = ""
	ret = reverseWords2(s)
	fmt.Println(ret)
	s = " "
	ret = reverseWords2(s)
	fmt.Println(ret)
}

/**
递归做法
*/
func reverseWords(s string) string {
	return recursion(s, 0)
}

func recursion(s string, start int) string {
	curr := ""
	if start == len(s) {
		return curr
	}
	for start < len(s) && s[start] == ' ' {
		start++
	}
	for start < len(s) && s[start] != ' ' {
		curr += string(s[start])
		start++
	}
	next := recursion(s, start)
	if next != "" {
		next += " "
	}
	return next + curr
}

/**

 */
func reverseWords2(s string) string {
	ret := ""
	word := ""
	for _, char := range s {
		if char == ' ' {
			if word != "" {
				ret = word + " " + ret
				word = ""
			}
		} else {
			word += string(char)
		}
	}
	if word != "" {
		ret = word + " " + ret
	}
	if len(ret) == 0 {
		return ret
	}else {
		return string(ret[:len(ret)-1])
	}
}
