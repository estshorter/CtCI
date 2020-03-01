package p08

// DrawLine solves ch05 p08
func DrawLine(screen []byte, width, x1, x2, y int) []byte {
	if x2 < x1 {
		return DrawLine(screen, width, x2, x1, y)
	}

	firstByte := x1 / 8
	posBitX1 := x1 % 8
	lastByte := x2 / 8
	posBitX2 := x2 % 8

	offset := width / 8 * y

	mask1 := 0xFF >> posBitX1
	mask2 := (0xFF << (7 - posBitX2))
	if firstByte == lastByte {
		screen[firstByte+offset] |= byte(mask1 & mask2)
		return screen
	}

	for i := firstByte + 1; i < lastByte; i++ {
		screen[i+offset] = 0xFF
	}
	screen[firstByte+offset] |= byte(mask1)
	screen[lastByte+offset] |= byte(mask2)
	return screen
}
