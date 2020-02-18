package queue

import (
	"testing"
)

func TestEnqueDeque(t *testing.T) {
	tests := []struct {
		name string
		vals []int
	}{
		{"1", []int{1, 2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := New()
			for _, val := range tt.vals {
				q.Enqueue(val)
			}
			for _, val := range tt.vals {
				deque := q.Dequeue()
				if val != deque {
					t.Errorf("Dequeue() = %v, want %v", deque, val)
				}

			}
		})
	}
}

func TestCapSize(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	q := New()
	for _, val := range a {
		q.Enqueue(val)
	}
	for idx := range a {
		q.Dequeue()
		if idx == 3 && cap(q.a) != 10 {
			t.Errorf("wrong cap %v, want %v", cap(q.a), 10)
		}
		if idx == 5 && cap(q.a) != 6 {
			t.Errorf("wrong cap %v, want %v", cap(q.a), 6)
		}
		if idx == 6 && cap(q.a) != 4 {
			t.Errorf("wrong cap %v, want %v", cap(q.a), 4)
		}
		if idx == 7 && cap(q.a) != 2 {
			t.Errorf("wrong cap %v, want %v", cap(q.a), 2)
		}
		if idx == 8 && cap(q.a) != 1 {
			t.Errorf("wrong cap %v, want %v", cap(q.a), 1)
		}
	}
}
