package p04

import "testing"

func TestMyQueue(t *testing.T) {
	s := New()
	s.Enqueue(1)
	s.Enqueue(2)
	s.Enqueue(3)
	s.Enqueue(4)
	s.Enqueue(5)
	s.Enqueue(6)
	s.Enqueue(7)
	s.Enqueue(8)

	for i := 0; i < 8; i++ {
		got := s.Dequeue()
		want := i + 1
		if got != want {
			t.Errorf("Got %v, want %v", got, want)
		}
	}

}
