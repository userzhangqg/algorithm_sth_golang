package codeTop

func copyRandomListReview(head *Node) *Node {
	cacheNode := make(map[*Node]*Node)

	var deepCopy func(node *Node) *Node
	deepCopy = func(node *Node) *Node {
		if node == nil {
			return node
		}
		if n, ok := cacheNode[node]; ok {
			return n
		}
		newNode := &Node{}
		newNode.Val = node.Val
		cacheNode[node] = newNode
		newNode.Next = deepCopy(node.Next)
		newNode.Random = deepCopy(node.Random)
		return newNode
	}
	return deepCopy(head)
}
