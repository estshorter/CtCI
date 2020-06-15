package p11

import (
	"math"
	"sort"
)

func sortPeakTrough(a []int) []int {
	sort.Ints(a)
	for i := 1; i < len(a); i += 2 {
		a[i-1], a[i] = a[i], a[i-1]
	}
	return a
}

func sortPeakTrough2(a []int) []int {
	for i := 1; i < len(a); i += 2 {
		biggestIndex := mapIndex(a, i-1, i, i+1)
		if i != biggestIndex {
			a[i], a[biggestIndex] = a[biggestIndex], a[i]
		}
	}
	return a
}

func mapIndex(array []int, a, b, c int) int {
	var aValue, bValue, cValue int
	if a >= 0 && a < len(array) {
		aValue = array[a]
	} else {
		aValue = math.MinInt32
	}
	if b >= 0 && b < len(array) {
		bValue = array[a]
	} else {
		bValue = math.MinInt32
	}
	if c >= 0 && c < len(array) {
		cValue = array[c]
	} else {
		cValue = math.MinInt32
	}
	maxValue := max(aValue, max(bValue, cValue))
	switch maxValue {
	case aValue:
		return a
	case bValue:
		return b
	case cValue:
		return c
	}
	return math.MinInt32
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
