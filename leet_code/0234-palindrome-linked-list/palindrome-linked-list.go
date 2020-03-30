package main

import "fmt"

func main() {
	var head = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}
	fmt.Println(isPalindrome(head))

	head = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
	}
	fmt.Println(isPalindrome(head))
}

/**
双指针做法
	先遍历一遍链表，将元素全部存入切片中
	再使用双端指针判断切片中的元素是否回文
*/
func isPalindrome(head *ListNode) bool {
	s := make([]int, 0)
	for head != nil {
		s = append(s, head.Val)
		head = head.Next
	}
	l, r := 0, len(s)-1
	for l <= r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

/**
反转链表的后半部分，然后比较和前半部分是否一致
*/
func isPalindrome2(head *ListNode) bool {

}

func reverList(head *ListNode) *ListNode {
	if head != nil 
}

type ListNode struct {
	Val  int
	Next *ListNode
}
