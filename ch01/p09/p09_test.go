package p09

import "testing"

func TestIsRotatedMatrix(t *testing.T) {
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
			"true",
			args{"waterbottle", "erbottlewat"},
			true,
		},
		{
			"false",
			args{"waterbotte", "erbottlewat"},
			false,
		},
		{
			"false2",
			args{"abcdef", "abcdfe"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsRotatedMatrix(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("IsRotatedMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
