package p02

// Point defines 2D points
type Point struct {
	x, y int
}

func getPath(maze [][]bool) []Point {
	if maze == nil || len(maze) == 0 {
		return nil
	}

	var path []Point
	memo := make(map[Point]bool)
	if ret, path := getPathInternal(maze, len(maze)-1, len(maze[0])-1, path, memo); ret {
		return path
	}
	return nil
}

func getPathInternal(maze [][]bool, row, col int, path []Point, memo map[Point]bool) (bool, []Point) {
	if col < 0 || row < 0 || !maze[row][col] {
		return false, path
	}
	p := Point{row, col}
	if _, ok := memo[p]; ok {
		return false, path
	}
	if isAtOrigin := (row == 0) && (col == 0); isAtOrigin {
		return true, append(path, p)
	}
	if ret, pathRet := getPathInternal(maze, row, col-1, path, memo); ret {
		return true, append(pathRet, p)
	}
	if ret, pathRet := getPathInternal(maze, row-1, col, path, memo); ret {
		return true, append(pathRet, p)
	}
	return false, path
}
