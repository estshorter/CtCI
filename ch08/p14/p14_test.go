package p14

import "testing"

func Test_countEval(t *testing.T) {
	type args struct {
		s      string
		result bool
		memo   map[Key]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"1^0|0|1", false, make(map[Key]int)}, 2},
		{"1", args{"0&0&0&1^1|0", true, make(map[Key]int)}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countEval(tt.args.s, tt.args.result, tt.args.memo); got != tt.want {
				t.Errorf("countEval() = %v, want %v", got, tt.want)
			}
		})
	}
}
