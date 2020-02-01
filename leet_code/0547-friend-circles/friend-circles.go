package main

import "fmt"

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
