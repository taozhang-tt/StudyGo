package main

import "fmt"

/**
820. 单词的压缩编码
	https://leetcode-cn.com/problems/short-encoding-of-words/
题目描述：
	给定一个单词列表，我们将这个列表编码成一个索引字符串 S 与一个索引列表 A。
	例如，如果这个列表是 ["time", "me", "bell"]，我们就可以将其表示为 S = "time#bell#" 和 indexes = [0, 2, 5]。
	对于每一个索引，我们可以通过从字符串 S 中索引的位置开始读取字符串，直到 "#" 结束，来恢复我们之前的单词列表。
	那么成功对给定单词列表进行编码的最小字符串长度是多少呢？
示例：
	输入: words = ["time", "me", "bell"]
	输出: 10
	说明: S = "time#bell#" ， indexes = [0, 2, 5] 。
*/

func main() {
	var words = []string{"me", "time"}
	ret := minimumLengthEncoding2(words)
	fmt.Println(ret)

	words = []string{"time", "time", "time"}
	ret = minimumLengthEncoding2(words)
	fmt.Println(ret)

	words = []string{"time", "me", "bell"}
	ret = minimumLengthEncoding2(words)
	fmt.Println(ret)

	words = []string{"t"}
	ret = minimumLengthEncoding2(words)
	fmt.Println(ret)
}

/**
字典树
*/
func minimumLengthEncoding(words []string) int {
	root := &TrieNode{ //初始化字典树的根节点
		IsEnd:       false,
		Children:    [26]*TrieNode{},
		HasChildren: false,
	}
	ret := 0
	for _, item := range words { //遍历每一个单词
		trie := root
		level := 0
		for i := len(item) - 1; i >= 0; i-- { //遍历每一个字符
			level++ //记录当前是字典树的第几层（根节点是第0层)，也即是当前遍历的字符的长度
			char := int32(item[i] - 'a')
			if trie.Children[char] != nil {	//字典树中已存在该分支
				trie = trie.Children[char]
				if trie.IsEnd {
					ret -= (level + 1) //从 S 中删除该单词 + #
					trie.IsEnd = false
				}
			} else {
				trie.Children[char] = &TrieNode{
					IsEnd:    false,
					Children: [26]*TrieNode{},
				}
				trie = trie.Children[char]
			}
		}
		//如果该节点没有其它孩子，则标记为叶子节点
		hasChildren := false
		for _, child := range trie.Children {
			if child != nil {
				hasChildren = true
				break
			}
		}
		if !hasChildren {
			trie.IsEnd = true
			ret += (level + 1)
		}
	}
	return ret
}

type TrieNode struct {
	Val         int32
	IsEnd       bool
	Children    [26]*TrieNode
	HasChildren bool
	Count       int
}

/**
字典树 2：采用了额外的空间记录所有单词在字典树中的节点
*/
func minimumLengthEncoding2(words []string) int {
	root := &TrieNode{ //初始化字典树的根节点
		Children: [26]*TrieNode{},
		Count:    0,
	}
	nodes := make(map[*TrieNode]int, 0)	//用于记录每个单词结尾对应的节点与单词的对应关系
	ret := 0

	for index, word := range words {
		trie := root
		for i := len(word) - 1; i >= 0; i-- {
			char := int32(word[i]-'a')
			if trie.Children[char] != nil {
				trie = trie.Children[char]
			} else {
				trie.Count++
				trie.Children[char] = &TrieNode{
					Children: [26]*TrieNode{},
					Count: 0,
				}
				trie = trie.Children[char]
			}
		}
		nodes[trie] = index
	}
	for node, index := range nodes {
		if node.Count == 0 {
			ret += (len(words[index]) + 1)
		}
	}
	return ret
}
