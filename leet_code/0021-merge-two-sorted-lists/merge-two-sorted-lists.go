package main

import "fmt"

/**
21. 合并两个有序链表
	https://leetcode-cn.com/problems/merge-two-sorted-lists/
题目描述：
	将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
示例
	输入：1->2->4, 1->3->4
	输出：1->1->2->3->4->4
*/

func main() {
	head1 := generateList(1, 3, 5)
	head2 := generateList(2, 4, 6)

	head := mergeTwoLists(head1, head2)
	printList(head)
}

//递归解法
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
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
