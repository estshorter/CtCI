package p03

import (
	"testing"
)

func Test_UrlifyTest(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"golang introduction",
			args{"golang introduction"},
			"golang%20introduction",
		},
		{
			"hoge",
			args{"hoge"},
			"hoge",
		},
		{
			"Keyboard shortcuts for Windows",
			args{"Keyboard shortcuts for Windows"},
			"Keyboard%20shortcuts%20for%20Windows",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UrlifyTest(tt.args.s); got != tt.want {
				t.Errorf("urlifyTest() = %v, want %v", got, tt.want)
			}
		})
	}
}
