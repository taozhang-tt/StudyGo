package main

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
