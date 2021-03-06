package main

import "fmt"

/**
206. 反转一个单链表。
	https://leetcode-cn.com/problems/reverse-linked-list/
题目描述：
	反转一个单链表。
示例:
	输入: 1->2->3->4->5->NULL
	输出: 5->4->3->2->1->NULL
*/
func main() {
	head := generateList(1, 2, 3, 4, 5)
	head = reverseList4(head)
	printList(head)
}

//添加了一个空的头指针
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	var next *ListNode
	for {
		if head == nil {
			return prev
		}
		next = head.Next
		head.Next = prev
		prev = head
		head = next
	}
}

//不添加空的头指针，需要先处理头节点的反转
func reverseList2(head *ListNode) *ListNode {
	var (
		next     *ListNode
		nextNext *ListNode
	)
	//处理头节点
	if head == nil {
		return head
	}
	next = head.Next
	head.Next = nil

	for {
		if next == nil {
			return head
		}
		nextNext = next.Next
		next.Next = head
		head = next
		next = nextNext
	}
}

//递归做法
func reverseList3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	finalHead := reverseList3(head.Next)
	head.Next.Next = head
	head.Next = nil
	return finalHead
}

func reverseList4(head *ListNode) *ListNode {
	var prev *ListNode
	for {
		if head == nil {
			return prev
		}
		head.Next, prev, head = prev, head, head.Next
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
