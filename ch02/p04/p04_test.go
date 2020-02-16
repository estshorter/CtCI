package p04

import (
	"ctci/container/list"
	"fmt"
	"testing"
)

func equal(l1 *list.List, l2 *list.List) bool {
	if l1.Len() != l2.Len() {
		fmt.Printf("wrong length: got %v, want %v\n", l1.Len(), l2.Len())
		return false
	}

	v1 := make([]int, l1.Len())
	v2 := make([]int, l2.Len())

	e1 := l1.Front()
	e2 := l2.Front()

	equal := true
	for i := 0; i < l1.Len(); i++ {
		v1[i] = e1.Value
		v2[i] = e2.Value

		if e1.Value != e2.Value {
			equal = false
		}
		e1 = e1.Next
		e2 = e2.Next
	}
	fmt.Printf("got %v, want %v\n", v1, v2)
	return equal
}

func TestPartition(t *testing.T) {
	a1 := [...]int{3, 5, 10, 5, 8, 2, 1}
	a2 := [...]int{3, 2, 1, 5, 10, 5, 8}
	l1 := list.NewWithSlice(a1[:])
	l2 := list.NewWithSlice(a2[:])

	type args struct {
		l *list.List
		x int
	}
	tests := []struct {
		name string
		args args
		want *list.List
	}{
		{"case1", args{l1, 5}, l2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Partition(tt.args.l, tt.args.x); !equal(got, tt.want) {
				t.Errorf("PseudoSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
