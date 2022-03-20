package codeTop

/*
剑指 Offer 22. 链表中倒数第k个节点

输入一个链表，输出该链表中倒数第k个节点。为了符合大多数人的习惯，本题从1开始计数，即链表的尾节点是倒数第1个节点。

例如，一个链表有 6 个节点，从头节点开始，它们的值依次是 1、2、3、4、5、6。这个链表的倒数第 3 个节点是值为 4 的节点。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof

思路：
遍历获取总数，再次遍历head，当剩余=k时返回head
*/

func getKthFromEnd(head *ListNode, k int) *ListNode {
	cur := head
	total := 0

	for cur != nil {
		total++
		cur = cur.Next
	}

	//n := 1
	//for n+k-1 != total {
	//	n++
	//	head = head.Next
	//}
	for total > k {
		total--
		head = head.Next
	}
	return head
}
