package main

import "fmt"

func main() {
	fmt.Println(totalNQueens(2))
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
