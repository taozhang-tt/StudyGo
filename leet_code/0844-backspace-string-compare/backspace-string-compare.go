package main

import "fmt"

/**
844. 比较含退格的字符串
	https://leetcode-cn.com/problems/backspace-string-compare/
题目描述：
	给定 S 和 T 两个字符串，当它们分别被输入到空白的文本编辑器后，判断二者是否相等，并返回结果。 # 代表退格字符。
	注意：如果对空文本输入退格字符，文本继续为空。
示例 1：
	输入：S = "ab#c", T = "ad#c"
	输出：true
	解释：S 和 T 都会变成 “ac”。
示例 2：
	输入：S = "ab##", T = "c#d#"
	输出：true
	解释：S 和 T 都会变成 “”。
示例 3：
	输入：S = "a##c", T = "#a#c"
	输出：true
	解释：S 和 T 都会变成 “c”。
示例 4：
	输入：S = "a#c", T = "b"
	输出：false
	解释：S 会变成 “c”，但 T 仍然是 “b”。
*/

func main() {
	var S, T string

	S = "xywrrmp"
	T = "xywrrmu#p"
	fmt.Println(backspaceCompare2(S, T))

	S = "ab#c"
	T = "ad#c"
	fmt.Println(backspaceCompare2(S, T))

	S = "ab##"
	T = "c#d#"
	fmt.Println(backspaceCompare2(S, T))

	S = "a##c"
	T = "#a#c"
	fmt.Println(backspaceCompare2(S, T))
}

/*
利用栈
statckS, stackT 分别记录两个字符串中的字符
遍历每个字符
if (当前字符 != '#') {
	当前字符入栈
} else {
	if (栈不空) {
		弹出栈顶元素
	}
}
*/
func backspaceCompare(S string, T string) bool {
	str1 := make([]byte, 0)
	for i := 0; i < len(S); i++ {
		if S[i] != '#' {
			str1 = append(str1, S[i])
		} else if len(str1) > 0 {
			str1 = str1[:len(str1)-1]
		}
	}
	str2 := make([]byte, 0)
	for i := 0; i < len(T); i++ {
		if T[i] != '#' {
			str2 = append(str2, T[i])
		} else if len(str2) > 0 {
			str2 = str2[:len(str2)-1]
		}
	}
	return string(str1) == string(str2)
}

/*
双指针法
	ps 指向 S 的最后一个字符，pt 指向 T 的最后一个字符，从后向前挪动指针
	skill 变量记录挪动到当前位置时 '#' 的数量
		if (当前位置的字符 == '#') {
			则 skill++；
		} else {
			if (skill > 0) {
				skill--;
			} else {
				break;
			}
		}
		向前挪动一个位置
*/
func backspaceCompare2(S string, T string) bool {
	ps, pt := len(S)-1, len(T)-1
	skillS, skillT := 0, 0
	for ps >= 0 || pt >= 0 {
		for ps >= 0 {
			if S[ps] == '#' {
				skillS++
			} else if skillS > 0 {
				skillS--
			} else {
				break
			}
			ps--
		}
		for pt >= 0 {
			if T[pt] == '#' {
				skillT++
			} else if skillT > 0 {
				skillT--
			} else {
				break
			}
			pt--
		}
		if pt >= 0 {
			if ps < 0 || T[pt] != S[ps] {
				return false
			}
		}
		if ps >= 0 {
			if pt < 0 || T[pt] != S[ps] {
				return false
			}
		}
		ps--
		pt--
	}
	return true
}
