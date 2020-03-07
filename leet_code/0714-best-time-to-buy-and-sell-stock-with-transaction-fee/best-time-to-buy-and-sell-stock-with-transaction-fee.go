package main

import "fmt"

/**
714. 买卖股票的最佳时机含手续费
	https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/
题目描述：
	给定一个整数数组 prices，其中第 i 个元素代表了第 i 天的股票价格 ；非负整数 fee 代表了交易股票的手续费用。
	你可以无限次地完成交易，但是你每次交易都需要付手续费。如果你已经购买了一个股票，在卖出它之前你就不能再继续购买股票了。
	返回获得利润的最大值。
示例 1:
	输入: prices = [1, 3, 2, 8, 4, 9], fee = 2
	输出: 8
	解释: 能够达到的最大利润:
	在此处买入 prices[0] = 1
	在此处卖出 prices[3] = 8
	在此处买入 prices[4] = 4
	在此处卖出 prices[5] = 9
	总利润: ((8 - 1) - 2) + ((9 - 4) - 2) = 8.
*/

func main() {
	prices := []int{1, 3, 2, 8, 4, 9}
	fmt.Println(maxProfit(prices, 2))
}

/**
动态规划：
	思考过程：
		1. 定义状态：dp[i] 表示第 i 天的最大利润 max profit; 将状态改成以 mp[i] 记录
		2. 递推方程：第 i 天可以买入股票 a[i]，也可以卖出股票 a[i]，则有 mp[i] = max(mp[i-1] + a[i], mp[i-1] - a[i])
		   上述的递推方程是有误的，因为我们要想在第 i 天卖出股票，那要求我们必须在第 i-1 天买入了股票，可以发现我们丢失了是否持有股票的信息；总结就是：当我们发现无法递推的时候，一般都是我们定义的状态不够，需要增加维度

		3. 定义状态：mp[i][j]，i 表示天数取值 0 到 n-1，j 表示是否拥有股票取值 0 或 1，第 i 天的最大利润为 max(mp[i][0], mp[i][1])
		4. 递推方程：mp[i][0] = max(mp[i-1][0], mp[i-1][1]+a[i])，mp[i][1] = max(mp[i-1][1], mp[i-1][0]-a[i])
*/
func maxProfit(prices []int, fee int) int {
	if len(prices) < 2 {
		return 0
	}
	ret := 0
	mp := make([][2]int, len(prices))
	mp[0][0], mp[0][1] = 0, -prices[0]
	for i := 1; i < len(prices); i++ {
		//不持有股份
		if mp[i-1][0] > mp[i-1][1]+prices[i]-fee {
			mp[i][0] = mp[i-1][0]
		} else {
			mp[i][0] = mp[i-1][1] + prices[i] - fee
		}
		//持有一股
		if mp[i-1][1] > mp[i-1][0]-prices[i] {
			mp[i][1] = mp[i-1][1]
		} else {
			mp[i][1] = mp[i-1][0] - prices[i]
		}
		if mp[i][0] > ret {
			ret = mp[i][0]
		}
		if mp[i][1] > ret {
			ret = mp[i][1]
		}
	}
	return ret
}
