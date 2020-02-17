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

	for i != 0 {
		q, r := divmod(i, 10)
		l.PushFront(r)
		i = q
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
func Sum(l1, l2 *list.List) *list.List {
	i1 := listToInt(l1)
	i2 := listToInt(l2)
	l := intToList(i1 + i2)
	return l
}

// SumReverse solves advanced problem 2.5
func SumReverse(l1, l2 *list.List) *list.List {
	i1 := listToIntReverse(l1)
	i2 := listToIntReverse(l2)
	l := intToListReverse(i1 + i2)
	return l
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// addLists solves problem 2.5
func addLists(l1, l2 *list.List) *list.List {
	l := list.New()
	lenMax := max(l1.Len(), l2.Len())
	e1, e2 := l1.Front(), l2.Front()
	for i, carry := 0, 0; i < lenMax || carry == 1; i++ {
		v1, v2 := 0, 0
		if i < l1.Len() {
			v1 = e1.Value
		}
		if i < l2.Len() {
			v2 = e2.Value
		}
		var val int
		carry, val = divmod(v1+v2+carry, 10)
		l.PushBack(val)
		e1, e2 = e1.Next, e2.Next
	}
	return l
}

// addListsAdvanced solves advanced problem 2.5
func addListsAdvanced(l1, l2 *list.List) *list.List {
	digitMax := max(l1.Len(), l2.Len()) + 1
	l := list.New()
	pushFrontAtDigit(l1.Front(), l2.Front(), l1.Len(), l2.Len(), l, digitMax)
	return l
}

func pushFrontAtDigit(e1, e2 *list.Element, l1Digit, l2Digit int, l *list.List, digit int) int {
	if digit == 1 {
		carry, val := divmod(e1.Value+e2.Value, 10)
		l.PushFront(val)
		return carry
	}
	e1Next, e2Next := e1, e2
	e1Value, e2Value := 0, 0
	if digit <= l1Digit {
		e1Next = e1.Next
		e1Value = e1.Value
	}
	if digit <= l2Digit {
		e2Next = e2.Next
		e2Value = e2.Value
	}
	carry := pushFrontAtDigit(e1Next, e2Next, l1Digit, l2Digit, l, digit-1)
	carry, val := divmod(e1Value+e2Value+carry, 10)
	if digit <= max(l1Digit, l2Digit) || val > 0 {
		l.PushFront(val)
	}
	return carry
}
