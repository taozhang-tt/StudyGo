package main

import "fmt"

func main() {
	head := generateList(1, 2, 3, 4, 5)
	printList(head)
	fmt.Println("----")
	head = reverseKGroup(head, 5)
	printList(head)
}

//给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
//
//k 是一个正整数，它的值小于或等于链表的长度。
//
//如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
//
//示例 :
//
//给定这个链表：1->2->3->4->5
//
//当 k = 2 时，应当返回: 2->1->4->3->5
//
//当 k = 3 时，应当返回: 3->2->1->4->5
func reverseKGroup(head *ListNode, k int) *ListNode {
	var (
		fake  = new(ListNode) //一个空的头节点
		prev  = fake          //记录待反转的那部分链表的前一个节点
		start = prev          //记录待反转的那部分链表的开始节点
		end   = prev          //记录待反转的那部分链表的结束节点
	)
	fake.Next = head
	for {
		for i := 0; i < k; i++ {
			end = end.Next
			if end == nil {
				break
			}
		}
		if end == nil { //如果不足 K 个，则结束
			break
		}
		//取出需要反转的这部分链表
		next := end.Next //记录待反转部分链表的下一个节点
		end.Next = nil	//将待反转的这部分链表与后面的链表断开
		start = prev.Next	//记录待反转部分链表的下一个节点
		prev.Next = reverseK(prev.Next)	//反转这部分链表
		start.Next = next //将反转后的链表与前一部分拼接

		prev = start
		end = start
	}
	return fake.Next
}

func reverseK(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	finalHead := reverseK(head.Next)
	head.Next.Next = head
	head.Next = nil
	return finalHead
}
