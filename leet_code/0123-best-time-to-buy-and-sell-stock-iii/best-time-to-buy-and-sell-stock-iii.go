package main

import "fmt"

func main() {
	prices := []int{3,3,5,0,0,3,1,4}
	fmt.Println(maxProfit(prices))
}

/**
动态规划：
	思考过程：
		1. 定义状态：dp[i] 表示第 i 天的最大利润 max profit; 将状态改成以 mp[i] 记录
		2. 递推方程：第 i 天可以买入股票 a[i]，也可以卖出股票 a[i]，则有 mp[i] = max(mp[i-1] + a[i], mp[i-1] - a[i])
		   上述的递推方程是有误的，因为我们要想在第 i 天卖出股票，那要求我们必须在第 i-1 天买入了股票，可以发现我们丢失了是否持有股票的信息；总结就是：当我们发现无法递推的时候，一般都是我们定义的状态不够，需要增加维度

		3. 定义状态：mp[i][j]，i 表示天数取值 0 到 n-1，j 表示是否拥有股票取值 0、1、2 分别表示未持有股票、持有股票、已出售股票；第 i 天的最大利润为 max(mp[i][0], mp[i][1]，mp[i][2])
		4. 递推方程：mp[i][0] = max(mp[i-1][0], mp[i-1][1]+a[i])，mp[i][1] = max(mp[i-1][1], mp[i-1][0]-a[i])

		5. 如果限制交易的次数，那定义的状态还需要增加一维用于记录交易信息
		6. 定义状态：mp[i][k][j]，i、j如上所述，k表示交易的次数
		7. 递推方程：每次递推要对 K 遍历，k 的取值为 0 到 K-1
			mp[i][k][0] = max(mp[i-1][k][0], mp[i-1][k-1][1]+a[i])
			mp[i][k][1] = max(mp[i-1][k][1], mp[i-1][k-1][0]-a[i])
*/

func maxProfit(prices []int) int {
	N := len(prices)
	K := 2
	if N == 0 {
		return 0
	}
	//初始化
	mp := make([][][]int, len(prices))
	ret := 0
	//递推
	for i := 0; i < len(prices); i++ {
		mp[i] = make([][]int, K+1)
		for k := 0; k <= K; k++ {
			mp[i][k] = make([]int, 2)
			mp[i][k][0] = 0
			mp[i][k][1] = -prices[i]
			if i == 0 {
				continue
			}
			//未持有股票
			if k == 0 || (mp[i-1][k][0] > mp[i-1][k][1]+prices[i]) {
				mp[i][k][0] = mp[i-1][k][0]
			} else {
				mp[i][k][0] = mp[i-1][k][1] + prices[i]
			}
			if mp[i][k][0] > ret {
				ret = mp[i][k][0]
			}
			//持有股票，在买进时记录交易次数
			if k == 0 || (mp[i-1][k][1] > mp[i-1][k-1][0]-prices[i]) {
				mp[i][k][1] = mp[i-1][k][1]
			} else {
				mp[i][k][1] = mp[i-1][k-1][0] - prices[i]
			}
			if mp[i][k][1] > ret {
				ret = mp[i][k][1]
			}
		}
	}
	return ret
}

