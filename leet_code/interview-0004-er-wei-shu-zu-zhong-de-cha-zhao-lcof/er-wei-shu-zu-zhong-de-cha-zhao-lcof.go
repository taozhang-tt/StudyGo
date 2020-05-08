package main

import "fmt"

func main()  {
	var matrix [][]int
	matrix = [][]int{
		[]int{1,   4,  7, 11, 15},
		[]int{2,   5,  8, 12, 19},
		[]int{3,   6,  9, 16, 22},
		[]int{10, 13, 14, 17, 24},
		[]int{18, 21, 23, 26, 30},
	}
	fmt.Println(findNumberIn2DArray(matrix, 5))
	fmt.Println(findNumberIn2DArray(matrix, 20))
}

/**
从右上角开始查询，如果当前元素等于 target 返回；如果当前元素大于 target 则消除当前列；如果当前元素小于 target 则消除当前行
*/
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	n, m := len(matrix), len(matrix[0])
	for i:=0; i<n; i++ {
		for j:=m-1; j >= 0; j-- {
			if matrix[i][j] == target {
				return true
			}
			if matrix[i][j] > target {	//消除当前列
				m -= 1
				j = m
			} else {	//消除当前行
				break
			}
		}
	}
	return false
}