package main

import "fmt"

/**
22. 括号生成
	https://leetcode-cn.com/problems/generate-parentheses/
题目描述：
	给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。
	例如，给出 n = 3，生成结果为：
	[
  	"((()))",
  	"(()())",
  	"(())()",
  	"()(())",
  	"()()()"
	]
*/

func main() {
	fmt.Println(generateParenthesis2(4))
}

/**
利用递归搜索的思想
	n 对括号，一共有 2n 个字符，相当于往 2n 个盒子里塞东西，塞的过程是有要求的：
	如果 左括号数量 == 右括号数量 == n，说明是一个结果，保存下来
	必须满足 右括号的数量 < 左括号的数量 < n

*/
func generateParenthesis(n int) []string {
	ret := make([]string, 0)
	ret = generate(n, 0, 0, "", ret)
	return ret
}

func generate(n, leftUsed, rightUsed int, curr string, ret []string) []string {
	if leftUsed == n && rightUsed == n { //左括号数量 == 右括号数量 == n，说明是一个结果，保存下来
		return append(ret, curr)
	}
	if leftUsed < n { //左括号是可以随时塞入的，只需要满足 左括号的数量 < n 即可
		ret = generate(n, leftUsed+1, rightUsed, curr+"(", ret)
	}
	if rightUsed < leftUsed && rightUsed < n { //如果想塞入右括号，需要满足 右括号数量 < 左括号数量 && 右括号数量 < n
		ret = generate(n, leftUsed, rightUsed+1, curr+")", ret)
	}
	return ret
}


func generateParenthesis2(n int) []string {
	ret := make([]string, 0)
	dfs(n, 0, 0, "", &ret)
	return ret
}

func dfs(n, leftUsed, rightUsed int, curr string, ret *[]string) {
	if leftUsed == n && rightUsed == n { //左括号数量 == 右括号数量 == n，说明是一个结果，保存下来
		*ret = append(*ret, curr)
		return
	}
	if leftUsed < n { //左括号是可以随时塞入的，只需要满足 左括号的数量 < n 即可
		dfs(n, leftUsed+1, rightUsed, curr+"(", ret)
	}
	if rightUsed < leftUsed && rightUsed < n { //如果想塞入右括号，需要满足 右括号数量 < 左括号数量 && 右括号数量 < n
		dfs(n, leftUsed, rightUsed+1, curr+")", ret)
	}
}
