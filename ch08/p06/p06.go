package p06

import (
	"ctci/container/stack"
	"fmt"
)

type tower struct {
	Name string
	stack.Stack
}

func newTower(name string) *tower {
	return &tower{name, *stack.New()}
}

func hanoi(n int) *tower {
	src := newTower("src")
	dst := newTower("dst")
	buf := newTower("buf")
	for i := n; i > 0; i-- {
		src.Push(i)
	}
	hanoiHelper(n, src, dst, buf)

	// fmt.Println(src)
	fmt.Println(*dst)
	// fmt.Println(buf)
	return dst
}

func hanoiHelper(n int, src, dst, buf *tower) {
	if n == 1 {
		move(src, dst)
		return
	}

	hanoiHelper(n-1, src, buf, dst)
	move(src, dst)
	hanoiHelper(n-1, buf, dst, src)
}

func move(src, dst *tower) {
	// fmt.Printf("%v -> %v, %v\n", src.Name, dst.Name, src.Peek())
	dst.Push(src.Pop())
}
