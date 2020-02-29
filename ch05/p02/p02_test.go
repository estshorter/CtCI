package p02

import (
	"testing"
)

func Test_printBinary(t *testing.T) {
	type args struct {
		v float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case3", args{0.5}, ".1"},
		{"case3", args{0.875}, ".111"},
		{"case4", args{1.0 / 2147483648.0}, ".0000000000000000000000000000001"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrintBinary(tt.args.v); got != tt.want {
				t.Errorf("printBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
