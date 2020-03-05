package main

import "fmt"

/**
合并排序的数组
https://leetcode-cn.com/problems/sorted-merge-lcci/
*/

func main() {
	//A := []int{1,2,3,0,0,0}
	//m := 3
	//B := []int{2,5,6}
	//n := 3
	A := []int{0}
	m := 0
	B := []int{1}
	n := 1
	merge4(A, m, B, n)
	fmt.Println(A)
}

//2. 直接合并后再进行一次排序

/**
3. 双指针做法（将两个数组看成队列）
	指针 a 指向 队列 A 的头部，指针 b 指向队列 B 的头部
*/
func merge3(A []int, m int, B []int, n int) {
	temp := make([]int, m+n)
	a := 0
	b := 0
	for a < m || b < n { //循环取两个队列的队首元素
		if a == m { //队列 A 已空，直接将队列 B 的队首元素放入
			temp[a+b] = B[b]
			b++
		} else if b == n { //队列 B 已空，直接将队列 A 的队首元素放入
			temp[a+b] = A[a]
			a++
		} else if A[a] < B[b] { //队列 A 的首部元素小于队列 B 的首部元素，将 A 的队首元素放入
			temp[a+b] = A[a]
			a++
		} else { //队列 B 的首部元素小于队列 A 的首部元素，将 B 的队首元素放入
			temp[a+b] = B[b]
			b++
		}
	}
	for index, value := range temp { //将最终结果拷贝到队列 A 中
		A[index] = value
	}
}

/**
1. 双指针做法，一个正向指针，一个逆向指针
*/
func merge(A []int, m int, B []int, n int) {
	for i := 0; i < n; i++ {
		j := m - 1 //m 记录的是数组 A 的元素个数
		for {
			if j < 0 || B[i] >= A[j] {
				A[j+1] = B[i]
				m++
				break
			}
			A[j+1] = A[j]
			j--
		}
	}
}

/**
4. 逆向双指针
*/
func merge4(A []int, m int, B []int, n int) {
	a := m - 1
	b := n - 1
	tail := m + n - 1
	for a >= 0 || b >= 0 {
		if a < 0 {
			A[tail] = B[b]
			b--
		} else if b < 0 {
			A[tail] = A[a]
			a--
		} else if A[a] > B[b] {
			A[tail] = A[a]
			a--
		} else {
			A[tail] = B[b]
			b--
		}
		tail--
	}
}
