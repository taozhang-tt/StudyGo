package main

import "fmt"

/**
24. 两两交换链表中的节点
	https://leetcode-cn.com/problems/swap-nodes-in-pairs/
题目描述：
	给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
	你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
实例：
	给定 1->2->3->4, 你应该返回 2->1->4->3.
*/

func main() {
	head := generateList(1, 2, 3, 4, 5)
	printList(head)
	fmt.Println("----")
	head = swapPairs4(head)
	printList(head)
}

//非递归解法，添加空的头节点
func swapPairs(head *ListNode) *ListNode {
	var (
		prev = new(ListNode) //空的头节点
		temp = prev
	)
	prev.Next = head
	for {
		if temp.Next == nil || temp.Next.Next == nil {
			return prev.Next
		}
		left := temp.Next
		right := temp.Next.Next

		left.Next = right.Next
		right.Next = left
		temp.Next = right

		temp = temp.Next.Next
	}
	return prev.Next
}

//递归解法, 根据 swapPairs 思路改写
func swapPairs2(head *ListNode) *ListNode {
	var (
		prev = new(ListNode)
	)
	prev.Next = head
	if prev.Next == nil || prev.Next.Next == nil {
		return prev.Next
	}
	next := head.Next
	head.Next = swapPairs2(next.Next)
	next.Next = head
	prev.Next = next

	return prev.Next
}

//递归解法，根据 swapPairs2 优化
func swapPairs3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = swapPairs3(next.Next)
	next.Next = head
	return next
}

func swapPairs4(head *ListNode) *ListNode {
	var prev = new(ListNode)
	prev.Next = head
	var initPrev = prev
	for {
		if prev.Next == nil || prev.Next.Next == nil {
			return initPrev.Next
		}
		a, b := prev.Next, prev.Next.Next
		prev.Next, b.Next, a.Next = b, a, b.Next
		prev = a
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func generateList(param ...int) *ListNode {
	head := new(ListNode)
	head.Val = 0
	current := head
	for _, value := range param {
		next := new(ListNode)
		next.Val = value
		current.Next = next
		current = current.Next
	}
	return head.Next
}

func printList(head *ListNode) {
	for {
		if head == nil {
			return
		}
		fmt.Print(head.Val, ", ")
		head = head.Next
	}
}
