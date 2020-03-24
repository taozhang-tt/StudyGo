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
func main() {
	head := &ListNode{
		Val: -3,
	}
	fmt.Println(detectCycle2(head))
}

//借助了map集合，空间复杂度高
func detectCycle(head *ListNode) *ListNode {
	var m = make(map[*ListNode]bool)
	for {
		if head == nil {
			return nil
		}
		if _, ok := m[head]; ok {
			return head
		}
		m[head] = true
		head = head.Next
	}
}

/**
数据计算法
	推导过程：
		假设环外长度为 a，环长为 b，及链表头到环到入口处长度为 a
		设置快慢两个指针 fast, slow；fast步长为2，slow步长为1；当两者在环内相遇时，记录快指针走的步长为 f，慢指针走的步长为 s，且假设此时 fast比 slow 在环内多转了 n 圈（n>=1)
		此时有：f = 2s, f = s + nb; => f = 2nb, s = nb
		对于入环节点步长 k 满足 k = a + mb; (m >= 0)
		且当两个指针相遇时，slow 走过的步长为 s = nb;(n>0)，此时如果 slow 再走 a 步，则刚好位于入环节点处
*/
func detectCycle2(head *ListNode) *ListNode {
	var slow, fast = head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	if fast == nil || fast.Next == nil {	//判断是否有环
		return nil
	}
	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
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
