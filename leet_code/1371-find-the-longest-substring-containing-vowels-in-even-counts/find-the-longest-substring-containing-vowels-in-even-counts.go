package main

import "fmt"

/**
1371. 每个元音包含偶数次的最长子字符串
	https://leetcode-cn.com/problems/find-the-longest-substring-containing-vowels-in-even-counts/
题目描述
	给你一个字符串 s ，请你返回满足以下条件的最长子字符串的长度：每个元音字母，即 'a'，'e'，'i'，'o'，'u' ，在子字符串中都恰好出现了偶数次。
示例 1：
	输入：s = "eleetminicoworoep"
	输出：13
	解释：最长子字符串是 "leetminicowor" ，它包含 e，i，o 各 2 个，以及 0 个 a，u 。
示例 2：
	输入：s = "leetcodeisgreat"
	输出：5
	解释：最长子字符串是 "leetc" ，其中包含 2 个 e 。
示例 3：
	输入：s = "bcbcbc"
	输出：6
	解释：这个示例中，字符串 "bcbcbc" 本身就是最长的，因为所有的元音 a，e，i，o，u 都出现了 0 次。
提示：
	1 <= s.length <= 5 x 10^5
	s 只包含小写英文字母。
*/
func main() {
	var s string
	s = "eleetminicoworoep"
	fmt.Println(findTheLongestSubstring3(s))
	s = "leetcodeisgreat"
	fmt.Println(findTheLongestSubstring3(s))
	s = "bcbcbc"
	fmt.Println(findTheLongestSubstring3(s))
}

/**
暴力法、枚举法(超时)
	枚举所有子串，统计每个子串中出现每个元音字母的次数
*/
func findTheLongestSubstring(s string) int {
	max := 0
	for i := 0; i < len(s); i++ {
		cnt := map[byte]int{
			'a': 0,
			'e': 0,
			'i': 0,
			'o': 0,
			'u': 0,
		}
		sum := 0
		for j := i; j < len(s); j++ {
			sum++
			if _, exist := cnt[s[j]]; exist {
				cnt[s[j]]++
			}
			if sum > max {
				sign := true
				for _, num := range cnt {
					if num%2 != 0 {
						sign = false
						break
					}
				}
				if sign {
					max = sum
				}
			}
		}
	}
	return max
}

/**
前缀和
	利用前缀和可以优化一些时间
	我们用 prefix[i][k] 记录在前 i 个字符中，第 k 个元音字母出现的次数
	那么对于子串 [i, j]，我们想判断它是否满足题设条件，即是判断 prefix[j][k] - prefix[i-1][k] 是否满足题设，这样就把时间复杂度降低到 O(1)，不过总体的时间复杂度还是 O(n方)
*/
func findTheLongestSubstring2(s string) int {
	max := 0
	//统计前缀和
	prefix := make(map[int]map[byte]int, len(s))
	prefix[-1] = map[byte]int{
		'a': 0,
		'e': 0,
		'i': 0,
		'o': 0,
		'u': 0,
	}
	for i := 0; i < len(s); i++ {
		prefix[i] = make(map[byte]int)
		for ch, num := range prefix[i-1] {
			prefix[i][ch] = num
		}
		if _, exist := prefix[i][s[i]]; exist {
			prefix[i][s[i]]++
		}
	}
	//
	for i := 0; i < len(s); i++ {
		for j := len(s) - 1; j > i; j-- {
			if j-i+1 > max {
				sign := true
				for ch, num := range prefix[j] {
					if (num-prefix[i-1][ch])%2 != 0 {
						sign = false
						break
					}
				}
				if sign {
					max = j-i+1
				}
			}
		}
	}
	return max
}

/**
前缀和 & 状态压缩(官方题解)
	我们用数字 0 和 1 标志元音字母的状态，用一个五位的二进制数表示每个元音字母的状态；对于五个元音字母，a 用最后一位表示，e 用右数第二位，依次类推；
	我们不需要完全记录每个元音字母出现了多少次，只需要记录每个元音字母出现了偶数次还是奇数次，由 异或 运算的性质可知，00001 表示字母 a 出现了 奇数次，其它字母出现了偶数次
	假设对于 [0, i] 有 00011 且对于 [0, j] 也刚好有 00011，那我们由 奇数 - 奇数 = 偶数、偶数 - 偶数 = 偶数 的性质可知: [i+1, j] 之间每个元音字母必出现了偶数次
	我们用一个hash表保存每个二进制状态对应的 i
*/
func findTheLongestSubstring3(s string) int {
	pos, ret, status := make(map[int]int, 32), 0, 0	//pos 保存子区间 [0, i] 各个元音字母出现的奇偶状态
	pos[0] = -1	//假想在字符串开头插入了一个不存在的字符，索引为 -1，它的奇偶状态为 00000，什么都没出现过，是 00000 不过分
	for i :=0; i<len(s); i++ {
		switch s[i] {
		case 'a':
			status ^= 1 << 0
		case 'e':
			status ^= 1 << 1
		case 'i':
			status ^= 1 << 2
		case 'o':
			status ^= 1 << 3
		case 'u':
			status ^= 1 << 4
		}
		if num, exist := pos[status]; exist {// [0, i] 对应的状态 status，在 [0, j] 且 0 < j < i 这个区间时已经出现过，那么 [j+1, i] 这个区间内各个元音字母都出现了偶数次，i 与 j+1 之间有 i - j 个字符
			ret = max(ret, i-num)
		} else {
			pos[status] = i
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
