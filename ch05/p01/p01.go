package p01

// Insert soves ch05-p01
func Insert(n, m, i, j uint32) uint32 {
	n2 := clear(n, i, j)
	m2 := shift(m, i, j)
	ret := update(n2, m2)
	return ret
}

func clear(n, i, j uint32) uint32 {
	var mask uint32 = (1 << (j - i + 1)) - 1
	mask = mask << i
	return n &^ mask
}

func shift(m, i, j uint32) uint32 {
	return m << i
}

func update(n, m uint32) uint32 {
	return n | m
}
