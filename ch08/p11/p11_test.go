package p11

import "testing"

func Test_makeChange(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"100", args{100}, 242},
		{"20", args{20}, 9},
		{"10", args{10}, 4},
		{"5", args{5}, 2},
		{"1", args{1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeChange(tt.args.n); got != tt.want {
				t.Errorf("makeChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
