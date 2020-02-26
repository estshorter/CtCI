package p01

import "fmt"

// Insert soves ch05-p01
func Insert(n, m, i, j uint32) uint32 {
	n2 := clear(n, i, j)
	m2 := makeMask(m, i, j)
	ret := update(n2, m2)
	fmt.Printf("%b %b %b\n", n2, m2, ret)
	return ret
}

func clear(n, i, j uint32) uint32 {
	var mask uint32 = (1 << (j - i + 1)) - 1
	mask = mask << i
	return n &^ mask
}

func makeMask(m, i, j uint32) uint32 {
	return m << i
}

func update(n, m uint32) uint32 {
	return n | m
}
