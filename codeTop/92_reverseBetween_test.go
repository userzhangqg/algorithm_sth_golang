package codeTop

import (
	"fmt"
	"testing"
)

func TestReverseBetween(t *testing.T) {
	head := &ListNode{Val: 3, Next: &ListNode{Val: 5}}
	res := reverseBetween(head, 1, 2)
	fmt.Println(res)
}
