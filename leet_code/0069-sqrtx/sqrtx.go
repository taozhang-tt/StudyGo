package main

import "fmt"

/**
69. x 的平方根
	https://leetcode-cn.com/problems/sqrtx/
题目描述：
	实现 int sqrt(int x) 函数。
	计算并返回 x 的平方根，其中 x 是非负整数。
	由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
示例:
	输入: 4
	输出: 2
*/

func main() {
	fmt.Println(mySqrt(2))
	fmt.Println(mySqrt2(4, 1e-4))
}

//leetcode 题解
func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	l, r, res := 1, x, 1
	for l <= r {
		m := (l + r) / 2
		if m == x/m {
			return m
		} else if m < x/m {
			l = m + 1
		} else {
			r = m - 1
			res = r
		}
	}
	return res
}

//加强版，根据精度计算
func mySqrt2(x int, precision float64) float64 {
	l, r := 0.0, float64(x)
	for l < r {
		m := (l + r) / 2
		if m == float64(x)/m {
			return m
		} else if m < float64(x)/m {
			l = m
		} else {
			r = m
		}
		if (r - l) < precision {
			break
		}
	}
	return l
}
