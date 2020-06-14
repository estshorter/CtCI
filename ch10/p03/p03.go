package p03

import (
	"fmt"
	"math"
)

func binarySearchRotate(a []int, v int) int {
	minIdx := argmin(a)
	return binarySearchRotateHelper(a, v, minIdx, (len(a)-1+minIdx)%len(a), minIdx)
}

func binarySearchRotateHelper(a []int, v, low, high, startIdx int) int {
	lowVirtual := idxRealToVirtual(low, startIdx, len(a))
	highVirtual := idxRealToVirtual(high, startIdx, len(a))

	for highVirtual >= lowVirtual {
		mid := idxVirtualToReal((highVirtual+lowVirtual)/2, startIdx, len(a))
		fmt.Println(lowVirtual, highVirtual, mid)
		if v < a[mid] {
			high = (mid - 1) % len(a)
			if high < 0 {
				high += len(a)
			}
		} else if v > a[mid] {
			low = (mid + 1) % len(a)
		} else {
			return mid
		}
		lowVirtual = idxRealToVirtual(low, startIdx, len(a))
		highVirtual = idxRealToVirtual(high, startIdx, len(a))
	}
	return -1
}

func idxVirtualToReal(vIdx, startIdx, lenSlice int) int {
	return (vIdx + startIdx) % lenSlice
}

func idxRealToVirtual(rIdx, startIdx, lenSlice int) int {
	temp := (rIdx - startIdx)
	if temp < 0 {
		temp += lenSlice
	}
	return temp
}

func argmin(a []int) int {
	minVal := math.MaxInt64
	idx := -1
	for i, v := range a {
		if v < minVal {
			minVal = v
			idx = i
		}
	}
	return idx
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func binarySearchRotate2(a []int, v int) int {
	return binarySearchRotateHelper2(a, v, 0, len(a)-1)
}

func binarySearchRotateHelper2(a []int, x, left, right int) int {
	mid := (left + right) / 2
	if x == a[mid] {
		return mid
	}
	if right < left {
		return -1
	}
	if a[left] < a[mid] {
		if x >= a[left] && x < a[mid] {
			return binarySearchRotateHelper2(a, x, left, mid-1)
		}
		return binarySearchRotateHelper2(a, x, mid+1, right)
	} else if a[mid] < a[left] {
		if x > a[mid] && x <= a[right] {
			return binarySearchRotateHelper2(a, x, mid+1, right)
		}
		return binarySearchRotateHelper2(a, x, left, mid-1)
	} else if a[left] == a[mid] {
		if a[mid] != a[right] {
			return binarySearchRotateHelper2(a, x, mid+1, right)
		}
		result := binarySearchRotateHelper2(a, x, left, mid-1)
		if result == -1 {
			return binarySearchRotateHelper2(a, x, mid+1, right)
		}
		return result
	}
	return -1
}
