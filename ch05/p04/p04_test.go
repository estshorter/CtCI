package p04

import "testing"

func TestAdjacentNum(t *testing.T) {
	type args struct {
		v uint
	}
	tests := []struct {
		name    string
		args    args
		want    uint
		want1   uint
		wantErr bool
	}{
		{"case1", args{0b11011}, 0b10111, 0b11101, false},
		{"case2", args{0b00011}, 0, 0, true},
		{"case3", args{1 << 63}, 0, 0, true},
		{"case4", args{(1<<62 | 1<<60)}, (1<<62 | 1<<59), (1<<62 | 1<<61), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := AdjacentNum(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("AdjacentNum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AdjacentNum() got = %b, want %b", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("AdjacentNum() got1 = %b, want %b", got1, tt.want1)
			}
		})
	}
}
