package p05

import (
	"ctci/container/list"
)

func pow(x, y int) int {
	z := 1
	for i := 0; i < y; i++ {
		z *= x
	}
	return z
}

func listToInt(l *list.List) int {
	ret := 0
	iter := 0
	for e := l.Front(); !l.IsConsumed(e); e = e.Next {
		ret += pow(10, iter) * e.Value
		iter++
	}
	return ret
}

func listToIntReverse(l *list.List) int {
	ret := 0
	iter := 0
	for e := l.Front(); !l.IsConsumed(e); e = e.Next {
		ret += pow(10, l.Len()-1-iter) * e.Value
		iter++
	}
	return ret
}

func intToList(i int) *list.List {
	l := list.New()
	if i == 0 {
		l.PushBack(0)
		return l
	}

	for i != 0 {
		q, r := divmod(i, 10)
		l.PushBack(r)
		i = q
	}
	return l
}

func intToListReverse(i int) *list.List {
	l := list.New()
	if i == 0 {
		l.PushBack(0)
		return l
	}

	var buf []int
	for i != 0 {
		q, r := divmod(i, 10)
		buf = append(buf, r)
		i = q
	}
	for j := len(buf) - 1; j >= 0; j-- {
		l.PushBack(buf[j])
	}
	return l
}

// https://stackoverflow.com/questions/43945675/division-with-returning-quotient-and-remainder
func divmod(numerator, denominator int) (quotient, remainder int) {
	quotient = numerator / denominator // integer division, decimals are truncated
	remainder = numerator % denominator
	return
}

// Sum solves problem 2.5
func Sum(l1 *list.List, l2 *list.List) *list.List {
	i1 := listToInt(l1)
	i2 := listToInt(l2)
	l := intToList(i1 + i2)
	return l
}

// SumReverse solves advanced problem 2.5
func SumReverse(l1 *list.List, l2 *list.List) *list.List {
	i1 := listToIntReverse(l1)
	i2 := listToIntReverse(l2)
	l := intToListReverse(i1 + i2)
	return l
}
