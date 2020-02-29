package p03

//CountSucc solves ch05 p03
func CountSucc(v uint) int {
	cntMax := 0
	a := 0
	b := 0
	allOne := true
	for ; v != 0; v >>= 1 {
		if v&1 == 0 {
			cntMax = max(cntMax, b+a+1)
			b = a
			a = 0
			allOne = false
		} else {
			a++
		}
	}
	if allOne {
		cntMax = max(cntMax, a)
	} else {
		cntMax = max(cntMax, b+a+1)
	}
	return cntMax
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
