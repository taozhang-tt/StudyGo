package main

import "fmt"

/**
914. 卡牌分组
	https://leetcode-cn.com/problems/x-of-a-kind-in-a-deck-of-cards/
题目描述：
	给定一副牌，每张牌上都写着一个整数。
	此时，你需要选定一个数字 X，使我们可以将整副牌按下述规则分成 1 组或更多组：
	每组都有 X 张牌。
	组内所有的牌上都写着相同的整数。
	仅当你可选的 X >= 2 时返回 true。
示例 1：
	输入：[1,2,3,4,4,3,2,1]
	输出：true
	解释：可行的分组是 [1,1]，[2,2]，[3,3]，[4,4]
*/

func main() {
	var deck = []int{1, 2, 3, 4, 4, 3, 2, 1}
	fmt.Println(hasGroupsSizeX(deck))
	deck = []int{1, 1, 1, 2, 2, 2, 3, 3}
	fmt.Println(hasGroupsSizeX(deck))
	deck = []int{1, 1}
	fmt.Println(hasGroupsSizeX(deck))
	deck = []int{1, 1, 2, 2, 2, 2}
	fmt.Println(hasGroupsSizeX(deck))
}

func hasGroupsSizeX(deck []int) bool {
	if len(deck) < 2 {
		return false
	}
	deckMap := make(map[int]int, len(deck)) //统计每张卡牌的个数
	for _, item := range deck {
		deckMap[item] = deckMap[item] + 1
	}
	min := 0
	for _, item := range deckMap { //寻找卡牌个数的最小值
		if min == 0 {
			min = item
			continue
		}
		if item < min {
			min = item
		}
	}
	for i := 2; i <= min; i++ { //寻找一个所有卡牌数的最小公约数
		sign := true
		for _, item := range deckMap {
			if item%i != 0 {
				sign = false
			}
		}
		if sign {
			return true
		}
	}
	return false
}
