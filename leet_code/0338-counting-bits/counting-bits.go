package main

import "fmt"

func main() {
	fmt.Println(countBits2(5))
}

//思路: 主要考察二进制数性质的, 即如果A/2 = B, 那么A比B多了一位, 并且A和B出了A的二进制的右边的一位以外其他都一样,
//举个栗子A = 11, 二进制就是1011, B = 5, 二进制是101, 所以我们可以看出其最左边是相等的, 只有A的最后一位不等.
func countBits(num int) []int {
	count := make([]int, num+1, num+1)
	for i:=1; i<=num; i++ {
		count[i] = count[i/2] + i%2
	}
	return count
}

/**
通过位运算
	对于正整数 A，B = A & (A-1)，相当于将 A 的2进制的最后一个1清0，这时候 B 肯定小于 A，且 A 的二进制1的位数 等于 B 的二进制1的位数再加1
 */
func countBits2(num int) []int {
	count := make([]int, num+1, num+1)
	for i:=1; i<=num; i++ {
		temp := uint32(i)
		count[i] = count[temp & (temp-1)] + 1
	}
	return count
}
