package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	a := ListNode{1, &ListNode{Val: 2}}
	b := ListNode{1, &ListNode{Val: 3}}
	fmt.Println(a == b)
}
