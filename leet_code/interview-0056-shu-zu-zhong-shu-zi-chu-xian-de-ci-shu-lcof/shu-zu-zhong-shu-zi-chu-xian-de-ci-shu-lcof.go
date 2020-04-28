package main

import "fmt"

/**
面试题56 - I. 数组中数字出现的次数
	https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof/
题目描述：
	一个整型数组 nums 里除两个数字之外，其他数字都出现了两次。请写程序找出这两个只出现一次的数字。要求时间复杂度是O(n)，空间复杂度是O(1)。
示例 1：
	输入：nums = [4,1,4,6]
	输出：[1,6] 或 [6,1]
示例 2：
	输入：nums = [1,2,10,4,1,4,3,3]
	输出：[2,10] 或 [10,2]
限制：
	2 <= nums <= 10000
*/

func main() {
	var nums []int
	var ret []int
	nums = []int{4, 1, 4, 6}
	ret = singleNumbers(nums)
	fmt.Println(ret)
}

/**
位运算
	预备知识：对于两个相等的数做异或运算，结果为 0，即 a ^ a = 0
	假设两个不相等的数分别为 a、b，那么 a ^ b != 0 必然成立；如果我们能把 nums 所有元素做异或操作，那么结果应该等于 a ^ b
	如果我们能将 nums 分为两组，第一组包含的元素除了 a 以外，其它数字都出现了两次；第二组除了 b 以外其它都出现了两次，那么第一组所有元素异或即为 a，相应的第二组为 b
	怎么分组？因为 a ^ b != 0，那么对于两者做异或的结果c的某一位肯定为1，我们可以取从右往左数的第一个为1的位置，然后遍历所有元素，如果该位为1的元素就分到第一组，否则分到第二组
*/
func singleNumbers(nums []int) []int {
	total := 0
	for _, item := range nums {
		total ^= item
	}
	sign := 1
	for sign&total == 0 {
		sign <<= 1
	}
	a, b := 0, 0
	for _, item := range nums {
		if item&sign == 0 {
			a ^= item
		} else {
			b ^= item
		}
	}
	return []int{a, b}
}
