package p08

import "ctci/container/tree"

// FindCommonAscendant solves ch04-p08
func FindCommonAscendant(t1, t2 *tree.Tree) *tree.Tree {
	d1 := depth(t1)
	d2 := depth(t2)

	var tLonger, tShorter *tree.Tree
	var diff, dmin int
	if d1 > d2 {
		tLonger = t1
		tShorter = t2
		diff = d1 - d2
		dmin = d2
	} else {
		tLonger = t2
		tShorter = t1
		diff = d2 - d1
		dmin = d1
	}
	for i := 0; i < diff; i++ {
		tLonger = tLonger.Parent
	}
	for i := 0; i < dmin; i++ {
		if tLonger.Parent == tShorter.Parent {
			return tLonger.Parent
		}
	}
	return nil
}

func depth(t *tree.Tree) int {
	cnt := 0
	t2 := t.Parent
	for t2 != nil {
		t2 = t2.Parent
		cnt++
	}
	return cnt
}

// FindCommonAscendant2 solves ch04-p08
func FindCommonAscendant2(root, p, q *tree.Tree) *tree.Tree {
	if t, isAncestor := commonAscendantRecurse(root, p, q); isAncestor {
		return t
	}
	return nil
}

func commonAscendantRecurse(root, p, q *tree.Tree) (*tree.Tree, bool) {
	if root == nil {
		return nil, false
	}
	if root == p && root == q {
		return root, true
	}

	rx, isAncestor := commonAscendantRecurse(root.Left, p, q)
	if isAncestor {
		return rx, true
	}

	ry, isAncestor := commonAscendantRecurse(root.Right, p, q)
	if isAncestor {
		return ry, true
	}

	if rx != nil && ry != nil {
		return root, true
	} else if root == p || root == q {
		isAncestor := rx != nil || ry != nil
		return root, isAncestor
	} else if rx != nil {
		return rx, false
	}
	return ry, false
}
