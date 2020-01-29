package main

import "fmt"

func main() {
	var x = 2.0
	var n = -2
	fmt.Println(myPow3(x, n))
}

//暴力求解，超时无法通过
func myPow(x float64, n int) float64 {
	ret := 1.0
	if n < 0 {
		x = 1/x
		n = -n
	}
	for i:=0; i<n; i++ {
		ret *= x
	}
	return ret
}

/**
分治法：递归做法
 */
func myPow2(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	if n < 0 {
		return 1/myPow2(x, -n)
	}
	if n % 2 > 0 {	// n为奇数
		return myPow2(x, n-1) * x
	}
	return myPow2(x*x, n/2)
}

/**
分治法：非递归做法
*/
func myPow3(x float64, n int) float64 {
	if n < 0 {
		x = 1/x
		n = -n
	}
	ret := 1.0
	for n > 0 {
		if n % 2 > 0 {
			ret *= x
		}
		x = x*x
		n = n / 2
	}
	return ret
}