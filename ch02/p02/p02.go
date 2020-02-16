package p02

import (
	"container/list"
	"fmt"
)

// FindLastKElement returns last kth elment of l
func FindLastKElement(l *list.List, k int) (*list.Element, error) {
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

// FindLastKElement2 returns last kth elment of l by using recursion
func FindLastKElement2(l *list.List, k int) *list.Element {
	e, _ := findLastKElementRecursion(l.Front(), k)
	return e
}

func findLastKElementRecursion(e *list.Element, k int) (*list.Element, int) {
	if e.Next() == nil {
		return nil, 0
	}
	e2, cnt := findLastKElementRecursion(e.Next(), k)
	cnt++
	if cnt == (k + 1) {
		return e.Next(), cnt
	}
	return e2, cnt
}
