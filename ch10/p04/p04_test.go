package p04

import "testing"

func Test_sortP04(t *testing.T) {
	type args struct {
		a []int
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{0, 1, 2, 3, 4, 5, 6, -1, -1, -1, -1, -1}, 3}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortP04(tt.args.a, tt.args.x); got != tt.want {
				t.Errorf("sortP04() = %v, want %v", got, tt.want)
			}
		})
	}
}
