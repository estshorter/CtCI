package p04

import (
	"fmt"
)

// AdjacentNum solves ch06 p04
func AdjacentNum(n uint) (uint, uint, error) {
	prev, err := getPrev(n)
	if err != nil {
		return 0, 0, err
	}

	next, err := getNext(n)
	if err != nil {
		return 0, 0, err
	}
	return prev, next, nil
}

func getNext(n uint) (uint, error) {
	c := n
	c0 := 0
	c1 := 0
	for ; (c&1 == 0) && (c != 0); c >>= 1 {
		c0++
	}
	for ; c&1 == 1; c >>= 1 {
		c1++
	}
	p := c0 + c1
	if p == 64 || p == 0 {
		return 0, fmt.Errorf("p = %v", c0+c1)
	}
	n |= (1 << p)
	n &^= ((1 << p) - 1)
	n |= (1 << (c1 - 1)) - 1
	return n, nil
}

func getPrev(n uint) (uint, error) {
	c := n
	c0 := 0
	c1 := 0
	for ; c&1 == 1; c >>= 1 {
		c1++
	}
	if c == 0 {
		return 0, fmt.Errorf("c becomes 0. c1 = %v", c1)
	}
	for ; (c&1 == 0) && (c != 0); c >>= 1 {
		c0++
	}
	p := c0 + c1
	n &^= ((1 << (p + 1)) - 1)
	var mask uint = ((1 << (c1 + 1)) - 1)
	n |= mask << (c0 - 1)
	return n, nil
}
