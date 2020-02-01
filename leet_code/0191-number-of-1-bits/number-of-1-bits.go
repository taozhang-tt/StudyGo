package main

import "fmt"

func main() {
	var num uint32
	num = 00000000000000000000000000001011
	fmt.Println(hammingWeight2(num))
	num = 00000000000000000000000010000000
	fmt.Println(hammingWeight2(num))
	//num = 11111111111111111111111111111101
	//fmt.Println(hammingWeight(num))
}

//依次枚举每一位
func hammingWeight(num uint32) int {
	count := 0
	for num > 0 {
		if num % 2 > 0 {
			count++
		}
		num /= 2
	}
	return count
}

//通过位运算加速； 通过 X & (X-1) 操作可以清零最低位的 1
func hammingWeight2(num uint32) int {
	count := 0
	for num > 0 {
		num = num & (num-1)
		count++
	}
	return count
}