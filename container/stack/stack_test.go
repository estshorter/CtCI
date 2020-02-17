package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	tests := []struct {
		name string
		vals []int
	}{
		{"1", []int{1, 2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := New()
			for _, val := range tt.vals {
				stack.Push(val)
			}
			for idx := range tt.vals {
				popped := stack.Pop()
				val := tt.vals[len(tt.vals)-1-idx]
				if val != popped {
					t.Errorf("IsPalindrome2() = %v, want %v", popped, val)
				}

			}
		})
	}
}

func TestPop(t *testing.T) {
	a := []int{1, 2, 3, 4}
	stack := New()
	for _, val := range a {
		stack.Push(val)
	}
	for idx := range a {
		stack.Pop()
		if idx == 2 && cap(stack.a) != 2 {
			t.Errorf("wrong cap %v, want %v", cap(stack.a), 2)
		}
		if idx == 3 && cap(stack.a) != 1 {
			t.Errorf("wrong cap %v, want %v", cap(stack.a), 0)
		}
	}
}
