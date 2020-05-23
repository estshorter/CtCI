package p06

import "testing"

func Test_hanoi(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
	}{
		{"1", args{1}},
		{"2", args{2}},
		{"4", args{4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dst := hanoi(tt.args.n)
			for i := 1; i <= tt.args.n; i++ {
				if dst.Pop() != i {
					panic(*dst)
				}
			}
		})
	}
}
