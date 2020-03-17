package main

import "fmt"

/**
1013. 将数组分成和相等的三个部分
	https://leetcode-cn.com/problems/partition-array-into-three-parts-with-equal-sum/
题目描述：
	给你一个整数数组 A，只有可以将其划分为三个和相等的非空部分时才返回 true，否则返回 false。
	形式上，如果可以找出索引 i+1 < j 且满足 (A[0] + A[1] + ... + A[i] == A[i+1] + A[i+2] + ... + A[j-1] == A[j] + A[j-1] + ... + A[A.length - 1]) 就可以将数组三等分。
 */
func main() {
	A := []int{0,2,1,-6,6,-7,9,1,2,0,1}
	fmt.Println(canThreePartsEqualSum(A))
	A = []int{0,2,1,-6,6,7,9,-1,2,0,1}
	fmt.Println(canThreePartsEqualSum(A))
	A = []int{3,3,6,5,-2,2,5,1,-9,4}
	fmt.Println(canThreePartsEqualSum(A))
}

func canThreePartsEqualSum(A []int) bool {
	total := 0
	for _, item := range A {
		total += item
	}
	if total % 3 > 0 {
		return false
	}
	single := total / 3
	sum := 0
	count := 0
	for i:=0; i<len(A) - 1; i++ {
		sum += A[i]
		if sum == single {
			sum = 0
			count ++
		}
	}
	if count >= 2 {
		return true
	}
	return false
}
