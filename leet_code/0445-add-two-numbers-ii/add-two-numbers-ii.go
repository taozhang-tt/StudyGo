package main

import "fmt"

/**
445. 两数相加 II
	https://leetcode-cn.com/problems/add-two-numbers-ii/
题目描述
	给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。
	你可以假设除了数字 0 之外，这两个数字都不会以零开头。
进阶：
	如果输入链表不能修改该如何处理？换句话说，你不能对列表中的节点进行翻转。
示例：
	输入：(7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
	输出：7 -> 8 -> 0 -> 7
*/

func main()  {
	var l1, l2, l3 *ListNode
	l1 = &ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 3,
					Next: nil,
				},
			},
		},
	}
	l2 = &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
				Next: nil,
			},
		},
	}
	l3 = addTwoNumbers(l1, l2)
	for l3 != nil {
		fmt.Println(l3.Val)
		l3 = l3.Next
	}

	l1 = new(ListNode)
	l2 = new(ListNode)
	l3 = addTwoNumbers(l1, l2)
	for l3 != nil {
		fmt.Println(l3.Val)
		l3 = l3.Next
	}
}

type ListNode struct {
    Val int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1 := make([]int, 0)
	stack2 := make([]int, 0)
	for l1 != nil {
		stack1 = append([]int{l1.Val}, stack1...)
		l1 = l1.Next
	}
	for l2 != nil {
		stack2 = append([]int{l2.Val}, stack2...)
		l2 = l2.Next
	}
	head := new(ListNode)
	carry := 0
	for len(stack1) > 0 || len(stack2) > 0 || carry > 0 {
		if len(stack1) > 0 {
			carry += stack1[0]
			stack1 = stack1[1:]
		}
		if len(stack2) > 0 {
			carry += stack2[0]
			stack2 = stack2[1:]
		}
		head.Val = carry % 10
		prev := new(ListNode)
		prev.Next = head
		head = prev
		carry /= 10
	}
	return head.Next
}