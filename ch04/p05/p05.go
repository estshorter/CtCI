package p05

import (
	"ctci/container/tree"
)

const maxInt = int(^uint(0) >> 1)
const minInt = -maxInt - 1

// isBST checks if a tree is BST
func isBST(t *tree.Tree) bool {
	return isBSTRecursive(t, minInt, maxInt)
}

func isBSTRecursive(t *tree.Tree, min int, max int) bool {
	if t == nil {
		return true
	}
	if t.Value <= min || t.Value > max {
		return false
	}
	if ret := isBSTRecursive(t.Left, min, t.Value); !ret {
		return false
	}
	if ret := isBSTRecursive(t.Right, t.Value, max); !ret {
		return false
	}
	return true
}
