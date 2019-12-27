package main

func main()  {
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
