package codeTop

/*
给你两个单链表的头节点headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。

图示两个链表在节点 c1 开始相交：

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/intersection-of-two-linked-lists

思路：
1. 哈希表
2. 双指针
错的人就算走过了对方的路也还是会错过
*/

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	hashMap := map[*ListNode]bool{}
	for headA != nil {
		hashMap[headA] = true
		headA = headA.Next
	}
	for headB != nil {
		if _, ok := hashMap[headB]; ok {
			return headB
		}
		headB = headB.Next
	}
	return nil
}
