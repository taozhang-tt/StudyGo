package main

import "fmt"

func main() {
	var K, N, ret int
	K, N = 1, 2
	ret = superEggDrop(K, N)
	fmt.Println(ret)
	K, N = 2, 6
	ret = superEggDrop(K, N)
	fmt.Println(ret)
	K, N = 3, 14
	ret = superEggDrop(K, N)
	fmt.Println(ret)
}

func superEggDrop(K int, N int) int {
	return dp(K, N)
}

var m = map[int]int{}

func dp(K int, N int) int {
	if _, exist := m[N*100+K]; !exist {
		ans := 0
		if N == 0 {
			ans = 0
		} else if K == 1 {
			ans = N
		} else {
			lo, hi := 1, N
			for (lo + 1) < hi {
				x := (lo + hi) / 2
				t1, t2 := dp(K-1, x-1), dp(K, N-x)
				if t1 < t2 {
					lo = x
				} else if t1 > t2 {
					hi = x
				} else {
					lo, hi = x, x
				}
			}
			t1 := dp(K-1, lo-1)
			if dp(K, N-lo) > t1 {
				t1 = dp(K, N-lo)
			}
			t2 := dp(K-1, hi-1)
			if dp(K, N-hi) > t2 {
				t2 = dp(K, N-hi)
			}
			if t1 < t2 {
				ans = 1 + t1
			} else {
				ans = 1 + t2
			}
		}
		m[N*100+K] = ans
	}
	return m[N*100+K]
}

func dp2(k, n, x int) int {
	if _, exist := m[n*100+k]; !exist {
		if n == 0 {
			return 0
		}
		if k == 1 {
			return n
		}
		ans := dp(k, n)
	} else {
		return m[n*100+k]
	}
}
