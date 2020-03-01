package p07

// SwapOddEvenBits solves ch05 p07
func SwapOddEvenBits(x uint) uint {
	return ((x & 0xaaaaaaaaaaaaaaaa) >> 1) | ((x & 0x5555555555555555) << 1)
}
