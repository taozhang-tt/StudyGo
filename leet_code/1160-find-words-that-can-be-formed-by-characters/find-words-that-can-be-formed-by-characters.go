package main

import "fmt"

/**
1160. 拼写单词
	https://leetcode-cn.com/problems/find-words-that-can-be-formed-by-characters/
题目描述：
	给你一份『词汇表』（字符串数组） words 和一张『字母表』（字符串） chars。
	假如你可以用 chars 中的『字母』（字符）拼写出 words 中的某个『单词』（字符串），那么我们就认为你掌握了这个单词。
	注意：每次拼写时，chars 中的每个字母都只能用一次。
	返回词汇表 words 中你掌握的所有单词的 长度之和。
示例 1：
	输入：words = ["cat","bt","hat","tree"], chars = "atach"
	输出：6
	解释：
	可以形成字符串 "cat" 和 "hat"，所以答案是 3 + 3 = 6。
 */
func main() {
	var words = []string{"cat","bt","hat","tree"}
	var chars = "atach"
	fmt.Println(countCharacters2(words, chars))
	words = []string{"hello","world","leetcode"}
	chars = "welldonehoneyr"
	fmt.Println(countCharacters2(words, chars))

}

/**
第一版
 */
func countCharacters(words []string, chars string) int {
	var charsMap = make(map[int32]int, 0)
	var ret int
	//统计所有的字符及其个数
	for _, char := range chars {
		charsMap[char] += 1
	}

	for _, word := range words {
		wordMap := make(map[int32]int, 0)
		wordCount := 0
		for _, char := range word {
			wordMap[char] += 1
			if wordMap[char] > charsMap[char] {
				continue
			}
		}
		for char, cnt := range wordMap {
			if charsMap[char] >= cnt {
				wordCount += cnt
			} else {
				wordCount = 0
				break
			}
		}
		ret += wordCount
	}
	return ret
}

/**
优化版
 */
func countCharacters2(words []string, chars string) int {
	var charsMap = make(map[int32]int, 0)
	var ret int
	//统计所有的字符及其个数
	for _, char := range chars {
		charsMap[char] += 1
	}
	for _, word := range words {
		ret += len(word)
		wordMap := make(map[int32]int, 0)	//用于记录单词中每个字符出现的次数
		for _, char := range word {
			wordMap[char] += 1
			if wordMap[char] > charsMap[char] {
				ret -= len(word)
				break
			}
		}
	}
	return ret
}

