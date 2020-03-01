package p07

import "testing"

func TestSwapOddEvenBits(t *testing.T) {
	type args struct {
		x uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{"case1", args{0b10101010}, 0b01010101},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SwapOddEvenBits(tt.args.x); got != tt.want {
				t.Errorf("SwapOddEvenBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
