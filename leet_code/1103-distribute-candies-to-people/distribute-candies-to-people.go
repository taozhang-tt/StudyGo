package main

import (
	"fmt"
	"math"
)

/**
1103. 分糖果 II
	https://leetcode-cn.com/problems/distribute-candies-to-people/
题目描述：

	排排坐，分糖果。
	我们买了一些糖果 candies，打算把它们分给排好队的 n = num_people 个小朋友。
	给第一个小朋友 1 颗糖果，第二个小朋友 2 颗，依此类推，直到给最后一个小朋友 n 颗糖果。
	然后，我们再回到队伍的起点，给第一个小朋友 n + 1 颗糖果，第二个小朋友 n + 2 颗，依此类推，直到给最后一个小朋友 2 * n 颗糖果。
	重复上述过程（每次都比上一次多给出一颗糖果，当到达队伍终点后再次从队伍起点开始），直到我们分完所有的糖果。注意，就算我们手中的剩下糖果数不够（不比前一次发出的糖果多），这些糖果也会全部发给当前的小朋友。
	返回一个长度为 num_people、元素之和为 candies 的数组，以表示糖果的最终分发情况（即 ans[i] 表示第 i 个小朋友分到的糖果数）。
示例 1：
	输入：candies = 7, num_people = 4
	输出：[1,2,3,1]
	解释：
	第一次，ans[0] += 1，数组变为 [1,0,0,0]。
	第二次，ans[1] += 2，数组变为 [1,2,0,0]。
	第三次，ans[2] += 3，数组变为 [1,2,3,0]。
	第四次，ans[3] += 1（因为此时只剩下 1 颗糖果），最终数组变为 [1,2,3,1]。
*/

func main() {
	candies := 90
	num_people := 4
	fmt.Println(distributeCandies3(candies, num_people))
}

/**
自己写的暴力解法
*/
func distributeCandies(candies int, num_people int) []int {
	ret := make([]int, num_people)
	circle := 0
	for candies > 0 {
		for i := 1; i <= num_people; i++ {
			nums := circle*num_people + i
			if candies < nums {
				ret[i-1] += candies
				candies = 0
				break
			}
			ret[i-1] += nums
			candies -= nums
		}
		circle++
	}
	return ret
}

/**
力扣官方的暴力解法
	tips：利用取摩运算来循环数组的元素
*/
func distributeCandies2(candies int, num_people int) []int {
	ret := make([]int, num_people)
	i := 0
	for candies > 0 {
		if i+1 < candies {
			ret[i%num_people] += i + 1
			candies -= i + 1
		} else {
			ret[i%num_people] += candies
			candies = 0
		}
		i++
	}
	return ret
}

/**
等差数列的性质
	公差：d
	首项: a1
	第n项：an = a1 + d*(n-1)
	前n项和：Sn = (a1+an)*n/2
	前n项和：Sn = a1*n + n*(n-1)*d/2

题解：
	1. 假设获得完整礼物的人数为p，这是一个公差为1的等差数列，前p项和为 Sp = p*(p+1)/2
	2. 假设所有糖果数量为 C，最后剩余的不够一份礼物的糖果数量为 remaining = C - Sp，则有 0 <= remaining < p + 1
		推导出 sqrt(2C + 0.25)-1.5 < p < sqrt(2C + 0.25) - 0.5
	3. 假设人数为 N，完整地分发礼物的回合数 rows = p / N
	4. 在 rows 个回合中，第 i 个人获得的糖果数为 d[i] = i + (i+N) + (i+2N) + ... + (i+(rows-1)*N)
		即是：i * rows + N*rows*(rows-1)/2
	5. 不完整的回合还有部分人领取到了完整的礼物 cols = p % N，这部分人里第 i 个人获取的糖果数为 d[i] = i + N * rows
	6. 最后以为拥有礼物的人获取的糖果数为 d[cols+1] = remaining
*/

func distributeCandies3(candies int, num_people int) []int {
	ret := make([]int, num_people)
	//收到完整礼物的人数
	p := int(math.Floor(math.Sqrt(2.0*float64(candies)+0.25) - 0.5))
	remaining := candies - (p+1)*p/2
	//完整的分发回合数
	rows := p / num_people
	cols := p % num_people
	for i := 0; i < num_people; i++ {
		ret[i] = (i+1)*rows + rows*(rows-1)*num_people/2
		if i < cols {
			ret[i] += i + 1 + rows*num_people
		}
	}
	ret[cols] += remaining
	return ret
}
