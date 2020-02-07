package main

import "fmt"

func main() {
	word1 := "a"
	word2 := "a"
	fmt.Println(minDistance(word1, word2))
}

/**
动态规划解题思路：
	状态定义：dp[i][j]，word1 的前 i 个字符与 word2 的前 j 个字符匹配所需要的最少步骤
	状态转换方程：
		（1）如果 word1[i] == word2[j]
			dp[i][j] = dp[i-1][j-1]
		（2）如果 word1[i] != word2[j]，那么对于有三种操作，分别是 插入（word1）、删除（word2）、替换
			dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
 */
func minDistance(word1 string, word2 string) int {
	len1 := len(word1)
	len2 := len(word2)
	dp := make([][]int, len1+1)
	for i:=0; i<=len1; i++ {
		dp[i] = make([]int, len2+1)
		dp[i][0] = i
	}
	for j:=0; j<=len2; j++ {
		dp[0][j] = j
	}
	for i:=1;i<=len1; i++ {
		for j:=1; j<=len2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				min := dp[i-1][j]
				if dp[i][j-1] < min {
					min = dp[i][j-1]
				}
				if dp[i-1][j-1] < min {
					min = dp[i-1][j-1]
				}
				dp[i][j] = min + 1
			}
		}
	}
	return dp[len1][len2]
}
