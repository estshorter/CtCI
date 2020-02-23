package p02

import (
	"ctci/container/tree"
)

// CreateMinimalBST solves ch04-p02
func CreateMinimalBST(s []int) *tree.Tree {
	if len(s) == 0 {
		return nil
	}
	mid := len(s) / 2
	t := &tree.Tree{Value: s[mid]}
	t.Left = CreateMinimalBST(s[0:mid])
	t.Right = CreateMinimalBST(s[mid+1 : len(s)])
	return t
}
