package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	curr := dummyHead
	carry := 0
	p := l1
	q := l2

	for p != nil || q != nil {
		sum := carry

		if p != nil {
			sum += p.Val
			p = p.Next
		}

		if q != nil {
			sum += q.Val
			q = q.Next
		}

		carry = sum / 10
		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next

	}
	if carry > 0 {
		curr.Next = &ListNode{Val: carry}
	}
	return dummyHead.Next
}
