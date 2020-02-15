package p02

import "testing"

func Test_checkPermutation(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"different length",
			args{s1: "hoge", s2: "hogehoge"},
			false,
		},
		{
			"reverse",
			args{s1: "hoge", s2: "egoh"},
			true,
		},
		{
			"dup",
			args{s1: "hogggge", s2: "egggohg"},
			true,
		},
		{
			"dup",
			args{s1: "hogggge", s2: "ggggohg"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckPermutation(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("checkPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
