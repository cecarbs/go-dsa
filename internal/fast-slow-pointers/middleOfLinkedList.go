package fastslowpointers

func MiddleOfLinkedList(head *ListNode) *ListNode {
	slow := head
	fast := head

	for fast.Next != nil && fast != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
