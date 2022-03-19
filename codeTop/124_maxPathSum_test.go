package codeTop

import (
	"fmt"
	"testing"
)

func Test_maxPathSum(t *testing.T) {
	root := &TreeNode{Val: -2, Left: &TreeNode{Val: 1}}
	res := maxPathSum(root)
	fmt.Println(res)
}
