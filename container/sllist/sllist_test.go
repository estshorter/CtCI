package sllist

import (
	"testing"
)

func TestInitBySlce(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	l := InitBySlice(a)
	e := l.Front()
	for _, aval := range a {
		v := e.Value
		if v != aval {
			t.Errorf("wrong list value %v, want %v", v, aval)
		}
		e = e.Next
	}
}
