package main

import "fmt"

func main() {
	// var matrix = [][]int{
	// 	[]int{1, 2, 3},
	// 	[]int{4, 5, 6},
	// 	[]int{7, 8, 9},
	// }
	// rotate(matrix)
	// fmt.Println(matrix)

	var matrix = [][]int{
		[]int{5, 1, 9, 11},
		[]int{2, 4, 8, 10},
		[]int{13, 3, 6, 7},
		[]int{15, 14, 12, 16},
	}
	rotate2(matrix)
	fmt.Println(matrix)
}

var dx = []int{0, 1, 0, -1} //右，下，左，上
var dy = []int{1, 0, -1, 0}

func rotate(matrix [][]int) {
	recursive(matrix, 0, 0, 0, len(matrix)-1, 0, len(matrix[0])-1, 0, len(matrix)-1)
}

func recursive(matrix [][]int, x, y, xl, xr, yl, yr, direction, step int) {
	if xl > xr || yl > yr {
		return
	}
	xy, xx, yy, dir := matrix[x][y], x, y, direction
	for i := 0; i < step; i++ { //顺时针走 step 步
		xx += dx[dir]
		yy += dy[dir]
		if xx < xl || xx > xr || yy < yl || yy > yr { //越过边界
			xx -= dx[dir]
			yy -= dy[dir]
			dir++
			dir = dir % 4
			xx += dx[dir]
			yy += dy[dir]
		}
	}

	xn, yn, sign := x+dx[direction], y+dy[direction], false
	if xn < xl { //越过上边界，下一步向右
		sign = true
	}
	if xn > xr { //越过下边界，下一步向左
		sign = true
	}
	if yn < yl { //左边界越界，下一步向上
		sign = true
	}
	if yn > yr { //右边界越界，下一步向下
		sign = true
	}
	if sign {
		xn, yn = xn-dx[direction], yn-dy[direction]
		direction++
		if direction == 4 {
			yl++
			yr--
			xr--
			xl++
			direction = direction % 4
			step -= 2
		}
		xn, yn = xn+dx[direction], yn+dy[direction]
	}
	recursive(matrix, xn, yn, xl, xr, yl, yr, direction, step)
	matrix[xx][yy] = xy
}

func rotate2(matrix [][]int) {
	var ret, l = make([][]int, len(matrix)), len(matrix)
	for i := 0; i < l; i++ {
		ret[i] = make([]int, l)
		copy(ret[i], matrix[i])
	}
	fmt.Println(ret)
	for x := 0; x < l; x++ {
		for y := 0; y < l; y++ {
			matrix[y][l-x-1] = ret[x][y]
		}
	}
}
