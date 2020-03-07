package main

import "fmt"

/**
231. 2的幂
	https://leetcode-cn.com/problems/power-of-two/
题目描述：
	给定一个整数，编写一个函数来判断它是否是 2 的幂次方。
*/

func main() {
	fmt.Println(isPowerOfTwo2(1))
	fmt.Println(isPowerOfTwo2(16))
	fmt.Println(isPowerOfTwo2(218))
}

//直接每次对 2 取余，如果余数大于0 则返回false
func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	if n == 1 {
		return true
	}
	for n > 1 {
		if n%2 != 0 {
			return false
		}
		n /= 2
	}
	return true
}

// 如果一个数是 2 对幂次方，那么它对二进制位里面有且仅有1位为1，其余均为0
func isPowerOfTwo2(n int) bool {
	if n <= 0 {
		return false
	}
	temp := uint32(n) //加上这一行，速度快了很多
	if temp&(temp-1) == 0 {
		return true
	}
	return false
}
