package main

import "fmt"

/**
52. N皇后 II
	https://leetcode-cn.com/problems/n-queens-ii/
题目描述：
	n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
示例:
	输入: 4
	输出: 2
	解释: 4 皇后问题存在如下两个不同的解法。
	[
 		[".Q..",  // 解法 1
  		"...Q",
  		"Q...",
  		"..Q."],

 		["..Q.",  // 解法 2
  		"Q...",
  		"...Q",
  		".Q.."]
	]
*/
func main() {
	fmt.Println(totalNQueens(0))
	fmt.Println(totalNQueens2(0))

}

func totalNQueens(n int) int {
	cols, pie, na := make(map[int]bool, 0), make(map[int]bool, 0), make(map[int]bool, 0)
	return dfs(n, 0, 0, cols, pie, na)
}

func dfs(n, row, total int, cols, pie, na map[int]bool) int {
	if row == n {
		return total + 1
	}
	for j := 0; j < n; j++ {
		if ok, _ := cols[j]; ok {
			continue
		}
		if ok, _ := pie[row+j]; ok {
			continue
		}
		if ok, _ := na[row-j]; ok {
			continue
		}
		cols[j] = true
		pie[row+j] = true
		na[row-j] = true
		total = dfs(n, row+1, total, cols, pie, na)
		delete(cols, j)
		delete(pie, row+j)
		delete(na, row-j)
	}
	return total
}

/**
利用位运算（以4皇后为例说明）
	每次按行遍历
	col 记录列上的皇后位置，比如：0100，说明在第二列放置了一个皇后；
	pie 记录皇后在撇上的攻击信息，比如：1000，说明当前行不能在第一列放置皇后；当向下扫描下一行时，应该将当前的 pie 左移一位
	na  记录皇后在捺上的攻击信息，比如：0010，说明当前行不能在第三列放置皇后；当向下扫描下一行时，应该将当前的 na  右移一位

*/
func totalNQueens2(n int) int {
	return dfsBit(n, 0, 0, 0, 0, 0)
}

func dfsBit(n, row, col, pie, na, total int) int {
	if row >= n {
		total++
		return total
	}
	bits := (^(col | pie | na)) & (1<<n - 1) //获取在该行上所有可以放置皇后的bit位（为1的位置）
	for bits > 0 {
		bit := bits & (-bits) //获取可以放置皇后的最后一个bit位
		total = dfsBit(n, row+1, col|bit, (pie|bit)<<1, (na|bit)>>1, total)
		bits = bits & (bits - 1) //去除可以放置皇后的最后一个bit位
	}
	return total
}
