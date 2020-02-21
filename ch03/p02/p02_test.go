package p02

import (
	"testing"
)

func TestStackWithMin_Pop(t *testing.T) {
	s := New()
	s.Push(2)
	s.Push(5)
	s.Push(7)
	s.Push(-1)
	s.Push(-1)

	min := []int{-1, -1, 2, 2, 2}
	// min := []int{-10, -10, -10, -10, -10}
	for i := 0; i < 5; i++ {
		got := s.Min()
		want := min[i]
		if got != want {
			t.Errorf("Got %v, want %v", got, want)
		}
		s.Pop()
	}

}
