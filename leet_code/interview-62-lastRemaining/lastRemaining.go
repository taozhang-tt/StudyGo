package main

import "fmt"

/**
面试题62. 圆圈中最后剩下的数字
	https://leetcode-cn.com/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/
题目描述
	0,1,,n-1这n个数字排成一个圆圈，从数字0开始，每次从这个圆圈里删除第m个数字。求出这个圆圈里剩下的最后一个数字。
	例如，0、1、2、3、4这5个数字组成一个圆圈，从数字0开始每次删除第3个数字，则删除的前4个数字依次是2、0、4、1，因此最后剩下的数字是3。
示例 1：
	输入: n = 5, m = 3
	输出: 3
示例 2：
	输入: n = 10, m = 17
	输出: 2
*/

func main() {
	fmt.Println(lastRemaining2(5, 3))
	fmt.Println(lastRemaining2(10, 17))
	fmt.Println(lastRemaining3(5, 3))
	fmt.Println(lastRemaining2(10, 17))
	fmt.Println(lastRemaining3(5, 1))
	fmt.Println(lastRemaining3(1, 2))
}

/**
(1)对于规模为 n 的问题：0, 1, 2, ..., (m-1), m, m+1, ..., n-2, n-1，第一个被删除的元素为 m-1
(2)删除掉 m-1 后，我们重新对元素做一个排序，下次将从第 m 个元素开始计数：m, m+1, ..., n-2, n-1, ..., 0, 1, ..., m-3, m-2; 这个时候问题规模变为了 n-1
那么问题就变成：m, m+1, ..., n-2, n-1, ..., 0, 1, ..., m-3, m-2 这 n-1 个元素排成一个圈，每次从这个圆圈里删除第m个数字；
那么在 n-1 规模问题下，它们对坐标被映射为
m, m+1, ..., n-2, 	  n-1,     ..., 0,   1,       ..., m-3, m-2
0, 1  , ..., n-(m+2), n-(m+1), ..., n-m, n-(m-1), ..., n-3, n-2
记 n 规模下坐标为 y， n-1 规模下坐标为 x，两者的坐标映射关系为 y = (x + m) % n
从而将规模为 n 的问题转化成了 规模为 n-1 的问题
*/

//利用上述思路使用递归
func lastRemaining(n int, m int) int {
	if n == 1 {
		return 0
	}
	x := lastRemaining(n-1, m)
	return (m + x) % n
}

//利用上述思路使用迭代
func lastRemaining2(n int, m int) int {
	f := 0
	for i := 2; i <= n; i++ {
		f = (f + m) % i
	}
	return f
}

//使用循环链表(超时)
func lastRemaining3(n int, m int) int {
	if m == 1 {
		return n-1
	}
	root := Genereate(n)
	for root.Next != root {
		for i := 0; i < m-1; i++ {
			root = root.Next
		}
		root.Next = root.Next.Next
	}
	return root.Val
}

func Genereate(n int) *ListNode {
	prev := &ListNode{
		Val: -1,
	}
	curr := prev
	for i := 0; i < n; i++ {
		curr.Next = &ListNode{
			Val: i,
		}
		curr = curr.Next
	}
	curr.Next = prev.Next
	return prev
}

type ListNode struct {
	Val  int
	Next *ListNode
}
