package p06

import "testing"

func TestSimpleCompress(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"aabcccaaa"}, "a2b1c3a3"},
		{"2", args{"abcd"}, "abcd"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SimpleCompress(tt.args.s); got != tt.want {
				t.Errorf("SimpleCompress() = %v, want %v", got, tt.want)
			}
		})
	}
}
