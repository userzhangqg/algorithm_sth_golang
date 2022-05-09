package codeTop

/*
剑指offer 09. 用两个栈来实现队列
https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/

用两个栈来实现一个队列。队列的声明如下，请实现它的两个函数appendTail和deleteHead，分别完成在队列尾部插入整数和在队列头部删除整数的功能。
若队列中没有元素，deleteHead 操作返回-1

!!! TODO
做过思路不熟
思路：
1. 双栈入栈出栈
*/

type CQueue struct {
	stack1 []int
	stack2 []int
}

func ConstructorOffer09() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.stack1 = append(this.stack1, value)
}

func (this *CQueue) DeleteHead() int {
	// stack2 为零时将stack1导入，导入完成此时队列首为栈顶。否则直接出栈，栈顶元素即为队列首
	if len(this.stack2) == 0 {
		if len(this.stack1) == 0 {
			return -1
		}
		for len(this.stack1) > 0 {
			val := this.stack1[len(this.stack1)-1]
			this.stack1 = this.stack1[:len(this.stack1)-1]
			this.stack2 = append(this.stack2, val)
		}
	}

	ans := this.stack2[len(this.stack2)-1]
	this.stack2 = this.stack2[:len(this.stack2)-1]
	return ans
}
