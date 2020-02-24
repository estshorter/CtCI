package p10

import (
	"ctci/container/tree"
	"strconv"
	"strings"
)

//ContainsTree solves ch04-p10
func ContainsTree(t1, t2 *tree.Tree) bool {
	sb1 := &strings.Builder{}
	sb2 := &strings.Builder{}
	getOrderString(t1, sb1)
	getOrderString(t2, sb2)
	str1 := sb1.String()
	str2 := sb2.String()
	return strings.Contains(str1, str2)
}

func getOrderString(t *tree.Tree, sb *strings.Builder) {
	if t == nil {
		sb.WriteString("X ")
		return
	}
	sb.WriteString(strconv.Itoa(t.Value))
	sb.WriteRune(' ')
	getOrderString(t.Left, sb)
	getOrderString(t.Right, sb)
}

// ContainsTree2 solves ch04-p10
func ContainsTree2(t1, t2 *tree.Tree) bool {
	if t2 == nil {
		return true
	}
	return subTree(t1, t2)
}

func subTree(t1, t2 *tree.Tree) bool {
	if t1 == nil {
		return false
	} else if t1.Value == t2.Value && matchTree(t1, t2) {
		return true
	}
	return subTree(t1.Left, t2) || subTree(t1.Right, t2)
}

func matchTree(t1, t2 *tree.Tree) bool {
	if t1 == nil && t2 == nil {
		return true
	} else if t1 == nil || t2 == nil {
		return false
	} else if t1.Value != t2.Value {
		return false
	}
	return matchTree(t1.Left, t2.Left) && matchTree(t1.Right, t2.Right)
}
