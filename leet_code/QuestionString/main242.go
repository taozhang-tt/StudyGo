package main

import (
	"fmt"
)

func main() {

	var (
		s = "anagram"
		t = "nagaram"
		//s = "rat"
		//t = "car"
	)
	fmt.Println(isAnagram(s, t))
}

func isAnagram(s string, t string) bool {
	var (
		tempSli = make(map[uint8]int, 26)
	)
	if len(s) != len(t) {
		return false
	}
	for i := 0; i < len(s); i++ {
		tempSli[s[i]] ++
		tempSli[t[i]] --
	}
	for _, value := range tempSli {
		if value != 0 {
			return false
		}
	}
	return true
}
