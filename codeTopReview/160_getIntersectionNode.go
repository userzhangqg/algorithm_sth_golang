package codeTopReview

/*
160. 相交链表

// 双指针
时间复杂度：O(m+n)O(m+n)，其中 mm 和 nn 是分别是链表 \textit{headA}headA 和 \textit{headB}headB 的长度。需要遍历两个链表各一次。

空间复杂度：O(m)O(m)，其中 mm 是链表 \textit{headA}headA 的长度。需要使用哈希集合存储链表 \textit{headA}headA 中的全部节点。

// hashMap
时间复杂度：O(m+n)O(m+n)，其中 mm 和 nn 是分别是链表 \textit{headA}headA 和 \textit{headB}headB 的长度。两个指针同时遍历两个链表，每个指针遍历两个链表各一次。

空间复杂度：O(1)O(1)。

*/

func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	hashMap := make(map[*ListNode]bool)

	for headA != nil {
		hashMap[headA] = true
		headA = headA.Next
	}

	for headB != nil {
		if _, ok := hashMap[headB]; ok {
			return headB
		} else {
			headB = headB.Next
		}
	}
	return nil
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 双指针
	if headA == nil || headB == nil {
		return nil
	}
	curA, curB := headA, headB
	for curA != curB {
		if curA != nil {
			curA = curA.Next
		} else {
			curA = headB
		}

		if curB != nil {
			curB = curB.Next
		} else {
			curB = headA
		}
	}
	return curA
}
