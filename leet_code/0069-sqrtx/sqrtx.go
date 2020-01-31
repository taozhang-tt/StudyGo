package main

import "fmt"

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
		m := (l+r)/2
		if m == float64(x)/m {
			return m
		} else if m < float64(x)/m {
			l = m
		} else {
			r = m
		}
		if (r-l)<precision {
			break
		}
	}
	return l
}
