package p06

import "ctci/container/tree"

// InOrderSucc solves ch04-p06
func InOrderSucc(t *tree.Tree) *tree.Tree {
	if t.Right != nil {
		return getMostLeftNode(t.Right)
	}
	t2 := t
	p := t.Parent
	for p != nil && p.Left != t2 {
		t2 = p
		p = p.Parent
	}
	return p
}

func getMostLeftNode(t *tree.Tree) *tree.Tree {
	for t.Left != nil {
		t = t.Left
	}
	return t
}
