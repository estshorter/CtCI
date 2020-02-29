package p03

import "testing"

func TestCountSucc(t *testing.T) {
	type args struct {
		v uint
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{0b10101110}, 5},
		{"case1", args{1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountSucc(tt.args.v); got != tt.want {
				t.Errorf("CountSucc() = %v, want %v", got, tt.want)
			}
		})
	}
}
