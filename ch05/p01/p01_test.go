package p01

import "testing"

func TestInsert(t *testing.T) {
	type args struct {
		n uint32
		m uint32
		i uint32
		j uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{"case1", args{0b10000000000, 0b10011, 3, 7}, 0b10010011000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Insert(tt.args.n, tt.args.m, tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Insert() = %b, want %b", got, tt.want)
			}
		})
	}
}
