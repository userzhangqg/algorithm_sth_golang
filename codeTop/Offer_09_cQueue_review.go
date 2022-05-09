package codeTop

type CQueueReview struct {
	inStack  []int
	outStack []int
}

func ConstructorCQueueReview() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTailReview(value int) {
	this.stack1 = append(this.stack1, value)
}

func (this *CQueue) DeleteHeadReview() int {
	if len(this.stack2) == 0 {
		if len(this.stack1) == 0 {
			return -1
		}
		for len(this.stack1) > 0 {
			val := this.stack1[len(this.stack1)-1]
			this.stack2 = append(this.stack2, val)
			this.stack1 = this.stack1[:len(this.stack1)-1]
		}
	}

	val := this.stack2[len(this.stack2)-1]
	this.stack2 = this.stack2[:len(this.stack2)-1]
	return val
}
