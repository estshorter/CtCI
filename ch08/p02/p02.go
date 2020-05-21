package p02

import (
	"github.com/golang-collections/collections/stack"
)

// Point defines 2D points
type Point struct {
	Row, Col int
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

func getPathStack(maze [][]bool) []Point {
	if maze == nil || len(maze) == 0 {
		return nil
	}

	memo := make(map[Point]bool)
	s := stack.New()
	sPath := stack.New()
	s.Push(Point{len(maze) - 1, len(maze[0]) - 1})
	for {
		p, ok := s.Pop().(Point)
		if !ok {
			return nil
		}

		if p.Row < 0 || p.Col < 0 || !maze[p.Row][p.Col] {
			if (sPath.Peek().(Point).Row - p.Row) == 0 {
				continue
			}
			n := s.Peek().(Point).Row - p.Row + 1
			for i := 0; i < n; i++ {
				sPath.Pop()
			}
			continue
		}
		if _, ok := memo[p]; ok {
			if (sPath.Peek().(Point).Row - p.Row) == 0 {
				continue
			}
			n := s.Peek().(Point).Row - p.Row + 1
			for i := 0; i < n; i++ {
				sPath.Pop()
			}
			continue
		}
		if (p.Row == 0) && (p.Col == 0) {
			sPath.Push(p)
			break
		}

		s.Push(Point{p.Row - 1, p.Col})
		s.Push(Point{p.Row, p.Col - 1})
		sPath.Push(p)
	}

	path := make([]Point, sPath.Len())
	n := sPath.Len()
	for i := 0; i < n; i++ {
		path[i] = sPath.Pop().(Point)
	}
	return path
}
