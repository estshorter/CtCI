package p05

import "testing"

func Test_sortWithEmpty(t *testing.T) {
	type args struct {
		a []string
		x string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]string{"at", "", "", "", "ball", "", "", "car", "", "", "dad", "", ""}, "ball"}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortWithEmpty(tt.args.a, tt.args.x); got != tt.want {
				t.Errorf("sortWithEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
