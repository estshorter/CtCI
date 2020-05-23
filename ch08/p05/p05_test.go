package p05

import "testing"

func Test_minProduct(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{31, 35}, 1085},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minProduct(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("minProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
