package p04

import "testing"

func Test_checkPalindromePermutation(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"hoge",
			args{"hoge"},
			false,
		},
		{
			"aabb cc ",
			args{"aabb cc "},
			true,
		},
		{
			" aaa bb cc",
			args{" aaa bb cc"},
			true,
		},
		{
			"bbbb ccccc a",
			args{"bbbb ccccc a"},
			false,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPalindromePermutation(tt.args.s); got != tt.want {
				t.Errorf("checkPalindromePermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
