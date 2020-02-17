package p06

import (
	"ctci/container/list"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	l1 := list.InitBySlice([]int{1, 0, 7, 0, 1})
	l2 := list.InitBySlice([]int{1, 0, 0, 1})
	l3 := list.InitBySlice([]int{1, 0, 1, 1})
	l4 := list.InitBySlice([]int{1, 0, 5, 1, 1})
	type args struct {
		l *list.List
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1", args{l1}, true},
		{"case1", args{l2}, true},
		{"case1", args{l3}, false},
		{"case1", args{l4}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.args.l); got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPalindrome2(t *testing.T) {
	l1 := list.InitBySlice([]int{1, 0, 7, 0, 1})
	l2 := list.InitBySlice([]int{1, 0, 0, 1})
	l3 := list.InitBySlice([]int{1, 0, 1, 1})
	l4 := list.InitBySlice([]int{1, 0, 5, 1, 1})

	type args struct {
		l *list.List
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1", args{l1}, true},
		{"case1", args{l2}, true},
		{"case1", args{l3}, false},
		{"case1", args{l4}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome2(tt.args.l); got != tt.want {
				t.Errorf("IsPalindrome2() = %v, want %v", got, tt.want)
			}
		})
	}
}
