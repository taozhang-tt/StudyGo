package main

import (
	"container/heap"
	"fmt"
)

//先实现一个最小堆
type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	var k = 3
	h := new(IntHeap)
	param := []int{4, 5, 8, 2, 3, 5, 10, 9, 4}
	for _, item := range param {
		if h.Len() < k {
			heap.Push(h, item)
		} else {
			top := heap.Pop(h).(int)
			if top < item {
				top = item
			}
			heap.Push(h, top)
		}

	}
	fmt.Println(heap.Pop(h))
}
