package p04

import "ctci/container/tree"

// IsBalanced solves ch04-p04
func IsBalanced(root *tree.Tree) bool {
	_, ret := isBalancedRecursive(root)
	return ret
}

func isBalancedRecursive(tr *tree.Tree) (int, bool) {
	if tr == nil {
		return -1, true
	}
	leftHeight, ret := isBalancedRecursive(tr.Left)
	if ret == false {
		return 0, false
	}
	rightHeight, ret := isBalancedRecursive(tr.Right)
	if ret == false {
		return 0, false
	}
	height := max(leftHeight, rightHeight) + 1
	return height, abs(leftHeight-rightHeight) <= 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
