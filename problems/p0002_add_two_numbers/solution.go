package p0002_add_two_numbers

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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode = nil

	var c int = 0
	var prev *ListNode = nil
	for l1 != nil || l2 != nil {
		var cur *ListNode = new(ListNode)
		cur.Val = c
		c = 0
		if l1 != nil {
			cur.Val += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			cur.Val += l2.Val
			l2 = l2.Next
		}

		c = cur.Val / 10
		cur.Val = cur.Val % 10

		if head == nil {
			head = cur
		}

		if prev != nil {
			prev.Next = cur
		}
		prev = cur
	}

	if c > 0 {
		prev.Next = new(ListNode)
		prev.Next.Val = c
	}

	return head
}
