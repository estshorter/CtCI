package p03

func magicFast(a []int) int {
	return magicFastAux(a, 0, len(a)-1)
}

func magicFastAux(a []int, start, end int) int {
	if start > end {
		return -1
	}
	mid := (start + end) / 2
	midVal := a[mid]
	if mid == midVal {
		return mid
	} else if mid < midVal {
		return magicFastAux(a, start, mid-1)
	}
	return magicFastAux(a, mid+1, end)

}

func magicFast2(a []int) int {
	return magicFast2Aux(a, 0, len(a)-1)
}

func magicFast2Aux(a []int, start, end int) int {
	if start > end {
		return -1
	}
	mid := (start + end) / 2
	midVal := a[mid]
	if mid == midVal {
		return mid
	}

	if left := magicFast2Aux(a, start, min(mid-1, midVal)); left >= 0 {
		return left
	}
	return magicFast2Aux(a, max(mid+1, midVal), end)
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
