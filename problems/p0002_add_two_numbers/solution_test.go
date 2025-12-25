package p0002_add_two_numbers

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		name string
		l1   []int
		l2   []int
		want []int
	}{
		{
			name: "Example 1",
			l1:   []int{2, 4, 3},
			l2:   []int{5, 6, 4},
			want: []int{7, 0, 8},
		},
		{
			name: "Example 2",
			l1:   []int{0},
			l2:   []int{0},
			want: []int{0},
		},
		{
			name: "Example 3",
			l1:   []int{9, 9, 9, 9, 9, 9, 9},
			l2:   []int{9, 9, 9, 9},
			want: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
		{
			name: "Edge Case: Carry at the end",
			l1:   []int{5},
			l2:   []int{5},
			want: []int{0, 1},
		},
		{
			name: "Edge Case: Different lengths with carry propagation",
			l1:   []int{1},
			l2:   []int{9, 9},
			want: []int{0, 0, 1},
		},
		{
			name: "Edge Case: One list is much longer",
			l1:   []int{0},
			l2:   []int{1, 2, 3},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := sliceToLinkedList(tt.l1)
			l2 := sliceToLinkedList(tt.l2)
			got := linkedListToSlice(addTwoNumbers(l1, l2))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func sliceToLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	curr := head
	for i := 1; i < len(nums); i++ {
		curr.Next = &ListNode{Val: nums[i]}
		curr = curr.Next
	}
	return head
}

func linkedListToSlice(head *ListNode) []int {
	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	// Return empty slice for nil head to match tt.want if needed
	// But according to problem constraints, lists are non-empty.
	return res
}
