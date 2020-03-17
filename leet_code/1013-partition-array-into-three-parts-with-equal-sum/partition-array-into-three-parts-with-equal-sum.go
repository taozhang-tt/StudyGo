package main

import "fmt"

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
