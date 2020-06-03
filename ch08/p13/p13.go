package p13

import (
	"sort"
)

// Box defines ...
type Box struct {
	w, h, d int
}

func (b *Box) canBeAbove(b2 *Box) bool {
	if b.h < b2.h && b.w < b2.w && b.d < b2.d {
		return true
	}
	return false
}

func createStack(boxes []Box) int {
	sort.Slice(boxes, func(i, j int) bool {
		return boxes[i].h > boxes[j].h
	})
	memo := make(map[int]int)
	maxHeight := 0
	for i := 0; i < len(boxes); i++ {
		tmp := createStackHelper(boxes, i, memo)
		if tmp > maxHeight {
			maxHeight = tmp
		}
	}
	return maxHeight
}

func createStackHelper(boxes []Box, index int, memo map[int]int) int {
	if _, ok := memo[index]; ok {
		return memo[index]
	}

	box := boxes[index]
	maxHeight := 0
	for i := index + 1; i < len(boxes); i++ {
		if boxes[i].canBeAbove(&box) {
			tmp := createStackHelper(boxes, index+1, memo)
			if tmp > maxHeight {
				maxHeight = tmp
			}
		}
	}
	maxHeight += box.h
	memo[index] = maxHeight
	return maxHeight
}
