/**
763. 划分字母区间
	https://leetcode-cn.com/problems/partition-labels/
题目描述：
	字符串 S 由小写字母组成。我们要把这个字符串划分为尽可能多的片段，同一个字母只会出现在其中的一个片段。返回一个表示每个字符串片段的长度的列表。
示例 1：
	输入：S = "ababcbacadefegdehijhklij"
	输出：[9,7,8]
	解释：
	划分结果为 "ababcbaca", "defegde", "hijhklij"。
	每个字母最多出现在一个片段中。
	像 "ababcbacadefegde", "hijhklij" 的划分是错误的，因为划分的片段数较少。
提示：
	S的长度在[1, 500]之间。
	S只包含小写字母 'a' 到 'z' 。
*/

package main

import "fmt"

func main() {
	var S string
	S = "ababcbacadefegdehijhklijss"
	fmt.Println(partitionLabels(S))
}

/**
分析可知：一个字符的所有下标必存在于一个分组内，即是说字符 a 若在分组一内，那么 a 最后出现的位置必存在于分组一
解题思路：
	先遍历一遍记录每个字符最后一次出现的位置 last = map[byte]int
	初始化两个指针 start, end := 0, 0
	从左到右遍历每个字符 ch，该字符当前的位置记录为 i, 该字符最后出现的位置为 last[ch]
		此时 end 记录的是字符 ch 之前任一字符在 S 中最后出现的最大值(max(last[c]), c是ch之前的任一字符)，当遍历到 ch 时，应该拿 end 和 last[ch] 做比较，并把较大值赋值给end
		判断 if (i == end)，如果两者相等，说明 start 到 end 位置所有的字符都只出现在 [start, end]这个区间内，他们可以被分割为一组
			把这个分组的大小（end-start+1）放入结果集
			start 挪动到 end 的下一个位置
*/
func partitionLabels(S string) []int {
	ret := make([]int, 0)
	last := make(map[byte]int)
	for i := 0; i < len(S); i++ {
		last[S[i]] = i
	}
	start, end := 0, 0
	for i, ch := range S {
		end = max(end, last[byte(ch)])
		if end == i {
			ret = append(ret, end-start+1)
			start = end+1
		}
	}
	return ret
}

/**
思路同上，都是官方题解，这是更优秀的写法
*/
func partitionLabels2(S string) []int {
	ret := make([]int, 0)
	last := make([]int, 26)
	for i := 0; i < len(S); i++ {
		last[S[i]-'a'] = i
	}
	start, end := 0, 0
	for i, ch := range S {
		end = max(end, last[ch-'a'])
		if end == i {
			ret = append(ret, end-start+1)
			start = end + 1
		}
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
