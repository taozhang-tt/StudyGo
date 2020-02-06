package main

import "fmt"

func main() {
	prices := []int{1, 2, 3, 0, 2}
	fmt.Println(maxProfit(prices))
}

/**
动态规划解题思路：
	状态定义：mp[i][j][k]；i 表示天数，取值 0 到 n-1；j 表示是否持有股票，取值 0、1；k 表示有无 cooldown，取值 0、1；

	状态转换分析：
		持有股票：
			（1）无冷冻：前一天持有且无冷冻、前一天持有且有冷冻（不成立）、前一天未持有且无冷冻&今天买入、前一天未持有且有冷冻（刚卖掉，今天无法操作）
			（2）有冷冻：这种情况不成立
		未持有股票：
			（1）无冷冻：前一天持有且无冷冻（今天不能卖，卖了就有冷冻）、前一天持有且有冷冻（不成立）、前一天未持有且无冷冻、前一天未持有且有冷冻
			（2）有冷冻：前一天持有且无冷冻（今天卖掉）、前一天持有且有冷冻（不成立）、前一天未持有且无冷冻（今天不可能有冷冻）、前一天未持有且有冷冻（今天只能买入、不可能有冷冻）
	状态转换方程：
		持有且无冷冻：mp[i][1][0] = max(mp[i-1][1][0], mp[i-1][0][0] - a[i])
		未持有且无冷冻：mp[i][0][0] = max(mp[i-1][0][0], mp[i-1][0][1])
		未持有且有冷冻: mp[i][0][1] = mp[i-1][1][0] + a[i]
*/
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	ret := 0
	mp := make([][][]int, len(prices))

	for i := 0; i < len(prices); i++ {
		mp[i] = make([][]int, 2)
		mp[i][0] = make([]int, 2)
		mp[i][1] = make([]int, 2)
		if i == 0 {
			mp[i][1][0] = -prices[i]
			continue
		}

		//持有且无冷冻
		if mp[i-1][1][0] > mp[i-1][0][0]-prices[i] {
			mp[i][1][0] = mp[i-1][1][0]
		} else {
			mp[i][1][0] = mp[i-1][0][0] - prices[i]
		}
		if mp[i][1][0] > ret {
			ret = mp[i][1][0]
		}
		//未持有且无冷冻
		if mp[i-1][0][0] > mp[i-1][0][1] {
			mp[i][0][0] = mp[i-1][0][0]
		} else {
			mp[i][0][0] = mp[i-1][0][1]
		}
		if mp[i][0][0] > ret {
			ret = mp[i][0][0]
		}
		//未持有且有冷冻
		mp[i][0][1] = mp[i-1][1][0] + prices[i]
		if mp[i][0][1] > ret {
			ret = mp[i][0][1]
		}
	}
	return ret
}
