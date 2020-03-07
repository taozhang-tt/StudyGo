package main

import "fmt"

/**
208. 实现 Trie (前缀树)
	https://leetcode-cn.com/problems/implement-trie-prefix-tree/
题目描述：
	实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。
示例:
	Trie trie = new Trie();
	trie.insert("apple");
	trie.search("apple");   // 返回 true
	trie.search("app");     // 返回 false
	trie.startsWith("app"); // 返回 true
	trie.insert("app");
	trie.search("app");     // 返回 true
说明:
	你可以假设所有的输入都是由小写字母 a-z 构成的。
	保证所有输入均为非空字符串。
*/

func main() {
	trie := Constructor()
	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))
	fmt.Println(trie.Search("app"))
	fmt.Println(trie.StartsWith("app"))
	trie.Insert("app")
	fmt.Println(trie.Search("app"))
}

/**
字典树的节点
  Val 保存父亲节点到该节点经过到那条路径上的字符，为了方便，这里保存的是该字符的 ASCII 码 减去 字符 'a' 的 ASCII 码
  IsEnd 标志字典树中保存的有无以该字符为结尾的word
  Children 该字符后面还可能出现的字符，题设是只会出现26个小写字母，所以使用了长度为26的数组，如果包好其它字符，则需要扩充
*/
type TrieNode struct {
	Val      int32
	IsEnd    bool
	Children [26]*TrieNode
}

type Trie struct {
	Root *TrieNode
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		Root: &TrieNode{
			Val:      0,
			IsEnd:    false,
			Children: [26]*TrieNode{},
		},
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	root := this.Root
	for _, item := range word { //遍历 word 的每一个字符
		if root.Children[item-'a'] == nil { //判断从根父亲结点出去的路径中有无该字符
			root.Children[item-'a'] = &TrieNode{ //如果没用，则父亲节点增加这一路径
				Val:      item - 'a',
				IsEnd:    false,
				Children: [26]*TrieNode{},
			}
		}
		root = root.Children[item-'a'] //父亲节点移动到前一个被遍历的字符对应的节点上，继续判断下一个字符
	}
	root.IsEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	root := this.Root
	for _, item := range word {
		if root.Children[item-'a'] == nil {
			return false
		} else {
			root = root.Children[item-'a']
		}
	}
	if root.IsEnd { //这里要判断是不是到了单词的结尾，因为也可能是某个单词的前缀
		return true
	}
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	root := this.Root
	for _, item := range prefix {
		if root.Children[item-'a'] == nil {
			return false
		} else {
			root = root.Children[item-'a']
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
