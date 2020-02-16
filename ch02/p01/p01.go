package p01

import (
	"container/list"
)

// RemoveDuplicate removes duplicated elements from linked list
func RemoveDuplicate(l *list.List) *list.List {
	exists := make(map[interface{}]bool)
	var eNext *list.Element

	for e := l.Front(); e != nil; e = eNext {
		eNext = e.Next()
		if _, ok := exists[e.Value]; ok == true {
			l.Remove(e)
		} else {
			exists[e.Value] = true
		}
	}
	return l
}

// RemoveDuplicateLowMemory removes duplicated elements from linked list
// with low memory at the expense of time complexity (O(n^2))
func RemoveDuplicateLowMemory(l *list.List) *list.List {
	var eNext *list.Element

	for e := l.Front(); e != nil; e = eNext {
		eNext = e.Next()
		for e2 := l.Front(); e2 != e; e2 = e2.Next() {
			if e.Value == e2.Value {
				l.Remove(e)
				break
			}
		}
	}
	return l
}
