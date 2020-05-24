package p10

import (
	"fmt"
	"strconv"

	"github.com/golang-collections/collections/queue"
)

// Color defines color of a pixel
type Color int

// Point defines 2d position
type Point struct {
	row, col int
}

const (
	black Color = iota
	white
	red
)

func (c Color) String() string {
	switch c {
	case black:
		return "b"
	case white:
		return "w"
	case red:
		return "r"
	default:
		return strconv.Itoa(int(c))
	}
}

func fill(img [][]Color, row, col int, fillColor Color, p Point) [][]Color {
	origColor := img[p.row][p.col]
	q := queue.New()
	q.Enqueue(p)
	for {
		if q.Len() == 0 {
			return img
		}
		p := q.Dequeue().(Point)
		if p.row < 0 || p.col < 0 || p.row >= row || p.col >= col || img[p.row][p.col] != origColor {
			continue
		}
		fmt.Println(p)
		img[p.row][p.col] = fillColor
		q.Enqueue(Point{p.row - 1, p.col})
		q.Enqueue(Point{p.row + 1, p.col})
		q.Enqueue(Point{p.row, p.col + 1})
		q.Enqueue(Point{p.row, p.col - 1})
	}
}
