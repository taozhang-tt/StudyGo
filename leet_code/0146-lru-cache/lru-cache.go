package main

import (
	"container/list"
	"fmt"
)

/**
146. LRU缓存机制
	https://leetcode-cn.com/problems/lru-cache/
题目描述：
	运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。
	获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
	写入数据 put(key, value) - 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。
进阶:
	你是否可以在 O(1) 时间复杂度内完成这两种操作？
示例:
	LRUCache cache = new LRUCache( 2 【缓存容量】 );
	cache.put(1, 1);
	cache.put(2, 2);
	cache.get(1);       // 返回  1
	cache.put(3, 3);    // 该操作会使得密钥 2 作废
	cache.get(2);       // 返回 -1 (未找到)
	cache.put(4, 4);    // 该操作会使得密钥 1 作废
	cache.get(1);       // 返回 -1 (未找到)
	cache.get(3);       // 返回  3
	cache.get(4);       // 返回  4
*/

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
