package main

import "fmt"

func main() {
	var nums []int
	nums = []int{2, 2, 2, 1}
	fmt.Println(singleNumber1(nums))
}

/**
Map
	统计每个数字出现的次数，然后返回出现了一次的数字
*/
func singleNumber(nums []int) int {
	m := make(map[int]int, 0)
	for _, num := range nums {
		m[num] += 1
	}
	for num, cnt := range m {
		if cnt == 1 {
			return num
		}
	}
	return 0
}

/**
Map
	统计所有的数字（去重）
	统计所有数字（不去重）
	(a+b+c)*3 - (a+a+a+b+b+b+c) = 2c
*/
func singleNumber1(nums []int) int {
	t1, t2, m := 0, 0, make(map[int]bool)
	for _, num := range nums {
		t2 += num
		if _, exist := m[num]; !exist {
			t1 += num
			m[num] = true
		}
	}
	return (t1*3 - t2) / 2
}
