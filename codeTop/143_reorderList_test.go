package codeTop

import (
	"testing"
)

func TestReorder(t *testing.T) {
	head := &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				3,
				&ListNode{
					4,
					&ListNode{
						5,
						nil,
					},
				},
			},
		},
	}
	reorderList(head)
}
