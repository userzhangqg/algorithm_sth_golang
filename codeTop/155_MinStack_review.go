package codeTop

import "math"

type MinStackReview struct {
	stack    []int
	minStack []int
}

func ConstructorReview() MinStackReview {
	return MinStackReview{
		stack:    []int{},
		minStack: []int{math.MaxInt64},
	}
}

func (this *MinStackReview) Push(val int) {
	this.stack = append(this.stack, val)
	this.minStack = append(this.minStack, min155Review(val, this.minStack[len(this.minStack)-1]))
}

func (this *MinStackReview) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStackReview) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStackReview) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

func min155Review(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
