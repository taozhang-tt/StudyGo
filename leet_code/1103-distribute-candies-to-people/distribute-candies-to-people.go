package main

import (
	"fmt"
	"math"
)

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
		if i+1<candies {
			ret[i%num_people] += i+1
			candies -= i+1
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
	p := int(math.Floor(math.Sqrt(2.0 * float64(candies) + 0.25)-0.5))
	remaining := candies - (p + 1) * p /2
	//完整的分发回合数
	rows := p/num_people
	cols := p%num_people
	for i:=0; i<num_people; i++ {
		ret[i] = (i+1) * rows + rows*(rows-1)*num_people/2
		if i < cols {
			ret[i] += i+1 + rows*num_people
		}
	}
	ret[cols] += remaining
	return ret
}