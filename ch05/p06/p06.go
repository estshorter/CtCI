package p06

func convertBit(a, b uint) int {
	var c1 uint = 0
	for xor := a ^ b; xor != 0; xor = xor & (xor - 1) {
		c1++
	}
	return int(c1)
}
