package main

import "fmt"

/**
51. N皇后
	https://leetcode-cn.com/problems/n-queens/
题目描述：
	n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
	给定一个整数 n，返回所有不同的 n 皇后问题的解决方案。
	每一种解法包含一个明确的 n 皇后问题的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
示例:
	输入: 4
	输出: [
 		[".Q..",  // 解法 1
  		"...Q",
  		"Q...",
  		"..Q."],

 		["..Q.",  // 解法 2
 		 "Q...",
  		"...Q",
  		".Q.."]
	]
	解释: 4 皇后问题存在两个不同的解法。
*/

func main() {
	fmt.Println(solveNQueens(0))
}

/**
N 皇后问题，对不能攻击的定义：
  1. 一行只能有一个皇后
  2. 一列只能有一个皇后
  3. 从右上到左下（简称撇）的对角线上只能有一个皇后
  4. 从左上到右下（简称捺）的对角线上只能有一个皇后
题解：
  我们以矩阵来描述棋盘，以 i 来记录行号，i 从 0 到 n-1；以 j 来记录列号，j 从 0 到 n-1
  1. 任意两个皇后的 i 和 j 均不能相同，这一点很明显
  2. 两个皇后不能处在同一撇上，这个撇通过分析棋盘发现可以表示为 i + j == 常数 c；比如 （0，1）和（1，0）处在同一撇上，两者的 i + j 均等于 1；再比如（0，2）、（1，1）、（2，0）
  3. 两个皇后不能处在同一捺上，这个捺通过分析棋盘发现可以表示为 j - j == 常数 c; 比如 （0，2）和（1，3）处在同一捺上，两者的 i - j 均等于 -2；再比如（0，1）、（1，2），（2，3）
所以我们在递归搜索的时候，可以通过上述三个条件进行剪枝，避免无答案的递归迭代
*/
func solveNQueens(n int) [][]string {
	var (
		cols, pie, na = make(map[int]bool, 0), make(map[int]bool, 0), make(map[int]bool, 0)
		status        = make(map[int]int, 0)
		ret           = make([]map[int]int, 0)
	)
	ret = dfs(n, 0, cols, pie, na, status, ret)
	return outPut(ret)
}

/**
n : 问题的规模
row ：行号
cols：已经被放置皇后的列号
pie：已经存在皇后的撇
na：已经存在皇后的捺
status：当前的状态 （key为i，value为j）
*/
func dfs(n, row int, cols, pie, na map[int]bool, status map[int]int, ret []map[int]int) []map[int]int {
	if row == n {
		item := make(map[int]int, 0)
		for key, value := range status {
			item[key] = value
		}
		ret = append(ret, item)
		return ret
	}
	for j := 0; j < n; j++ {
		if ok, _ := cols[j]; ok { //列冲突
			continue
		}
		if ok, _ := pie[row+j]; ok { //撇冲突
			continue
		}
		if ok, _ := na[row-j]; ok { //捺冲突
			continue
		}
		status[row] = j   //记录皇后位置
		cols[j] = true    //记录列位置
		pie[row+j] = true //记录撇
		na[row-j] = true  //记录捺
		ret = dfs(n, row+1, cols, pie, na, status, ret)
		delete(status, row)
		delete(cols, j)
		delete(pie, row+j)
		delete(na, row-j)
	}
	return ret
}

//查找map中是否存在value
func outPut(m []map[int]int) [][]string {
	var ret [][]string
	for _, item := range m {
		answer := make([]string, 0)
		length := len(item)
		for j := 0; j < length; j++ {
			str := ""
			for i := 0; i < length; i++ {
				if i == item[j] {
					str += "Q"
				} else {
					str += "."
				}
			}
			answer = append(answer, str)
		}
		ret = append(ret, answer)
	}
	return ret
}
