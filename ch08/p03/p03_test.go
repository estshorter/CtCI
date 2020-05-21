package p03

import (
	"testing"
)

func Test_magicFast(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{-10, -9, -1, 1, 2, 3, 5, 7, 9, 12, 13}}, 7},
		{"2", args{[]int{-5, 0, 1, 2, 4, 10, 12}}, 4},
		{"3", args{[]int{-1, 0, 1, 2, 5, 10, 12}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := magicFast(tt.args.a); got != tt.want {
				t.Errorf("magicFast() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_magicFast2(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{-10, -9, -1, 1, 2, 3, 5, 7, 9, 12, 13}}, 7},
		{"2", args{[]int{-5, 0, 1, 2, 4, 10, 12}}, 4},
		{"3", args{[]int{1, 0, 1, 2, 5, 10, 12}}, -1},
		{"4", args{[]int{0, 1, 1, 1, 5, 5, 6, 9}}, 0},
		{"5", args{[]int{-1, 0, 1, 1, 5, 5, 6, 9}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := magicFast2(tt.args.a); got != tt.want {
				t.Errorf("magicFast2() = %v, want %v", got, tt.want)
			}
		})
	}
}
