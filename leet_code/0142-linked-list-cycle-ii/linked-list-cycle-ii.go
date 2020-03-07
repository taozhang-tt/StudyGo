package main

import "fmt"

/**
142. 环形链表 II
	https://leetcode-cn.com/problems/linked-list-cycle-ii/
题目描述：
	给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
	为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。
	说明：不允许修改给定的链表。
示例：
	输入：head = [3,2,0,-4], pos = 1
	输出：tail connects to node index 1
	解释：链表中有一个环，其尾部连接到第二个节点。
*/
func main()  {

}

//借助了map集合，空间复杂度高
func detectCycle(head *ListNode) *ListNode {
	var m = make(map[*ListNode]bool)
	for {
		if head ==  nil {
			return nil
		}
		if _, ok := m[head]; ok {
			return head
		}
		m[head] = true
		head = head.Next
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