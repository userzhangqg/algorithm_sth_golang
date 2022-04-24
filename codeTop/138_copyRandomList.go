package codeTop

/*
138. 复制带随机指针的链表
https://leetcode-cn.com/problems/copy-list-with-random-pointer/

!!! TODO
未做出
思路：
1. map + 递归
*/

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	cachedNode := make(map[*Node]*Node)

	var deepCopy func(node *Node) *Node
	deepCopy = func(node *Node) *Node {
		if node == nil {
			return node
		}

		if v, ok := cachedNode[node]; ok {
			return v
		}
		newNode := &Node{Val: node.Val}
		cachedNode[node] = newNode
		newNode.Next = deepCopy(node.Next)
		newNode.Random = deepCopy(node.Random)
		return newNode
	}
	return deepCopy(head)
}
