package p05

import "testing"

func TestCanOneTimeConvert(t *testing.T) {
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
			"pale vs ple",
			args{"pale", "ple"},
			true,
		},
		{
			"pales vs pale",
			args{"pales", "pale"},
			true,
		},
		{
			"pale vs bale",
			args{"pale", "bale"},
			true,
		},
		{
			"pale vs bake",
			args{"pale", "bake"},
			false,
		},
		{
			"pale vs paleeeee",
			args{"pale", "paleeeee"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanOneTimeConvert(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("CanOneTimeConvert() = %v, want %v", got, tt.want)
			}
		})
	}
}
