package p05

import (
	"ctci/container/stack"
	"reflect"
	"testing"
)

func Test_sort(t *testing.T) {
	s1 := stack.New()
	s1.Push(2)
	s1.Push(3)
	s1.Push(1)
	s1.Push(4)
	s1.Push(5)
	s1.Push(2)
	s2 := stack.New()
	s2.Push(5)
	s2.Push(4)
	s2.Push(3)
	s2.Push(2)
	s2.Push(2)
	s2.Push(1)

	type args struct {
		s *stack.Stack
	}
	tests := []struct {
		name string
		args args
		want *stack.Stack
	}{
		{"case1", args{s1}, s2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sort(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sort() = %v, want %v", got, tt.want)
			}
		})
	}
}
