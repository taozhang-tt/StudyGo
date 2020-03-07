package main

import "fmt"

/**
212. 单词搜索 II
	https://leetcode-cn.com/problems/word-search-ii/
题目描述：
	给定一个二维网格 board 和一个字典中的单词列表 words，找出所有同时在二维网格和字典中出现的单词。
	单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母在一个单词中不允许被重复使用。
示例:

输入:
	words = ["oath","pea","eat","rain"] and board =
	[
	  ['o','a','a','n'],
	  ['e','t','a','e'],
	  ['i','h','k','r'],
	  ['i','f','l','v']
	]
输出: ["eat","oath"]
	说明:
	你可以假设所有输入都由小写字母 a-z 组成。
*/

func main() {
	//board := [][]byte{
	//	[]byte{'o','a','a','n'},
	//	[]byte{'e','t','a','e'},
	//	[]byte{'i','h','k','r'},
	//	[]byte{'i','f','l','v'},
	//}
	//words := []string{
	//	"oath",
	//	"pea",
	//	"eat",
	//	"rain",
	//}
	board := [][]byte{
		[]byte{'a', 'b'},
		[]byte{'a', 'a'},
	}
	words := []string{"aba", "baa", "bab", "aaab", "aaa", "aaaa", "aaba"}
	ret := findWords(board, words)
	fmt.Println(ret)
}

//先实现一个字典树

type TrieNod struct {
	Val      int32
	Children [26]*TrieNod
	IsEnd    bool
}

type Trie struct {
	Root *TrieNod
}

func (obj *Trie) Insert(word string) {
	root := obj.Root
	for _, item := range word {
		if root.Children[item-'a'] == nil {
			root.Children[item-'a'] = &TrieNod{
				Val:      item - 'a',
				Children: [26]*TrieNod{},
				IsEnd:    false,
			}
		}
		root = root.Children[item-'a']
	}
	root.IsEnd = true
}

func (obj *Trie) Search(word string) bool {
	root := obj.Root
	for _, item := range word {
		if root.Children[item-'a'] == nil {
			return false
		}
		root = root.Children[item-'a']
	}
	if !root.IsEnd {
		return false
	}
	return true
}

func (obj *Trie) StartWith(word string) bool {
	root := obj.Root
	for _, item := range word {
		if root.Children[item-'a'] == nil {
			return false
		}
		root = root.Children[item-'a']
	}
	return true
}

func findWords(board [][]byte, words []string) []string {
	//先把所有单词保存到字典树中
	trie := &Trie{
		Root: &TrieNod{
			Val:      0,
			Children: [26]*TrieNod{},
			IsEnd:    false,
		}}
	for _, item := range words {
		trie.Insert(item)
	}
	m := make(map[string]bool, 0)
	for x := 0; x < len(board); x++ {
		for y := 0; y < len(board[0]); y++ {
			dfs(board, trie, "", x, y, m)
		}
	}
	ret := make([]string, 0)
	for value, _ := range m {
		ret = append(ret, value)
	}
	return ret
}

func dfs(board [][]byte, trie *Trie, current string, x, y int, ret map[string]bool) ([][]byte, bool) {
	char := board[x][y]
	current += string(char)
	if trie.Search(current) {
		ret[current] = true
	}
	//剪枝
	if !trie.StartWith(current) {
		return board, false
	}
	//移动方向，上下左右
	board[x][y] = '@'
	dx := [4]int{-1, 1, 0, 0}
	dy := [4]int{0, 0, -1, 1}
	sign := false
	for i := 0; i < 4; i++ {
		xx := x + dx[i]
		yy := y + dy[i]
		if xx >= 0 && xx < len(board) && yy >= 0 && yy < len(board[0]) && board[xx][yy] != '@' {
			if _, ok := dfs(board, trie, current, xx, yy, ret); ok {
				//board[x][y] = char
				//return board, ok
				sign = true
			}
		}
	}
	board[x][y] = char
	return board, sign
}
