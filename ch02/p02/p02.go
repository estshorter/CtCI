package p02

import (
	"container/list"
	"fmt"
)

func findLastKElement(l *list.List, k int) (*list.Element, error) {
	e1 := l.Front()
	e2 := l.Front()
	for i := 0; i < k; i++ {
		e1 = e1.Next()
		if e1 == nil {
			return nil, fmt.Errorf("k is larger than l.Len(). k:%v, l.Len():%v", k, l.Len())
		}
	}
	for e1 = e1.Next(); e1 != nil; e1 = e1.Next() {
		e2 = e2.Next()
	}
	return e2, nil
}
