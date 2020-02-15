package p06

func intToASCIIRune(i int) rune {
	return rune(i + 48)
}

// SimpleCompress compesses string
func SimpleCompress(s string) string {
	rs := []rune(s)
	out := make([]rune, 2*len(rs))
	iout := 0
	cnt := 0
	for idx := 0; idx < len(rs); idx++ {
		cnt++
		if idx+1 >= len(rs) || rs[idx] != rs[idx+1] {
			out[iout] = rs[idx]
			out[iout+1] = intToASCIIRune(cnt)
			iout += 2
			cnt = 0
		}
	}
	if iout < len(rs) {
		return string(out[:iout])
	}
	return s
}
