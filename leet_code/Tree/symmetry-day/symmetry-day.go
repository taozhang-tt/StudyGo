package main

import (
	"fmt"
)

func main() {
	ret := make([][8]int, 0)
	curr := [8]int{}
	for i:=0; i<10; i++ {
		ret = dfs(curr, 0, ret)
	}
	fmt.Println(ret)
	fmt.Println(len(ret))
}

func dfs(curr [8]int, level int, ret [][8]int) [][8]int {
	if level > 3 {
		if isValid(curr) {
			ret = append(ret, curr)
		}
		return ret
	}
	//剪枝
	for i:=0; i<10; i++ {
		if level == 3 && i > 1 {
			continue
		}
		if level == 1 && i > 3 {
			continue
		}
		curr[level] = i
		curr[7-level] = i
		ret = dfs(curr, level+1, ret)
	}
	return ret
}

func isValid(date [8]int) (ret bool) {
	var (
		year         int
		month        int
		day          int
		mdMap        = map[int]int{
			1:  31,
			2:  28,
			3:  31,
			4:  30,
			5:  31,
			6:  30,
			7:  31,
			8:  31,
			9:  30,
			10: 31,
			11: 30,
			12: 31,
		}
	)
	day = date[6] * 10 + date[7]
	month = date[4] * 10 + date[5]
	if day == 0 || month == 0 || month > 12 {
		return false
	}
	year = date[0] * 1000 + date[1]*100 + date[2]*10 + date[3]
	//判断是否为闰年
	if (year%100 == 0 && year%400 == 0) || (year%100 != 0 && year%4 == 0) {
		mdMap[2] = 29
	}
	if day > mdMap[month] {
		return false
	}
	return true
}
