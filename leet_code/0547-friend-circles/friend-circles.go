package main

import "fmt"

/**
547. 朋友圈
	https://leetcode-cn.com/problems/friend-circles/
题目描述：
	班上有 N 名学生。其中有些人是朋友，有些则不是。他们的友谊具有是传递性。如果已知 A 是 B 的朋友，B 是 C 的朋友，那么我们可以认为 A 也是 C 的朋友。所谓的朋友圈，是指所有朋友的集合。
	给定一个 N * N 的矩阵 M，表示班级中学生之间的朋友关系。如果M[i][j] = 1，表示已知第 i 个和 j 个学生互为朋友关系，否则为不知道。你必须输出所有学生中的已知的朋友圈总数。
示例 1:
	输入:
	[[1,1,0],
	 [1,1,0],
	 [0,0,1]]
	输出: 2
	说明：已知学生0和学生1互为朋友，他们在一个朋友圈。
	第2个学生自己在一个朋友圈。所以返回2。
*/

func main() {
	M := [][]int{
		[]int{1, 1, 0},
		[]int{1, 1, 0},
		[]int{0, 0, 1},
	}
	fmt.Println(findCircleNum(M))

	M = [][]int{
		[]int{1, 1, 0},
		[]int{1, 1, 1},
		[]int{0, 1, 1},
	}
	fmt.Println(findCircleNum(M))

	M = [][]int{
		[]int{1, 0, 0, 1},
		[]int{0, 1, 1, 0},
		[]int{0, 1, 1, 1},
		[]int{1, 0, 1, 1},
	}
	fmt.Println(findCircleNum(M))
}

/*
使用颜色法
*/
func findCircleNum(M [][]int) int {
	ret := 0
	for x := 0; x < len(M); x++ {
		for y := 0; y < len(M[0]); y++ {
			if M[x][y] == 1 {
				dfsFloodFill(M, x, y)
				ret++
			}
		}
	}
	return ret
}

func dfsFloodFill(M [][]int, x, y int) {
	M[x][y] = 0
	//扫描x，y 不变
	for xx := 0; xx < len(M); xx++ {
		if M[xx][y] == 1 {
			dfsFloodFill(M, xx, y)
		}
	}
	//扫描y，x 不变
	for yy := 0; yy < len(M[0]); yy++ {
		if M[x][yy] == 1 {
			dfsFloodFill(M, x, yy)
		}
	}
}
