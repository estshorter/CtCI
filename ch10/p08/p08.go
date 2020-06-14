package p08

func showDuplicated(a []int) []int {
	bv := make([]byte, 32000/8+1)
	ret := make(map[int]struct{})
	for _, v := range a {
		b := bv[v/8]
		oneHot := byte(1 << (v % 8))
		if b&(oneHot) != 0 {
			ret[v] = struct{}{}
		}
		bv[v/8] |= oneHot
	}
	retAry := []int{}
	for k := range ret {
		retAry = append(retAry, k)
	}
	return retAry
}
