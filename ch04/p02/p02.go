package p02

import (
	"ctci/container/tree"
)

// CreateMinimalBST solves ch04-p02
func CreateMinimalBST(s []int) *tree.Tree {
	switch len(s) {
	case 1:
		return &tree.Tree{Value: s[0]}
	case 2:
		t := &tree.Tree{Value: s[1]}
		t.Left = &tree.Tree{Value: s[0]}
		return t
	default:
		mid := len(s) / 2
		t := &tree.Tree{Value: s[mid]}
		t.Left = CreateMinimalBST(s[0:mid])
		t.Right = CreateMinimalBST(s[mid+1 : len(s)])
		return t
	}
}
