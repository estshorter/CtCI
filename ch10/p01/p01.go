package p01

func merge(a, b []int, na int) []int {
	i := 0
	j := 0
	for j < len(b) {
		if i >= na || a[i+j] > b[j] {
			for k := na + j; k > i+j; k-- {
				a[k] = a[k-1]
			}
			a[i+j] = b[j]
			j++
		} else {
			i++
		}
	}
	return a
}

func merge2(a, b []int, na int) []int {
	idx := na + len(b) - 1
	i := na - 1
	j := len(b) - 1
	for j >= 0 {
		if i >= 0 && a[i] > b[j] {
			a[idx] = a[i]
			i--
		} else {
			a[idx] = b[j]
			j--
		}
		idx--
	}
	return a
}
