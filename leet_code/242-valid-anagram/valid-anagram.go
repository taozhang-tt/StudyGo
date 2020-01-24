package main

import "fmt"

//先对两个字符串排序，比较排序后对字符串是否相等
func isAnagram(s string, t string) bool {
	ss := []uint8(s)
	tt := []uint8(t)
	sortedString(ss, 0, len(ss)-1)
	sortedString(tt, 0, len(tt)-1)
	if string(ss) == string(tt) {
		return true
	}
	return false
}

//使用快排对字符串排序
func sortedString(s []uint8, left int, right int) {
	if left >= right {
		return
	}
	l := left
	r := right
	mid := s[l]
	for l < r {
		for l < r && s[r] > mid {
			r--
		}
		if l < r {
			s[l] = s[r]
			l++
		}
		for l < r && s[l] < mid {
			l++
		}
		if l < r {
			s[r] = s[l]
			r--
		}
	}
	s[l] = mid
	sortedString(s, left, l-1)
	sortedString(s, l+1, right)
}

//使用map集合统计的方式
func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[uint8]int, 26)
	for i := 0; i < len(s); i++ {
		m[s[i]] ++
		m[t[i]] --
	}
	for _, value := range m {
		if value != 0 {
			return false
		}
	}
	return true
}

func main() {
	s := "anagram"
	t := "nagaram"
	ret := isAnagram2(s, t)
	fmt.Println(ret)
}
