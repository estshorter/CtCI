package p03

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
