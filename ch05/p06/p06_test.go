package p06

import "testing"

func Test_convertBit(t *testing.T) {
	type args struct {
		a uint
		b uint
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{0b11101, 0b01111}, 2},
		{"case1", args{0b1011101, 0b0101111}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertBit(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("convertBit() = %v, want %v", got, tt.want)
			}
		})
	}
}
