package main

/**
876. 链表的中间结点
	https://leetcode-cn.com/problems/middle-of-the-linked-list/
题目描述：
	给定一个带有头结点 head 的非空单链表，返回链表的中间结点。
	如果有两个中间结点，则返回第二个中间结点。
*/

import "fmt"

func main()  {
	var head = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: nil,
		},
	}
	ret := middleNode(head)
	fmt.Println(ret.Val)
}

func middleNode(head *ListNode) *ListNode {
	var slow, fast = head, head
	for fast.Next != nil {
		slow = slow.Next
		if fast.Next.Next == nil {
			return slow
		}
		fast = fast.Next.Next
	}

	return slow
}

type ListNode struct {
	Val  int
	Next *ListNode
}