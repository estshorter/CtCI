package p04

func sortP04(a []int, x int) int {
	idx := 1
	for a[idx] != -1 && a[idx] < x {
		idx *= 2
	}
	return binarySearch(a, x, 0, idx)
}

func binarySearch(a []int, x, left, right int) int {
	mid := (right + left) / 2
	if left > right {
		return -1
	}
	aMid := a[mid]
	if aMid > x || aMid == -1 {
		return binarySearch(a, x, left, mid-1)
	} else if aMid < x {
		return binarySearch(a, x, mid+1, right)
	} else {
		return mid
	}
}
