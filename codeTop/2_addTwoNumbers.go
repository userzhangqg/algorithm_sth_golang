package codeTop

/*
2. 两数相加
https://leetcode-cn.com/problems/add-two-numbers/
给你两个非空的链表，表示两个非负的整数。他们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。
请你将两个数相加，并以相同形式返回一个表示和的链表。
你可以假设除了数字0以外，这两个数都不会以0开头。


！！！ 看提示后作出
思路：
从前向后进位推导，注意最后进位
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy

	var tmp int

	for l1 != nil || l2 != nil {
		var v int
		v = v + tmp
		if l1 != nil {
			v = l1.Val + v
			l1 = l1.Next
		}
		if l2 != nil {
			v = l2.Val + v
			l2 = l2.Next
		}

		tmp = v / 10
		cur.Next = &ListNode{Val: v % 10}
		cur = cur.Next
	}
	if tmp == 1 {
		cur.Next = &ListNode{Val: 1}
	}
	return dummy.Next
}
