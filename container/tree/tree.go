package tree

import (
	"fmt"
)

//A Tree is a binary tree with integer values.
type Tree struct {
	Parent *Tree
	Left   *Tree
	Value  int
	Right  *Tree
}

// Insert inserts a new value to a tree
func (t *Tree) Insert(v int) {
	t.insertRecursive(v)
}

func (t *Tree) insertRecursive(v int) *Tree {
	if t == nil {
		return &Tree{nil, nil, v, nil}
	}
	if v < t.Value {
		t.Left = t.Left.insertRecursive(v)
	} else {
		t.Right = t.Right.insertRecursive(v)
	}
	return t
}

// InsertWithParent inserts a new value to a tree and specifies a parent tree
func (t *Tree) InsertWithParent(v int) {
	t.insertWithParentRecursive(v, nil)
}

func (t *Tree) insertWithParentRecursive(v int, p *Tree) *Tree {
	if t == nil {
		return &Tree{p, nil, v, nil}
	}
	if v < t.Value {
		t.Left = t.Left.insertWithParentRecursive(v, t)
	} else {
		t.Right = t.Right.insertWithParentRecursive(v, t)
	}
	return t
}

func (t *Tree) String() string {
	if t == nil {
		return "()"
	}
	s := ""
	if t.Left != nil {
		s += t.Left.String() + " "
	}
	s += fmt.Sprint(t.Value)
	if t.Right != nil {
		s += " " + t.Right.String()
	}
	return "(" + s + ")"
}
