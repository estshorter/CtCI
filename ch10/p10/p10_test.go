package p10

import "testing"

func Test_getRankOfNumber(t *testing.T) {
	track(20)
	track(15)
	track(10)
	track(5)
	track(13)
	track(25)
	track(23)
	track(24)
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{10}, 1},
		{"2", args{24}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRankOfNumber(tt.args.number); got != tt.want {
				t.Errorf("getRankOfNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
