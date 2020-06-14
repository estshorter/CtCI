package p03

import (
	"testing"
)

func Test_binarySearchRotate2(t *testing.T) {
	type args struct {
		a []int
		v int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14}, 5}, 8},
		{"2", args{[]int{15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14}, 15}, 0},
		{"3", args{[]int{1, 3, 4, 5, 7, 10, 14}, 5}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearchRotate2(tt.args.a, tt.args.v); got != tt.want {
				t.Errorf("binarySearchRotate2() = %v, want %v", got, tt.want)
			}
		})
	}
}
