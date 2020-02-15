package p08

func zeronize(mat [][]int) [][]int {
	m := len(mat)
	n := len(mat[0])
	zeroColMap := make(map[int]int)
	zeroRowMap := make(map[int]int)
	for col := 0; col < m; col++ {
		for row := 0; row < n; row++ {
			if mat[col][row] == 0 {
				zeroColMap[col]++
				zeroRowMap[row]++
			}
		}
	}
	for col := range zeroColMap {
		for row := 0; row < n; row++ {
			mat[col][row] = 0
		}
	}
	for row := range zeroRowMap {
		for col := 0; col < m; col++ {
			mat[col][row] = 0
		}
	}

	return mat
}
