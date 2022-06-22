package codeTopReview

/*
146. LRU缓存
https://leetcode.cn/problems/lru-cache/

思路：
1. 双向链表
2. 伪头部，尾部减少判断
3. 抽象move，init等操作
4. size > cap 时移除尾部
*/

type CacheLinkedNode struct {
	key, value int
	prev, next *CacheLinkedNode
}

type LRUCache struct {
	size, capacity int
	cache          map[int]*CacheLinkedNode
	head, tail     *CacheLinkedNode
}

func initNode(key, value int) *CacheLinkedNode {
	return &CacheLinkedNode{key: key, value: value}
}

func Constructor(capacity int) LRUCache {
	// 使用伪头，尾 避免插入时判断是否有头尾
	head := initNode(0, 0)
	tail := initNode(0, 0)
	head.next = tail
	head.prev = tail
	tail.next = head
	tail.prev = head
	return LRUCache{capacity: capacity, cache: map[int]*CacheLinkedNode{}, head: head, tail: tail}
}

func (this *LRUCache) addToHead(node *CacheLinkedNode) {
	//this.head.next = node
	//node.prev = this.head
	// 错误代码
	//this.tail.prev = node
	//node.next = this.tail

	node.prev = this.head
	node.next = this.head.next

	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) moveToHead(node *CacheLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeNode(node *CacheLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) removeTail() *CacheLinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	} else {
		node := this.cache[key]
		this.moveToHead(node)
		return node.value
	}
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; !ok {
		node := initNode(key, value)
		this.addToHead(node)
		this.cache[key] = node
		this.size++
		if this.size > this.capacity {
			removed := this.removeTail()
			delete(this.cache, removed.key)
			this.size--
		}
	} else {
		node := this.cache[key]
		node.value = value
		this.moveToHead(node)
	}
}

/*
type CacheListNode struct {
	Val int
	Pre *CacheListNode
	Next *CacheListNode
}


type LRUCache struct {
	capacity int
	head *CacheListNode
	cache map[int]*CacheListNode
}


func Constructor(capacity int) LRUCache {
	return LRUCache{capacity: capacity}
}

func (this *LRUCache)updateCache(cur *CacheListNode)  {
	cur.Pre = cur.Next
	pre := this.head.Pre

	this.head.Pre = cur
	cur.Next = this.head

	pre.Next = cur
	cur.Pre = pre
	this.head = cur
}


func (this *LRUCache) Get(key int) int {
	cur := this.cache[key]
	if cur != nil {
		this.updateCache(cur)
		return cur.Val
	}else {
		return -1
	}
}


func (this *LRUCache) Put(key int, value int)  {
	cur := this.cache[key]
	if cur != nil {
		this.updateCache(cur)
		cur.Val = value
	}else {
		if len(this.cache) == this.capacity {
			this.DeleteTail()
		}
		this.InsertHead(key, value)
	}
}

func (this *LRUCache) DeleteTail()  {
	this.head.Pre = this.head.Pre.Pre
	this.head.Pre.Pre.Next = this.head
	this.capacity--
}

func (this *LRUCache) InsertHead(key int, value int)  {
	cur := &CacheListNode{Val: value, Pre: this.head.Pre, Next: this.head}
	this.cache[key] = cur
	this.capacity++
	pre := this.head.Pre
	this.head.Pre = cur
	pre.Next = cur
	this.head = cur
}
*/
