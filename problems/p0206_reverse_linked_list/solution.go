package p0206_reverse_linked_list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	// return sln_v1(head)
	return sln_v2(head)
	// return sln_v3(head)
}

func sln_v1(head *ListNode) *ListNode {
	var prev *ListNode = nil
	current := head

	for current != nil {
		temp := current.Next
		current.Next = prev
		prev = current
		current = temp
	}

	return prev
}

func sln_v2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := sln_v2(head.Next)
	head.Next.Next = head
	head.Next = nil

	return newHead
}

func sln_v3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	stack := []*ListNode{}
	curr := head
	for curr != nil {
		stack = append(stack, curr)
		curr = curr.Next
	}

	dummy := &ListNode{}
	curr = dummy
	for i := len(stack) - 1; i >= 0; i-- {
		curr.Next = stack[i]
		curr = curr.Next
	}
	curr.Next = nil

	return dummy.Next
}
