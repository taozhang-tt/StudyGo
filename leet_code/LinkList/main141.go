package main

func main() {

}

//利用快慢指针
func hasCycle(head *ListNode) bool {
	var fast, slow = head, head
	for {
		if fast == nil || fast.Next ==  nil {
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
}

//利用map做记录
func hasCycle2(head *ListNode) bool {
	var m = make(map[*ListNode]bool, 0)
	for {
		if head ==  nil {
			return false
		}
		if _, ok := m[head]; ok {
			return true
		}
		m[head] = true
		head = head.Next
	}
}
