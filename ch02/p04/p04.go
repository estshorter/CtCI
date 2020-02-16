package p04

import (
	"ctci/container/list"
)

// Partition solves ch02-p04
func Partition(l *list.List, x int) *list.List {
	head := list.New()
	tail := list.New()

	var eNext *list.Element
	for e := l.Front(); !l.IsConsumed(e); e = eNext {
		eNext = e.Next
		if e.Value < x {
			head.PushBackElement(e)
		} else {
			tail.PushBackElement(e)
		}
	}

	head.PushBackList(tail)
	return head
}
