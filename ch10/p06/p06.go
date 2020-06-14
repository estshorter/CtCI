package p06

func count(a []int) []int {
	maxInt8 := 127

	bv := make([]byte, maxInt8/8+1)
	for _, v := range a {
		hotVector := byte(1 << (v % 8))
		bv[v/8] |= hotVector
	}

	ret := []int{}
	for i := 0; i < int(maxInt8); i++ {
		b := bv[i/8]
		if b&byte(1<<(i%8)) == 0 {
			ret = append(ret, i)
		}
	}
	return ret
}
