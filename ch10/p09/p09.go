package p09

func searchSortedMatrix(a [][]int, x int) (int, int) {
	row := 0
	col := len(a[0]) - 1
	for row < len(a) && col >= 0 {
		if a[row][col] == x {
			return row, col
		} else if a[row][col] > x {
			col--
		} else {
			row--
		}
	}
	return -1, -1
}
