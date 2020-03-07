package main

import "fmt"

/**
70. 爬楼梯
	https://leetcode-cn.com/problems/climbing-stairs/
题目描述：
	假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
	每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
	注意：给定 n 是一个正整数。
示例 1：
	输入： 2
	输出： 2
解释： 有两种方法可以爬到楼顶。
	1.  1 阶 + 1 阶
	2.  2 阶
*/

func main() {
	fmt.Println(climbStairs4(44))
}

/**
回溯做法：从下往上算 (超时)
	f(0) = f(1) + f(2)
	从第0个台阶可以跳一步到第1个台阶上，也可以跳两个台阶到第2个台阶上，那么从第0个台阶跳到第n个台阶的方式 = 从第1个跳到第n个 + 从第2个跳到第n个
*/
func climbStairs(n int) int {
	if n > 1 {
		return F(n, 1) + F(n, 2)
	}
	return 1
}

func F(n, curr int) int {
	if n-curr < 2 {
		return 1
	}
	return F(n, curr+1) + F(n, curr+2)
}

/**
回溯做法 的加速版
*/
func climbStairs2(n int) int {
	m := make([]int, n+1, n+1)
	if n > 1 {
		return F2(n, 1, m) + F2(n, 2, m)
	}
	return 1
}

func F2(n, curr int, m []int) int {
	if n-curr < 2 {
		return 1
	}
	if m[curr+1] == 0 {
		m[curr+1] = F2(n, curr+1, m)
	}
	if m[curr+2] == 0 {
		m[curr+2] = F2(n, curr+2, m)
	}
	return m[curr+1] + m[curr+2]
}

/**
回溯做法：从上往下算
	要想跳到第n阶，那可以从第 n-1 阶跳一步上来 或者是从第 n-2 阶跳两步上来，那么则有 跳到第 n 阶的方式 f(n) = f(n-1) + f(n-2)
*/

func climbStairs3(n int) int {
	m := make([]int, n+1, n+1)
	return F3(n, m)
}

func F3(n int, m []int) int {
	if n < 2 {
		return 1
	}
	if m[n-1] == 0 {
		m[n-1] = F3(n-1, m)
	}
	if m[n-2] == 0 {
		m[n-2] = F3(n-2, m)
	}
	return m[n-1] + m[n-2]
}

/*
DP做法：
	递推公式 f(n) = f(n-1) + f(n-2)
*/

func climbStairs4(n int) int {
	f0, f1 := 1, 1
	for i := 2; i <= n; i++ {
		f1, f0 = f1+f0, f1
	}
	return f1
}
