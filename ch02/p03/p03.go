package p03

import "ctci/container/list"

// DeleteElement return list ..
func DeleteElement(l *list.List, mark *list.Element) *list.List {
	mark.Value = mark.Next.Value
	mark.Next = mark.Next.Next
	l.N--
	return l
}
