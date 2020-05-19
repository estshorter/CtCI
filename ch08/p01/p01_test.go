package p01

import (
	"testing"
)

func Test_countStepNaive(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1", args{4}, 7,
		},
		{
			"2", args{5}, 13,
		},
		{
			"30", args{30}, 53798080,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countStepNaive(tt.args.n); got != tt.want {
				t.Errorf("countStepNaive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countStepMemo(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1", args{4}, 7,
		},
		{
			"2", args{5}, 13,
		},
		{
			"30", args{30}, 53798080,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countStepMemo(tt.args.n); got != tt.want {
				t.Errorf("countStepMemo() = %v, want %v", got, tt.want)
			}
		})
	}
}
