package p05

import "fmt"

func sortWithEmpty(a []string, x string) int {
	return sortWithEmptyHelper(a, x, 0, len(a)-1)
}

func sortWithEmptyHelper(a []string, x string, left, right int) int {
	fmt.Println(left, right, (left+right)/2)
	if left > right {
		return -1
	}
	mid := (left + right) / 2
	amid := a[mid]

	if amid == "" {
		l := mid - 1
		r := mid + 1
		for {
			if l < left && r > right {
				return -1
			} else if r <= right && a[r] != "" {
				amid = a[r]
				mid = r
				break
			} else if l >= left && a[l] != "" {
				amid = a[l]
				mid = l
				break
			}
			l--
			r++
		}
	}
	if x > amid {
		return sortWithEmptyHelper(a, x, mid+1, right)
	} else if x < amid {
		return sortWithEmptyHelper(a, x, left, mid-1)
	}
	return mid
}
