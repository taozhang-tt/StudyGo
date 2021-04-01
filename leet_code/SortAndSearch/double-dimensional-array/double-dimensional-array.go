package main

import "fmt"

/**
题目描述
	二维数组排序，将二维数组 arr 按照 sort 中指定的列排序
*/
func main() {
	var arr [][]int
	var sort []int
	arr = [][]int{
		[]int{2, 4, 7, 10},
		[]int{7, 1, 4, 3},
		[]int{6, 3, 4, 5},
		[]int{1, 4, 7, 9},
	}
	sort = []int{
		0,
	}
	_ = sort

	quickSort(arr, 1, 0, len(arr)-1)
	fmt.Println(arr)
}

func doubleDimensionalArray(arr [][]int, sort []int) {
	
}

func quickSort(arr [][]int, sort, L, R int) {
	if L >= R {
		return
	}
	l, r, mid := L, R, arr[L]
	for l < r {
		for l < r && arr[r][sort] > mid[sort] {
			r--
		}
		if l < r {
			arr[l] = arr[r]
			l++
		}
		for l < r && arr[l][sort] < mid[sort] {
			l++
		}
		if l < r {
			arr[r] = arr[l]
			r--
		}
	}
	arr[l] = mid
	quickSort(arr, sort, L, r-1)
	quickSort(arr, sort, r+1, R)
}
