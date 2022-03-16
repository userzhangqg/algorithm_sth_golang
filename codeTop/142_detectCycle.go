package codeTop

/*
142. 环形链表II
https://leetcode-cn.com/problems/linked-list-cycle-ii/

给定一个链表的头结点head，返回链表开始如环的第一个节点。如果链表无环，则返回null。
如果链表中有某个节点，可以通过连续跟踪next指针再次到达，则链表中存在环。
*/

func detectCycle(head *ListNode) *ListNode {
	hashMap := make(map[*ListNode]bool)
	for head != nil {
		if hashMap[head] {
			return head
		} else {
			hashMap[head] = true
		}
		head = head.Next
	}
	return nil
}
