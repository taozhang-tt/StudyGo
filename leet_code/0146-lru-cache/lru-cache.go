package main

import (
	"container/list"
	"fmt"
)

func main() {
	//cache := Constructor(2)
	//cache.Put(1, 1)
	//cache.Put(2, 2)
	//fmt.Println(cache.Get(1))       // 返回  1
	//cache.Put(3, 3)
	//fmt.Println(cache.Get(2))       // 返回 -1 (未找到)
	//cache.Put(4, 4)
	//fmt.Println(cache.Get(1))       // 返回 -1 (未找到)
	//fmt.Println(cache.Get(3))       // 返回  3
	//fmt.Println(cache.Get(4))       // 返回  4

	cache := Constructor(2)
	cache.Put(2, 1)
	cache.Put(1, 1)
	cache.Put(2, 3)
	cache.Put(4, 1)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(2))
}

type LRUCache struct {
	Len int
	L   *list.List
	M   map[int]int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Len: capacity,
		L:   list.New(),
		M:   make(map[int]int, 0),
	}
}

func (this *LRUCache) Get(key int) int {
	if value, ok := this.M[key]; ok {
		this.L.MoveToFront(findElement(this.L.Front(), key))
		return value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	//先查询是否存在该 key，如果存在直接移动到头部
	if _, ok := this.M[key]; ok {
		front := findElement(this.L.Front(), key)
		front.Value = key
		this.M[key] = value
		this.L.MoveToFront(front)
		return
	}
	this.L.PushFront(key)
	this.M[key] = value
	if this.L.Len() > this.Len {
		ele := this.L.Back()
		this.L.Remove(ele)
		delete(this.M, ele.Value.(int))
	}
}

func findElement(ele *list.Element, value int) *list.Element {
	for ele != nil {
		if ele.Value.(int) == value {
			return ele
		}
		ele = ele.Next()
	}
	return nil
}
