package p12

// Point defines 2D points
type Point struct {
	Row, Col int
}

const gridsize = 8

func placeQueen() [][gridsize]int {
	return placeQueenHelper(0, [gridsize]int{}, [][gridsize]int{})
}

func placeQueenHelper(row int, columns [gridsize]int, placements [][gridsize]int) [][gridsize]int {
	if row == gridsize {
		return append(placements, columns)
	}
	for c := 0; c < gridsize; c++ {
		if checkValid(columns, row, c) {
			columns[row] = c
			placements = placeQueenHelper(row+1, columns, placements)
		}
	}
	return placements
}

func checkValid(columns [gridsize]int, row, col int) bool {
	for r2 := 0; r2 < row; r2++ {
		c2 := columns[r2]
		if col == c2 {
			return false
		}
		var cDist int
		if c2 > col {
			cDist = c2 - col
		} else {
			cDist = col - c2
		}
		rDist := row - r2
		if cDist == rDist {
			return false
		}
	}
	return true
}
